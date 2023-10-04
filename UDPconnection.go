package index

/*

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "ndlb_encode.h"

transactionStart_t node;
msgHdr_t msgHdr;
wrapheader_t wrapHeader;
MethodEntry_t node1;
MethodExit_t node2;
transactionEnd_t transactionEnd;
transactionEncodeHttp_t node3;

int WrapHeader(char *s){
    int len = 0;
    memcpy(s, "^",1);
    len += 1;
    memcpy(s + len , (char *)&(wrapHeader.wrapheadervar), sizeof(wrapHeader.wrapheadervar));
    len += sizeof(wrapheadervar_t);

  return len;
}
int WrapHeaderlen(int apiReqLen,int awsReqLen,int funcNameLen,int tagslength,short agentType,short messageType){
    int len = 0;
    
    wrapHeader.wrapheadervar.apiReqLen = apiReqLen ;
    wrapHeader.wrapheadervar.awsReqLen = awsReqLen ;
    wrapHeader.wrapheadervar.funcNameLen = funcNameLen ;
    wrapHeader.wrapheadervar.tagslength = tagslength ;
    wrapHeader.wrapheadervar.agentType = agentType ;
    wrapHeader.wrapheadervar.messageType = messageType;
    wrapHeader.wrapheadervar.fpId = 0;
    wrapHeader.wrapheadervar.appId = -1;
    wrapHeader.wrapheadervar.protocolType = PROTOCOL_UDP;
    wrapHeader.wrapheadervar.whLen = sizeof(wrapheadervar_t)+wrapHeader.wrapheadervar.awsReqLen+wrapHeader.wrapheadervar.apiReqLen+wrapHeader.wrapheadervar.funcNameLen+wrapHeader.wrapheadervar.tagslength+1;
    
    len += sizeof(wrapheadervar_t);
    len += 1;
  return len;
}
int ValueStore(char *s,char *value,int len,int number)
{
      memcpy(s + len, value ,number);
      len += number;
      return len;
}

int StartTransactionlen(int fp_header,int url,int btHeaderValue,int ndCookieSet,int nvCookieSet,int correlationHeader,long long flowpathinstance,long long qTimeMS,long long startTimeFP,int len)
{
    node.transactionStartVar.qTimeMS = qTimeMS ;
    node.transactionStartVar.startTimeFP = startTimeFP;
    node.transactionStartVar.fp_header = fp_header ;
    node.transactionStartVar.url = url ;
    node.transactionStartVar.btHeaderValue = btHeaderValue ;
    node.transactionStartVar.ndCookieSet = ndCookieSet ;
    node.transactionStartVar.nvCookieSet = nvCookieSet ;
    node.transactionStartVar.correlationHeader = correlationHeader ;
    node.transactionStartVar.flowpathinstance = flowpathinstance;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof (transactionStartVar_t) + msgHdr.header_len + node.transactionStartVar.fp_header +
    node.transactionStartVar.ndCookieSet + node.transactionStartVar.nvCookieSet +
    node.transactionStartVar.correlationHeader +
    node.transactionStartVar.btHeaderValue +
    node.transactionStartVar.url + 3;
    msgHdr.msg_type = 2;

    len += sizeof(msgHdr);
    len += 2;
    len += sizeof(node.transactionStartVar) ;

    return len;
}
int StartTransaction(char *s,int len)
{
    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.transactionStartVar), sizeof(node.transactionStartVar));
    len += sizeof(node.transactionStartVar) ;
    return len;
}

int MethodEntryFunction(char *s,int len)
{
    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node1.MethodEntryVar), sizeof(MethodEntryVar_t));
    len += sizeof(MethodEntryVar_t);

    return len;
}
int MethodEntryFunctionlen(int urlParameter,int methodName,int query_string,int mid,long long flowpathinstance,long long threadId,long long startTime,int len)
{

    node1.MethodEntryVar.methodName = methodName ;
    node1.MethodEntryVar.threadId= threadId;
    node1.MethodEntryVar.query_string = query_string ;
    node1.MethodEntryVar.urlParameter = urlParameter ;
    node1.MethodEntryVar.mid = mid ;
    node1.MethodEntryVar.startTime = startTime ;
    node1.MethodEntryVar.flowpathinstance = flowpathinstance;
    node1.MethodEntryVar.query_parameter = 0;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof(MethodEntryVar_t) + msgHdr.header_len + node1.MethodEntryVar.methodName +
    node1.MethodEntryVar.query_string+ node1.MethodEntryVar.urlParameter + 3;
    msgHdr.msg_type = 0;
    
    len += sizeof(msgHdr);
    len += 2;
    len += sizeof(MethodEntryVar_t);
    return len;
}

int MethodExitFunction(char *s,int len)
{
   
    memcpy(s+len, "^",1);
    len += 1;

    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);

    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node2.MethodExitVar), sizeof(node2.MethodExitVar));
    len += sizeof(MethodExitVar_t);

    return len;
}

int MethodExitFunctionlen(int statusCode,int mid,int eventType,int isCallout,long long duration,long long threadId,long long cpuTime,long long flowpathinstance,long long tierCallOutSeqNum,long long endTime,int methodName,int backend_header,int requestNotificationPhase,int len)
{

    node2.MethodExitVar.statusCode = statusCode ;
    node2.MethodExitVar.mid = mid ;
    node2.MethodExitVar.eventType = eventType;
    node2.MethodExitVar.isCallout = isCallout;
    node2.MethodExitVar.duration = duration;
    node2.MethodExitVar.threadId = threadId;
    node2.MethodExitVar.cpuTime = cpuTime;
    node2.MethodExitVar.tierCallOutSeqNum = tierCallOutSeqNum ;
    node2.MethodExitVar.endTime = endTime ;
    node2.MethodExitVar.methodName = methodName ;
    node2.MethodExitVar.backend_header = backend_header;
    node2.MethodExitVar.requestNotificationPhase = requestNotificationPhase ;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof(MethodExitVar_t)+ msgHdr.header_len + node2.MethodExitVar.methodName + node2.MethodExitVar.backend_header +
    node2.MethodExitVar.requestNotificationPhase + 3;
    msgHdr.msg_type = 1;
    
    len += sizeof(msgHdr);
    len += 2;
    len += sizeof(MethodExitVar_t);

    return len;
}

int EndTransaction(char *s,int len)
{
    int N=sizeof (transactionEnd_t);
    memcpy(s+len, "^",1);
    len += 1;

    memcpy(s + len , (char *)&(msgHdr), sizeof(msgHdr_t));
    len += sizeof(msgHdr_t);

    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (void *)&transactionEnd, N);
    len += N;

    return len;

}

int EndTransactionlen(int statuscode,long long endTime,long long flowpathinstance,long long cpuTime ,int len)
{
    int N = sizeof(transactionEnd_t);
    transactionEnd.statuscode = statuscode ;
    transactionEnd.endTime = endTime ;
    transactionEnd.flowpathinstance = flowpathinstance;
    transactionEnd.cpuTime = cpuTime;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = msgHdr.header_len + N + 3;
    msgHdr.msg_type = 3;

    len += sizeof(msgHdr_t);
    len += 2;
    len += N;

    return len;

}
int ReqRespHeader(char *s,int len)
{
    memcpy(s+len, "^",1);
    len += 1;

    
    memcpy(s + len , (char *)&(msgHdr), sizeof(msgHdr_t));
    len += sizeof(msgHdr_t);

    memcpy(s + len , "|",1);
    len += 1;

    memcpy(s + len, (char *)&(node3.transactionEncodeVarHttp), sizeof(node3.transactionEncodeVarHttp));
    len += sizeof(transactionEncodeVarHttp_t);

    return len;
}
int ReqRespHeaderlen(int statuscode,int buffer_len,int type_len,long long flowpathinstance,int len)
{
    
    node3.transactionEncodeVarHttp.statuscode = statuscode;
    node3.transactionEncodeVarHttp.buffer_len = buffer_len;
    node3.transactionEncodeVarHttp.type_len = type_len;
    node3.transactionEncodeVarHttp.flowpathinstance = flowpathinstance;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = msgHdr.header_len + sizeof(transactionEncodeVarHttp_t)  + node3.transactionEncodeVarHttp.type_len +
    node3.transactionEncodeVarHttp.buffer_len + 3;
    msgHdr.msg_type = 6;

    len += sizeof(msgHdr_t);
    len += 2;
    len += sizeof(transactionEncodeVarHttp_t);

    return len;
}

int last(char *s,int len)
{
    memcpy(s + len, "\n", 1);
}
*/
import "C"

import (
    "context"
    "encoding/json"
    //"fmt"
    "github.com/aws/aws-lambda-go/lambdacontext"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
    "log"
    "net"
    "os"
    "strings"
    "time"
    "unsafe"
)

var Ndcookie string
var Nvcookie string
var NDCOOKIE_VALUE string
var Ckheader string

type appRequest struct {
    APIReqId string
    AWSReqId string
    FuncName string
    Tags     string
}
var Appinfo appRequest
func Header(buf []byte) C.int {

    len := C.WrapHeader((*C.char)(unsafe.Pointer(&buf[0])))

    len = sendperameter(buf, len, Appinfo.APIReqId, Appinfo.AWSReqId, Appinfo.FuncName, Appinfo.Tags)
    return len
}
func Headerlen(msgType C.short,ctx context.Context) C.int {
    
    
    if lambdacontext.FunctionName != "" {
        Appinfo.FuncName = lambdacontext.FunctionName
        //log.Printf("FUNCTION NAME1: %s", lambdacontext.FunctionName)
    } else {
        Appinfo.FuncName = "main_test1"
    }
    if Apireqestid != "" {
        Appinfo.APIReqId  = Apireqestid
    } else {
        Appinfo.APIReqId  = "NotFound"
    }
    lc, _ := lambdacontext.FromContext(ctx)
    if lc.AwsRequestID != "" {
        Appinfo.AWSReqId = lc.AwsRequestID
        //log.Printf("REQUEST ID: %s", lc.AwsRequestID)
    } else {
        Appinfo.AWSReqId = "NotFound"
    }
    var Tier, server, appName, tagkey string
    if os.Getenv("CAV_APP_AGENT_TIER") == "" {
        Tier = "default"
    } else {
        Tier = os.Getenv("CAV_APP_AGENT_TIER")
    }
    if os.Getenv("CAV_APP_AGENT_SERVER") == "" {
        server = "default"
    } else {
        server = os.Getenv("CAV_APP_AGENT_SERVER")
    }
    if os.Getenv("CAV_APP_AGENT_INSTANCE") == "" {
        appName = os.Getenv("_HANDLER")
    } else {
        appName = os.Getenv("CAV_APP_AGENT_INSTANCE")
    }
    var final []byte
    if os.Getenv("CAV_APP_AGENT_TAGS") != "" {
        tagkey = os.Getenv("CAV_APP_AGENT_TAGS")
        tagfield := strings.Split(tagkey, ",")
        //var myMap map[string]string
        myMap := make(map[string]string)
        //final = "Technolgy:go"
        myMap["Technolgy"] = "go"
        for _, i := range tagfield {
            key := strings.Split(i, "=")
            namespace := strings.Split(key[1], ":")
            tagkeys := namespace[1]
            //fmt.Println(tagkeys)
            var tagvalue string
            switch namespace[0] {
            case "aws":
                tagvalue = findtagvalue(tagkeys, ctx)
            case "inline":
                tagvalue = tagkeys
            case "env":
                tagvalue = os.Getenv(tagkeys)
            default:
                log.Println("Unknow tag")
            }
            if key[0] != "" && tagvalue != "" {
                myMap[key[0]] = tagvalue
            }

        }

        final, err = json.Marshal(myMap)
        if err != nil {
            log.Printf("Error: %s", err.Error())
        }

    }

    Appinfo.Tags = "tierName=" + Tier + ";ndAppServerHost=" + server + ";appName=" + appName + ";tags=" + string(final)
    //fmt.Println(string(final))
    var apiReqLen = C.int(len(Appinfo.APIReqId))
    var awsReqLen = C.int(len(Appinfo.AWSReqId))
    var funcNameLen = C.int(len(Appinfo.FuncName))
    var tagslength = C.int(len(Appinfo.Tags))
    var agentType = C.short(4)
    var messageType = C.short(0)
    if msgType != 0 {
        messageType = msgType
    }
    
    len := C.WrapHeaderlen(apiReqLen, awsReqLen, funcNameLen, tagslength, agentType, messageType)
    len = len + apiReqLen + awsReqLen + funcNameLen + tagslength

    return len
}

func findtagvalue(k string, ctx context.Context) string {

    mySession := session.Must(session.NewSession())
    r := resourcegroupstaggingapi.New(mySession, aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))
    lc, _ := lambdacontext.FromContext(ctx)
    var paginationToken string = ""
    var in *resourcegroupstaggingapi.GetResourcesInput
    var out *resourcegroupstaggingapi.GetResourcesOutput
    var err error

    if len(paginationToken) == 0 {
        in = &resourcegroupstaggingapi.GetResourcesInput{
            ResourcesPerPage: aws.Int64(50),
        }
        out, err = r.GetResources(in)
        if err != nil {
            log.Println(err)
        }
    } else {
        in = &resourcegroupstaggingapi.GetResourcesInput{
            ResourcesPerPage: aws.Int64(50),
            PaginationToken:  &paginationToken,
        }
    }
    out, err = r.GetResources(in)
    if err != nil {
        log.Println(err)
    }
    for _, resource := range out.ResourceTagMappingList {

        if *resource.ResourceARN == lc.InvokedFunctionArn {
            var myMap map[string]string
            for _, a := range resource.Tags {

                data, err := json.Marshal(a)
                if err != nil {
                    log.Println("error ", err)
                }
                json.Unmarshal(data, &myMap)
            }
           
            if myMap["Key"] == k {
                a := myMap["Value"]
                return string(a)
            }

        }

    }
    return ""
}

func sendperameter(buf []byte, lenght C.int, nums ...string) C.int {
    for _, num := range nums {
        a := C.CString(num)
        lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, lenght, C.int(len(num)))
        C.free(unsafe.Pointer(a))
    }
    return lenght
}

type aiRecord struct {
    conn net.Conn
}

var aiRecObj *aiRecord = nil
var err error

func NewAIRecord() *aiRecord {
    r := aiRecord{}
    ipAddeers := os.Getenv("CAV_APP_AGENT_PROXYIP")
    port := os.Getenv("CAV_APP_AGENT_PROXYPORT")
    r.conn, err = net.Dial("udp", ipAddeers+":"+port)
    if err != nil {
        log.Printf("Not able to make connection with CAV_APP_AGENT %v", err)
    }
    // fmt.Println("conn value", r.conn)
    return &r
}

func CloseUDP() {
    aiRecObj.conn.Close()
    // log.Println("close")
}

func UDPConnection() {

    log.Println("udp_call")

    aiRecObj = NewAIRecord()
    // fmt.Println(aiRecObj)

    time.Sleep(800 * time.Millisecond)

}

type NDMsg struct {
    CookieName       string
    DomainName       string
    HeaderInResponse int
    Mpos             int
}

var CkEnable bool
var CkName string
var CkDomainName string
var CkResHeader int = 0
var CkMethodPos int = 0
var Buffer []byte

var Bt_header *string

func ReceiveMessageFromServer() {

    for {

        request := make([]byte, 1024)
        // time.Sleep(time.Millisecond * 200)
        _, _ = aiRecObj.conn.Read(request)

        /*if err != nil {
            log.Println("not able to recived data")

        }*/
        for i := 0; i < len(request)-1; i++ {
            if request[i] != 0 {
                Buffer = append(Buffer, request[i])
            }
        }
        if Closeroutine {
            break
        }
        if Buffer != nil {
            datareceived()
        }
    }

}


func StartTransactionMessage(ctx context.Context, CorrelationHeader string) {

    lenght1 := Headerlen(0,ctx)

    var fp_header1 = "dummy_fp_header"
    btHeaderValue1 := *Bt_header
    

    var fp_header = C.int(len(fp_header1))
    var url = C.int(len(Url_path))
    var btHeaderValue = C.int(len(btHeaderValue1))
    var ndCookieSet = C.int(len(Ndcookie))
    var nvCookieSet = C.int(len(Nvcookie))
    var correlationHeader = C.int(len(CorrelationHeader))
    var flowpathinstance = C.longlong(0)
    var startTimeFP = C.longlong(0)
    var qTimeMS = C.longlong(0)

    lenght1 = C.StartTransactionlen(fp_header, url, btHeaderValue, ndCookieSet, nvCookieSet, correlationHeader, flowpathinstance, qTimeMS, startTimeFP, lenght1)
    lenght1 = lenght1 + fp_header + url + btHeaderValue + ndCookieSet + nvCookieSet + correlationHeader 
    lenght1++
   
    var buf = make([]byte, lenght1)
    lenght := Header(buf)
    lenght = C.StartTransaction((*C.char)(unsafe.Pointer(&buf[0])),lenght)
    lenght = sendperameter(buf, lenght, fp_header1, Url_path, btHeaderValue1, Ndcookie, Nvcookie, CorrelationHeader)
    
    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    log.Println(lenght)
    _, err := aiRecObj.conn.Write(buf)

    log.Println("send data_start")
    if err != nil {
        log.Println("err not null")

    }
    msg := "StartTransactionMessage function called "
    NDNFSendMessage(ctx, msg)

}

func method_entry(ctx context.Context, MethodName string) {
    lenght1 := Headerlen(0,ctx)   
    

    query_string1 := "select * from countries"
    urlParameter1 := ""

    var urlParameter = C.int(len(urlParameter1))
    var methodName = C.int(len(MethodName))
    var query_string = C.int(len(query_string1))
    var mid = C.int(0)
    var flowpathinstance = C.longlong(0)
    var threadId = C.longlong(0)
    var startTime = C.longlong(0)
    
    lenght1 = C.MethodEntryFunctionlen(urlParameter, methodName, query_string, mid, flowpathinstance, threadId, startTime, lenght1)
    lenght1 = lenght1 + urlParameter + methodName + query_string 
    lenght1++
    
    var buf = make([]byte, lenght1)
    lenght := Header(buf)
    lenght = C.MethodEntryFunction((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    lenght = sendperameter(buf, lenght, MethodName, query_string1, urlParameter1)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    log.Println(lenght)
    _, err := aiRecObj.conn.Write(buf)
    log.Println("send data_MEntry")
    if err != nil {
        log.Println(err)

    }
    msg := "method_entry function called"
    NDNFSendMessage(ctx, msg)
}

func method_exit(ctx context.Context, MethodName string, statuscode int) {

    

    backend_header1 := "NA|10.20.0.85|NA|NA|mydb|mysql|NA|NA|NA|root"
    requestNotificationPhase1 := ""

    var statusCode = C.int(statuscode)
    var mid = C.int(0)
    var eventType = C.int(1)
    var isCallout = C.int(1)
    var duration = C.longlong(363)
    var threadId = C.longlong(0)
    var cpuTime = C.longlong(0)
    var flowpathinstance = C.longlong(0)
    var tierCallOutSeqNum = C.longlong(45)
    var endTime = C.longlong(0)
    var methodName = C.int(len(MethodName))
    var backend_header = C.int(len(backend_header1))
    var requestNotificationPhase = C.int(len(requestNotificationPhase1))
    
    lenght1 := Headerlen(0,ctx)
    lenght1 = C.MethodExitFunctionlen(statusCode, mid, eventType, isCallout, duration, threadId, cpuTime, flowpathinstance, tierCallOutSeqNum, endTime, methodName, backend_header, requestNotificationPhase, lenght1)
    lenght1 = lenght1 + methodName + backend_header + requestNotificationPhase
    lenght1++
    
    var buf = make([]byte, lenght1)
    lenght := Header(buf)
    lenght = C.MethodExitFunction((*C.char)(unsafe.Pointer(&buf[0])),lenght)
    lenght = sendperameter(buf, lenght, MethodName, backend_header1, requestNotificationPhase1)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    log.Println(lenght)
    _, err := aiRecObj.conn.Write(buf)
    log.Println("send data_MExit")
    if err != nil {
        log.Println(err)

    }
    msg := "method_exit function called"
    NDNFSendMessage(ctx, msg)

}

func end_business_transaction(ctx context.Context, statuscode int) {
    var statusCode = C.int(statuscode)
    var endTime = C.longlong(0)
    var flowpathinstance = C.longlong(0)
    var cpuTime = C.longlong(0)
    
    lenght1 := Headerlen(0,ctx)
    lenght1 = C.EndTransactionlen(statusCode, endTime, flowpathinstance, cpuTime, lenght1)
    lenght1++
    
    var buf = make([]byte, lenght1)
    lenght := Header(buf)
    lenght = C.EndTransaction((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    
    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    log.Println(lenght)
    _, err := aiRecObj.conn.Write(buf)
    log.Println("send data_end")
    if err != nil {
        log.Println(err)

    }
    msg := "end_business_transaction function called"
    NDNFSendMessage(ctx, msg)

}

func SendReqRespHeder(ctx context.Context, buffer string, Headertype string, statuscode int) {
  
    lenght1 := Headerlen(0,ctx)
    var statusCode = C.int(statuscode)
    var buffer_len = C.int(len(buffer))
    var type_len = C.int(len(Headertype))
    var flowpathinstance = C.longlong(0)
    
    lenght1 = C.ReqRespHeaderlen(statusCode, buffer_len, type_len, flowpathinstance, lenght1)
    lenght1 = lenght1 + buffer_len + type_len 
    lenght1++
    
    var buf = make([]byte, lenght1)
    lenght := Header(buf)
    lenght = C.ReqRespHeader((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    lenght = sendperameter(buf, lenght, buffer, Headertype)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    log.Println(lenght)
    _, err := aiRecObj.conn.Write(buf)
    log.Println("send headerReqResp")
    if err != nil {
        log.Println(err)

    }
    msg := "SendReqRespHeder function called"
    NDNFSendMessage(ctx, msg)
}

func NVCookieMessage(ctx context.Context) {
    len := Headerlen(4,ctx)
    var buf = make([]byte, len)
    _ = Header(buf)
    _, err := aiRecObj.conn.Write(buf)
    if err != nil {
        log.Println("Error in NVCookieMessage", err)

    }

    
    msg := "NVCookieMessage function called"
    NDNFSendMessage(ctx, msg)

}
func NDCookieMessage(ctx context.Context) {
    len := Headerlen(5,ctx)
    var buf = make([]byte, len)
    _ = Header(buf)
    _, err := aiRecObj.conn.Write(buf)

    if err != nil {
        log.Println("Error in NDCookieMessage ", err)

    }

   
    msg := "NDCookieMessage function called"
    NDNFSendMessage(ctx, msg)

}
func NDCookieValue(ctx context.Context) {
    len := Headerlen(6,ctx)
    var buf = make([]byte,len)
    _ = Header(buf)
    _, err := aiRecObj.conn.Write(buf)

    if err != nil {
        log.Println("Error in NDNFCookieMessage ", err)

    }

    msg := "NDCookieValue "
    NDNFSendMessage(ctx, msg)

}

func responsecookies() {
    Ckheader = "Name=" + CkName + ";Value=" + NDCOOKIE_VALUE + ";MaxAge=1000;Path=/;Domain=" + CkDomainName + ";Secure;HTTPOnly;"
    log.Println("cookies value", Ckheader)

}

func datareceived() {
    req := string(Buffer)

    if strings.Contains(req, "NVCookie") {
        c := strings.Split(req, ":")
        log.Println("NVCookie", c[1])
        if c[1] == "" {
            Nvcookie = ""
        }
        Buffer = nil
        Nvcookie = c[1]
    } else if strings.Contains(req, "CookieName") {
        var ndmsg NDMsg
        json.Unmarshal(Buffer, &ndmsg)
        log.Println("NDCookie=", ndmsg)
        if ndmsg.CookieName == "" {
            CkEnable = false
            Ndcookie = ""
        }
        CkResHeader = ndmsg.HeaderInResponse
        CkMethodPos = ndmsg.Mpos
        CkDomainName = ndmsg.DomainName
        CkName = ndmsg.CookieName
        CkEnable = true
        Ndcookie = CkName
        Buffer = nil
    } else if strings.Contains(req, "NDcookieValue") {

        c := strings.Split(req, ":")
        log.Println("value of NDCOOKIE=", c[1])
        NDCOOKIE_VALUE = c[1]
        Buffer = nil
    }
}
func NDNFSendMessage(ctx context.Context, msg string) {
    len1 := Headerlen(7,ctx)
    var buf = make([]byte, len1)
    lenght := Header(buf)
   
    size := C.int(len(msg))
    a := C.CString(msg)

    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, lenght, size)
    _, err := aiRecObj.conn.Write(buf)
    if err != nil {
        log.Println("Error in ND-NFSendMessage", err)

    }

}
