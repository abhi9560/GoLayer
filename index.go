package index

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-lambda-go/events"
	//"errors"
	"log"
	"reflect"
	"runtime"
	
	
)



var (
	// CurrentContext is the last create lambda context object.
	CurrentContext context.Context
)

var Apireqid  = ""
func WrapHandler(handler interface{}) interface{} {

	
	coldStart := true

	return func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
		
		
		ctx = context.WithValue(ctx, "cold_start", coldStart)
		
		UDPConnection()
		url_path := lambdacontext.FunctionName
		nvValue := NVCookieMessage(ctx)
   		ndValue := ""//NDCookieMessage(ctx)
		StartTransactionMessage(ctx ,url_path, "",ndValue,nvValue)
		handlerType := reflect.TypeOf(handler)
		if handlerType.NumIn() == 0 {
			return reflect.ValueOf(nil), nil
		}
	
		messageType := handlerType.In(handlerType.NumIn() - 1)
		
		
		var methodName string
		reqHeader := ""
		
		Sqs  	:= reflect.TypeOf(events.SQSEvent{})
		Sns  	:= reflect.TypeOf(events.SNSEvent{}) 
		S3  	:= reflect.TypeOf(events.S3Event{})
		ApiReq 	:= reflect.TypeOf(events.APIGatewayProxyRequest{})
		
		switch messageType{
			case Sqs:
				methodName = "index.SQSEventHandler"
				SQSEventCall(msg)
			
			case Sns :
				
				methodName = "index.SNSEventHandler"
				SNSEventCall(msg)
			case S3 :
		        	
				methodName = "index.S3EndpointHandler"
				S3EventCall(msg)
			case ApiReq :
				
				methodName = "index.ApiEndpointHandler"
				reqHeader,Apireqid = ApiGatewayCall(msg,reqHeader)
				log.Println("apireqid",Apireqid)
			
			default:
				methodName = runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		
			
		}
		statuscode := 200
		method_entry(ctx,methodName)
		if reqHeader != ""{
			//SendReqRespHeder(ctx ,reqHeader , "req" ,statuscode)
		}
		CurrentContext = ctx
		//NDNFCookieMessage(ctx)
		result, err := callHandler(ctx,msg, handler,messageType)
		if err != nil {
			statuscode = 500
		}
	
		method_exit(ctx,methodName,statuscode)
		
		end_business_transaction(ctx,statuscode)
		
		//CloseUDP()
		coldStart = false
		CurrentContext = nil
		return result, err
	}
}

func callHandler(ctx context.Context,msg json.RawMessage, handler interface{},messageType reflect.Type) (interface{}, error) {
	ev, err := unmarshalEventForHandler(msg,messageType)
	if err != nil {
		return nil, err
	}
	handlerType := reflect.TypeOf(handler)
	

	args := []reflect.Value{}
	
	if handlerType.NumIn() == 1 {
		// When there is only one argument, argument is either the event payload, or the context.
		contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
		firstArgType := handlerType.In(0)
		if firstArgType.Implements(contextType) {
			args = []reflect.Value{reflect.ValueOf(ctx)}
		} else {
			args = []reflect.Value{ev.Elem()}

		}
	} else if handlerType.NumIn() == 2 {
		// Or when there are two arguments, context is always first, followed by event payload.
		args = []reflect.Value{reflect.ValueOf(ctx), ev.Elem()}
		
		
	}

	handlerValue := reflect.ValueOf(handler)
	
	log.Println("handlerValue   ",handlerValue)
	output := handlerValue.Call(args)
	log.Println("full output",output)
	var response interface{}
	
	var errResponse error

	if len(output) > 0 {
		// If there are any output values, the last should always be an error
		val := output[len(output)-1].Interface()
		log.Println("val of output",val)
		if errVal, ok := val.(error); ok {
			errResponse = errVal
		}
	}
	
	if len(output) > 1 {
		// If there is more than one output value, the first should be the response payload.
		response    = output[0].Interface()
		ApiResp    := reflect.TypeOf(events.APIGatewayProxyRequest{})
		
		
		switch messageType{
			
			case ApiResp:
				respHeader := ""
				str, err := json.Marshal(response)
				if err != nil {
					log.Println("unable to marshal json", err)
				}
				respHeader,statuscode := ApiResponseCall(str,respHeader)
				
				log.Println("all header value 2",statuscode)
        			//SendReqRespHeder(ctx ,respHeader , "resp" ,statuscode)
		}
		
	}

	return response, errResponse
}

func unmarshalEventForHandler(ev json.RawMessage, messageType reflect.Type) (reflect.Value, error) {
	
	newMessage := reflect.New(messageType)
	json.Unmarshal(ev,newMessage.Interface())
	
	return newMessage, err
}

