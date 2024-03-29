package index

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	//"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func AmplifyApps(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	var config amplify.ListAppsInput
	amplifyClient := amplify.NewFromConfig(*client.AWSClient)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amplify")
	if err != nil {
		log.Warnln("Couldn't fetch Amplify cost and usage:", err)
	}
	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AWSAmplify"),
		Filters: []types.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(client.AWSClient.Region),
				Type:  types.FilterTypeTermMatch,
			},
		},
	})
	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS Amplify: %v", err)
		return resources, err
	}

	priceMap , err := awsUtils.GetPriceMap(pricingOutput, "group") // priceMap
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
		return resources, err
	}
	fmt.Println("map come in ",priceMap)

	for {
		output, err := amplifyClient.ListApps(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, app := range output.Apps {
            		estimatedCost := 0.0

					metricsDurationOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
						StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
						EndTime:    aws.Time(time.Now()),
						MetricName: aws.String("App Duration"),
						Namespace:  aws.String("AWS/Amplify"),
						Dimensions: []cloudwatchTypes.Dimension{
							{
								Name:  aws.String("AppName"),
								Value: app.Name,
							},
						},
						Period: aws.Int32(3600),
						Statistics: []cloudwatchTypes.Statistic{
							cloudwatchTypes.StatisticSum,
						},
					})
            		// if app.Tier == types.AppTierBasic {
                	// 	estimatedCost = 10.0 
            		// } else if app.Tier == types.AppTierPremium {
                	// 	estimatedCost = 20.0 
            		// } else {
                	// 	estimatedCost = 0.5  
            		//}
				if metricsDurationOutput != nil && len(metricsDurationOutput.Datapoints) > 0 {
					estimatedCost = *metricsDurationOutput.Datapoints[0].Average
				}
				tags := make([]models.Tag, 0)
				tagsResp, err := amplifyClient.ListTagsForResource(context.Background(), &amplify.ListTagsForResourceInput{
					ResourceArn: app.AppArn,  
				})
		
				if err == nil {
					for key, value := range tagsResp.Tags {
						tags = append(tags, models.Tag{
							Key:   key,
							Value: value,
						})
					}
				}
		
			

            	resources = append(resources, models.Resource{
                	Provider:  "AWS",
                	Account:   client.Name,
                	Service:   "Amplify",
                	ResourceId: *app.AppId,
                	Region:    client.AWSClient.Region,
                	Name:      *app.Name,
                	Cost:      estimatedCost,
                	Metadata: map[string]string{
                    		"serviceCost": fmt.Sprint(serviceCost),
                	},
					Tags:       tags,
                	FetchedAt: time.Now(),
                	Link: fmt.Sprintf("https://console.aws.amazon.com/amplify/v1/home/#/apps/%s", *app.AppId),
            		})
				}


		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Amplify",
		"resources": len(resources),
	}).Info("Fetched Amplify apps")

	return resources, nil
}
