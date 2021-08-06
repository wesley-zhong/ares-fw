# ares-fw 
ares-fw(ares-framework) include web  request(http) and tcp request

# web request
integrate [gin](https://github.com/gin-gonic/gin)
 include restful request 
 ```
 addr/{serviceName}/{methodName}  
 ```
or rpc request

```
 requestBody ={
     "serviceName":"service",
     "methodName":"method",
     "payLoad":"body"
 }

```

 # tcp request
 integrate  [gnet](https://github.com/panjf2000/gnet)
  called method by registed msgId