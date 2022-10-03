package index

/*
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "ndlb_encode.h"
int WrapHeader(char *s,int apiReqLen,int awsReqLen,int funcNameLen,int tagslength,short agentType,short messageType){
    int len = 0;
    wrapheader_t wrapHeader;
    memcpy(s, "^",1);
    len += 1;
    
    wrapHeader.wrapheadervar.apiReqLen = apiReqLen ;
    wrapHeader.wrapheadervar.awsReqLen = awsReqLen ;
    wrapHeader.wrapheadervar.funcNameLen = funcNameLen ;
    wrapHeader.wrapheadervar.tagslength = tagslength ;
    wrapHeader.wrapheadervar.agentType = agentType ;
    wrapHeader.wrapheadervar.messageType = messageType;
    wrapHeader.wrapheadervar.whLen = sizeof(wrapheadervar_t)+wrapHeader.wrapheadervar.awsReqLen+wrapHeader.wrapheadervar.apiReqLen+wrapHeader.wrapheadervar.funcNameLen+wrapHeader.wrapheadervar.tagslength+1;
    memcpy(s + len , (char *)&(wrapHeader.wrapheadervar), sizeof(wrapHeader.wrapheadervar));
    len += sizeof(wrapheadervar_t);
  return len;
}
int ValueStore(char *s,char *value,int len,int number)
{
      memcpy(s + len, value ,number);
      len += number;
      return len;
}
int StartTransaction(char *s,int fp_header,int url,int btHeaderValue,int ndCookieSet,int nvCookieSet,int correlationHeader,long long flowpathinstance,long qTimeMS,long long startTimeFP,int len)
{
    transactionStart_t node;
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    
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
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.transactionStartVar), sizeof(node.transactionStartVar));
    len += sizeof(node.transactionStartVar) ;
    return len;
}
int MethodEntryFunction(char *s,int urlParameter,int methodName,int query_string,int mid,long long flowpathinstance,long threadId,long long startTime,int len)
{
    msgHdr_t msgHdr;
    MethodEntry_t node;
    memcpy(s+len, "^",1);
    len += 1;
    
    node.MethodEntryVar.methodName = methodName ;
    node.MethodEntryVar.threadId= threadId;
    node.MethodEntryVar.query_string = query_string ;
    node.MethodEntryVar.urlParameter = urlParameter ;
    node.MethodEntryVar.mid = mid ;
    node.MethodEntryVar.startTime = startTime ;
    node.MethodEntryVar.flowpathinstance = flowpathinstance;
    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof(MethodEntryVar_t) + msgHdr.header_len + node.MethodEntryVar.methodName +
    node.MethodEntryVar.query_string+ node.MethodEntryVar.urlParameter + 3;
    msgHdr.msg_type = 0;
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.MethodEntryVar), sizeof(MethodEntryVar_t));
    len += sizeof(MethodEntryVar_t);
    return len;
}
int MethodExitFunction(char *s,int statusCode,int mid,int eventType,int isCallout,long duration,long threadId,long long cpuTime,long long flowpathinstance,long long tierCallOutSeqNum,long long endTime,int methodName,int backend_header,int requestNotificationPhase,int len)
{
    MethodExit_t node;
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    node.MethodExitVar.statusCode = statusCode ;
    node.MethodExitVar.mid = mid ;
    node.MethodExitVar.eventType = eventType;
    node.MethodExitVar.isCallout = isCallout;
    node.MethodExitVar.duration = duration;
    node.MethodExitVar.threadId = threadId;
    node.MethodExitVar.cpuTime = cpuTime;
    node.MethodExitVar.tierCallOutSeqNum = tierCallOutSeqNum ;
    node.MethodExitVar.endTime = endTime ;
    node.MethodExitVar.methodName = methodName ;
    node.MethodExitVar.backend_header = backend_header;
    node.MethodExitVar.requestNotificationPhase = requestNotificationPhase ;
    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof(MethodExitVar_t)+ msgHdr.header_len + node.MethodExitVar.methodName + node.MethodExitVar.backend_header +
    node.MethodExitVar.requestNotificationPhase + 3;
    msgHdr.msg_type = 1;
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.MethodExitVar), sizeof(node.MethodExitVar));
    len += sizeof(MethodExitVar_t);
    return len;
}
int EndTransaction(char *s,int statuscode,long long endTime,long long flowpathinstance,long long cpuTime ,int len)
{
    transactionEnd_t transactionEnd;
    int N=sizeof (transactionEnd_t);
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    
    transactionEnd.statuscode = statuscode ;
    transactionEnd.endTime = endTime ;
    transactionEnd.flowpathinstance = flowpathinstance;
    transactionEnd.cpuTime = cpuTime;
   
    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = msgHdr.header_len + N + 3;
    msgHdr.msg_type = 3;
    memcpy(s + len , (char *)&(msgHdr), sizeof(msgHdr_t));
    len += sizeof(msgHdr_t);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (void *)&transactionEnd, N);
    len += N;
    return len;
}
int ReqRespHeader(char *s,int statuscode,int buffer_len,int type_len,long long flowpathinstance,int len)
{
    transactionEncodeHttp_t node;
    node.transactionEncodeVarHttp.statuscode = statuscode;
    node.transactionEncodeVarHttp.buffer_len = buffer_len;
    node.transactionEncodeVarHttp.type_len = type_len;
    node.transactionEncodeVarHttp.flowpathinstance = flowpathinstance;
    
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    
    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = msgHdr.header_len + sizeof(transactionEncodeVarHttp_t)  + node.transactionEncodeVarHttp.type_len + 
    node.transactionEncodeVarHttp.buffer_len + 3;
    msgHdr.msg_type = 6;
    
    memcpy(s + len , (char *)&(msgHdr), sizeof(msgHdr_t));
    len += sizeof(msgHdr_t);
    
    memcpy(s + len , "|",1);
    len += 1;
    
     memcpy(s + len, (char *)&(node.transactionEncodeVarHttp), sizeof(node.transactionEncodeVarHttp));
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
    "fmt"
    "log"
    "net"
    "time"
    "github.com/aws/aws-lambda-go/lambdacontext"
    "unsafe"
    "os"
    "context"
    
)



func Header(buf []byte,ctx context.Context) C.int {

    var funcName string
    var awsReqId string
    if lambdacontext.FunctionName != "" {
        funcName = lambdacontext.FunctionName
        log.Printf("FUNCTION NAME1: %s", lambdacontext.FunctionName)
    }else {
        funcName = "main_test1"
    }
    var apiReqId = "akkdnfjnflffk"
    lc, _ := lambdacontext.FromContext(ctx)
    if lc.AwsRequestID != "" {
        awsReqId = lc.AwsRequestID
        log.Printf("REQUEST ID: %s", lc.AwsRequestID)
    }else {
         awsReqId =  "asfddhgfjhgj"
    }
    var Tier,server,appName string
    if os.Getenv("CAV_APP_AGENT_TIER") == "" {
        Tier = "default"
    }else {
        Tier = os.Getenv("CAV_APP_AGENT_TIER")
    }
    if os.Getenv("CAV_APP_AGENT_SERVER") == "" {
        server = "default"
    }else {
        server = os.Getenv("CAV_APP_AGENT_SERVER")
    }
    if os.Getenv("CAV_APP_AGENT_INSTANCE") == "" {
        appName =os.Getenv("_HANDLER")
    }else {
        appName = os.Getenv("CAV_APP_AGENT_INSTANCE")
    }
    var tags = "tierName="+Tier+";ndAppServerHost="+server+";appName=" + appName

    var apiReqLen = C.int(len(apiReqId))
    var awsReqLen = C.int(len(awsReqId))
    var funcNameLen = C.int(len(funcName))
    var tagslength = C.int(len(tags))
    var agentType = C.short(4)
    var messageType = C.short(0)

    len := C.WrapHeader((*C.char)(unsafe.Pointer(&buf[0])), apiReqLen, awsReqLen, funcNameLen, tagslength, agentType, messageType)
    
    a := C.CString(apiReqId)
    b := C.CString(awsReqId)
    c := C.CString(funcName)
    d := C.CString(tags)
    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))
    defer C.free(unsafe.Pointer(d))

    len = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, len, apiReqLen)
    len = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), b, len, awsReqLen)
    len = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), c, len, funcNameLen)
    len = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), d, len, tagslength)

    return len
}


type aiRecord struct {
    conn net.Conn
}

var aiRecObj *aiRecord = nil
var err error

func NewAIRecord() *aiRecord {
    r := aiRecord{}
    ipAddeers := os.Getenv("CAV_APP_AGENT_PROXYIP")
    port      := os.Getenv("CAV_APP_AGENT_PROXYPORT")
    r.conn, err = net.Dial("udp", ipAddeers+":"+port)
    if err != nil {
        fmt.Printf("Not able to make connection with CAV_APP_AGENT %v", err)
    }
    fmt.Println("conn value", r.conn)
    return &r
}

func CloseUDP() {
    aiRecObj.conn.Close()
    fmt.Println("close")
}

func UDPConnection() {

    fmt.Println("udp_call")

    aiRecObj = NewAIRecord()
    fmt.Println(aiRecObj)

    time.Sleep(1 * time.Second)

}

func ReceiveMessageFromServer() {
    request := make([]byte, 1024)
    a, err := aiRecObj.conn.Read(request)
    ar b []byte
    for i := 0; i < len(request)-1; i++ {
            if request[i] != 0 {
                b = append(b, request[i])
            }
     }
    req := string(b)
    fmt.Println("request=", req, err)
    a := strings.Split(req, ":")
    log.Println("NVCookie", a[1])
   // return a[1]
   // log.Println("request=", string(request), a)
}

/*func generate_bt() {
    id := uuid.New().String()
    var i big.Int
    i.SetString(strings.Replace(id, "-", "", 4), 16)
    return i.String()
}*/

func StartTransactionMessage(ctx context.Context,bt_name string, CorrelationHeader string) {

    var buf = make([]byte, 1024)
    lenght := Header(buf, ctx)
    var fp_header1 = "dummy_fp_header"
    var url1 = bt_name
    btHeaderValue1 := "dummy_btHeaderValue"
    ndCookieSet1 := ""
    nvCookieSet1 := ""

    var fp_header = C.int(len(fp_header1))
    var url = C.int(len(url1))
    var btHeaderValue = C.int(len(btHeaderValue1))
    var ndCookieSet = C.int(len(ndCookieSet1))
    var nvCookieSet = C.int(len(nvCookieSet1))
    var correlationHeader = C.int(len(CorrelationHeader))
    var flowpathinstance = C.longlong(0)
    var startTimeFP = C.longlong(0)
    var qTimeMS = C.long(0)

    lenght = C.StartTransaction((*C.char)(unsafe.Pointer(&buf[0])), fp_header, url, btHeaderValue, ndCookieSet, nvCookieSet, correlationHeader, flowpathinstance, qTimeMS, startTimeFP, lenght)

    a := C.CString(fp_header1)
    b := C.CString(url1)
    c := C.CString(btHeaderValue1)
    d := C.CString(ndCookieSet1)
    e := C.CString(nvCookieSet1)
    f := C.CString(CorrelationHeader)
    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))
    defer C.free(unsafe.Pointer(d))
    defer C.free(unsafe.Pointer(e))
    defer C.free(unsafe.Pointer(f))

    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, lenght, fp_header)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), b, lenght, url)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), c, lenght, btHeaderValue)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), d, lenght, ndCookieSet)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), e, lenght, nvCookieSet)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), f, lenght, correlationHeader)
    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    _, err := aiRecObj.conn.Write(buf)

    fmt.Println("send data_start")
    if err != nil {
        log.Println("not able to send data")

    }

    

}

func method_entry(ctx context.Context,MethodName string) {
    var buf = make([]byte, 1024)
    lenght := Header(buf, ctx)

    
    query_string1 := "select * from countries"
    urlParameter1 := ""

    var urlParameter = C.int(len(urlParameter1))
    var methodName = C.int(len(MethodName))
    var query_string = C.int(len(query_string1))
    var mid = C.int(0)
    var flowpathinstance = C.longlong(0)
    var threadId = C.long(0)
    var startTime = C.longlong(0)

    lenght = C.MethodEntryFunction((*C.char)(unsafe.Pointer(&buf[0])), urlParameter, methodName, query_string, mid, flowpathinstance, threadId, startTime, lenght)

    a := C.CString(MethodName)
    b := C.CString(query_string1)
    c := C.CString(urlParameter1)
    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))

    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, lenght, methodName)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), b, lenght, query_string)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), c, lenght, urlParameter)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)

    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send data_MEntry")
    if err != nil {
        log.Println(err)

    }
}


func method_exit(ctx context.Context,MethodName string,statuscode int) {

    var buf = make([]byte, 1024)
    lenght := Header(buf, ctx)

    
    backend_header1 := "NA|10.20.0.85|NA|NA|mydb|mysql|NA|NA|NA|root"
    requestNotificationPhase1 := ""

    var statusCode = C.int(statuscode)
    var mid = C.int(0)
    var eventType = C.int(1)
    var isCallout = C.int(1)
    var duration = C.long(363)
    var threadId = C.long(0)
    var cpuTime = C.longlong(0)
    var flowpathinstance = C.longlong(0)
    var tierCallOutSeqNum = C.longlong(45)
    var endTime = C.longlong(0)
    var methodName = C.int(len(MethodName))
    var backend_header = C.int(len(backend_header1))
    var requestNotificationPhase = C.int(len(requestNotificationPhase1))

    lenght = C.MethodExitFunction((*C.char)(unsafe.Pointer(&buf[0])),statusCode, mid, eventType, isCallout, duration, threadId, cpuTime, flowpathinstance, tierCallOutSeqNum, endTime, methodName, backend_header, requestNotificationPhase, lenght)
    a := C.CString(MethodName)
    b := C.CString(backend_header1)
    c := C.CString(requestNotificationPhase1)

    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))

    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, lenght, methodName)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), b, lenght, backend_header)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), c, lenght, requestNotificationPhase)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send data_MExit")
    if err != nil {
        log.Println(err)

    }

}

func end_business_transaction(ctx context.Context,statuscode int) {

    var buf = make([]byte, 1024)
    lenght := Header(buf, ctx)

    var statusCode = C.int(statuscode)
    var endTime = C.longlong(0)
    var flowpathinstance = C.longlong(0)
    var cpuTime = C.longlong(0)

    lenght = C.EndTransaction((*C.char)(unsafe.Pointer(&buf[0])), statusCode, endTime, flowpathinstance, cpuTime, lenght)
    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send data_end")
    if err != nil {
        log.Println(err)

    }

}


func SendReqRespHeder(ctx context.Context,buffer string,Headertype string,statuscode int) {
    var buf = make([]byte, 1024)
    lenght := Header(buf, ctx)
    
    var statusCode = C.int(statuscode)
    var buffer_len = C.int(len(buffer))
    var type_len = C.int(len(Headertype))
    var flowpathinstance = C.longlong(0)
   
    lenght = C.ReqRespHeader((*C.char)(unsafe.Pointer(&buf[0])),statusCode,buffer_len,type_len,flowpathinstance,lenght)
     a := C.CString(buffer)
     b := C.CString(Headertype)
    

    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    

    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), a, lenght, buffer_len)
    lenght = C.ValueStore((*C.char)(unsafe.Pointer(&buf[0])), b, lenght, type_len)
    

    C.last((*C.char)(unsafe.Pointer(&buf[0])), lenght)
    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send headerReqResp")
    if err != nil {
        log.Println(err)

     }
}
func NDCookieMessage(ctx context.Context)  {
    var buf = make([]byte, 1024)
    _ = Header(buf, 5, ctx)
    _, err := aiRecObj.conn.Write(buf)
    if err != nil {
        log.Fatal(err)

    }
    ReceiveMessageFromServer()
    fmt.Println("NDCookieMessage recived")
   // return ndValue
}
