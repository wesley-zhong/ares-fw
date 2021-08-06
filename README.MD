ares-fw (ares-framework) include web  request(http) and tcp request

# web request
 include restful request and rpc request
 You can call the http request by  addr/{serviceName}/{methodName}  or rpc call requestBody ={
     "serviceName":"service",
     "methodName":"method",
     "payLoad":"body"
 }

 # tcp request
  called method by registed msgId