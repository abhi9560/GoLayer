#pragma pack(1)


typedef struct wrapheadervar_t {
int whLen;
int apiReqLen;
int awsReqLen;
int funcNameLen;
int tagslength;
short agentType;
short messageType;

}wrapheadervar_t;

typedef struct wrapheader_t {
wrapheadervar_t wrapheadervar;
char *apiReqId;
char *awsReqId;
char *funcName;
char *tags;
}wrapheader_t;

typedef struct msgHdr_t
{
  int header_len;
  int total_len;
  int msg_type;
}msgHdr_t;

extern msgHdr_t msgHdr;

typedef struct transactionStartVar_t
{ 
  int fp_header;
  int url;
  int btHeaderValue;
  int ndCookieSet;
  int nvCookieSet;
  int correlationHeader;
  long long flowpathinstance;
  long qTimeMS;
  long long startTimeFP;
}transactionStartVar_t;

extern transactionStartVar_t transactionStartVar;

typedef struct transactionStart_t
{ 
  transactionStartVar_t transactionStartVar;
  
  char *fp_header;
  char *url;
  char *btHeaderValue;
  char *ndCookieSet;
  char *nvCookieSet;
  char *correlationHeader;
}transactionStart_t;

extern transactionStart_t transactionStart;



typedef struct transactionEnd_t
{
  int statuscode;
  long long flowpathinstance;
  long long endTime;
  long long cpuTime;
}transactionEnd_t;

extern transactionEnd_t transactionEnd;

typedef struct MethodEntryVar_t
{
  int mid;  

  long long flowpathinstance;
  long threadId;
  long long startTime;

  int methodName;
  int query_string;
  int urlParameter;
}MethodEntryVar_t;

extern MethodEntryVar_t MethodEntryVar;

typedef struct MethodEntry_t
{
  struct MethodEntryVar_t MethodEntryVar;

  char *methodName;
  char *query_string;
  char *urlParameter;
}MethodEntry_t;

extern MethodEntry_t MethodEntry;


typedef struct MethodExitVar_t
{
  int statusCode;
  int mid;
  int eventType;
  int isCallout;
  long threadId;
  long duration;
  long long flowpathinstance;
  long long cpuTime;

  int methodName;
  int backend_header;
  int requestNotificationPhase;
  long long tierCallOutSeqNum;
  long long endTime;
}MethodExitVar_t;

extern MethodExitVar_t MethodExitVar;

typedef struct MethodExit_t
{
  MethodExitVar_t MethodExitVar;

  char *methodName;
  char *backend_header;
  char *requestNotificationPhase;

}MethodExit_t;

extern MethodExit_t MethodExit;

typedef struct transactionEncodeVarHttp_t
{
    int  statuscode;
    int buffer_len;
    int type_len;
    long long flowpathinstance;

}transactionEncodeVarHttp_t;
extern transactionEncodeVarHttp_t transactionEncodeVarHttp;

typedef struct transactionEncodeHttp_t
{
    struct transactionEncodeVarHttp_t transactionEncodeVarHttp;
    char* buffer;
    char* type;
} transactionEncodeHttp_t;
extern transactionEncodeHttp_t transactionEncodeHttp;



msgHdr_t msgHdr;
transactionStartVar_t transactionStartVar;
transactionStart_t transactionStart;
MethodEntryVar_t MethodEntryVar;
MethodEntry_t MethodEntry;
MethodExitVar_t MethodExitVar;
MethodExit_t MethodExit;
transactionEncodeVarHttp_t transactionEncodeVarHttp;
transactionEncodeHttp_t transactionEncodeHttp;



