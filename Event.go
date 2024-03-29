package index

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"unsafe"
)

var ApiResponse = events.APIGatewayProxyResponse{}
var Apireqestid string

func SQSEventCall(msg json.RawMessage) {

	eh := events.SQSEvent{}

	err := json.Unmarshal(msg, &eh)
	if err != nil {
		log.Println("error in SQS.json file", err)
	}
	for _, record := range eh.Records {

		log.Printf("\nRECORDS:- %d", unsafe.Sizeof(record))
		log.Printf("\nMesssageId:- %s \nEventSource:- %s \nBody:- %s \nAttributes:- %s \n", record.MessageId, record.EventSource, record.Body, record.Attributes)
	}
}

func SNSEventCall(msg json.RawMessage) {
	eh := events.SNSEvent{}

	err := json.Unmarshal(msg, &eh)
	if err != nil {
		log.Println("error in SNS.json file", err)
	}
	for _, record := range eh.Records {
		snsRecord := record.SNS
		log.Printf("\nRECORDS:- %d", unsafe.Sizeof(record))
		log.Printf("\nEventSource:- %s \nEventVersion:- %s \nTimestamp: %s \nMessage:- %s \nSignature:- %s \nSigningCertURL:- %s \n", record.EventSource, record.EventVersion, snsRecord.Timestamp, snsRecord.Message, snsRecord.Signature, snsRecord.SigningCertURL)
	}
}

func S3EventCall(msg json.RawMessage) {
	eh := events.S3Event{}

	err := json.Unmarshal(msg, &eh)
	if err != nil {
		log.Println("error in S3.json file", err)
	}
	for _, record := range eh.Records {
		s3 := record.S3
		log.Printf("\nRECORDS:- %d\nEventName:= %s \nEventSource:- %s\nEventVersion:- %s\nSourceIPAddress:- %s\nBucketName:- %s \nObjectKey:- %s \n", unsafe.Sizeof(record), record.EventName, record.EventSource, record.EventVersion, record.RequestParameters.SourceIPAddress, s3.Bucket.Name, s3.Object.Key)

	}
}
func ApiGatewayCall(msg json.RawMessage, reqHeader string) string {
	request := events.APIGatewayProxyRequest{}
	err := json.Unmarshal(msg, &request)
	if err != nil {
		log.Println("error in ApiGateway.json file", err)
	}
	Url_path = request.Path + "|" + request.HTTPMethod
	Apireqestid = request.RequestContext.RequestID
	return MakeHeader(reqHeader, request.Headers)

}

func MakeHeader(Header string, request map[string]string) string {
	var btheader string
	Bt_header = nil
	//btheader = " "
	b := []rune{127}
	for key, value := range request {
		Header += key + "$"
		Header += value
		Header += string(b)
		btheader += key + "=" + value + "&"
	}
	Bt_header = &btheader
	return Header
}

func ApiResponseCall(msg []byte, respHeader string) (string, events.APIGatewayProxyResponse) {

	err := json.Unmarshal(msg, &ApiResponse)
	if err != nil {
		log.Println("error in ApiGatewayResponse.json file", err)
	}

	if len(ApiResponse.Headers) == 0 {
		ApiResponse.Headers = map[string]string{"ResponseHeaders": "NotFound"}
	}
	if Ckheader != "" {
		ApiResponse.Headers = map[string]string{"set-cookie": Ckheader}
	}
	respHeader = MakeHeader(respHeader, ApiResponse.Headers)
	return respHeader, ApiResponse
}
