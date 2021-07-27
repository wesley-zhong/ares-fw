package webserver

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"netease.com/core"
)

type AresWebapp struct {
}

func (webApp *AresWebapp) WebAppStart(coreInst *core.Core, addr string) {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		//inner rpc
		var rpcReq core.RpcReq
		if err := c.ShouldBind(&rpcReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": rpcReq})

	})

	r.POST("/:serviceName/:methodName", func(c *gin.Context) {
		//inner restful or single server restful
		var rpcReq core.RpcReq
		rpcReq.ServiceName = c.Params.ByName("serviceName")
		rpcReq.MethodName = c.Params.ByName("methodName")

		aresMethod := coreInst.GetCallFun(rpcReq.ServiceName, rpcReq.MethodName)
		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		if err := c.ShouldBind(realParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})

	r.POST("/game/:serviceName/:methodName", func(c *gin.Context) {
		//by proxy such as nginx
		var rpcReq core.RpcReq
		rpcReq.ServiceName = c.Params.ByName("serviceName")
		rpcReq.MethodName = c.Params.ByName("methodName")
		var body string
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rpcReq.PayLoad = body
		c.JSON(200, gin.H{"data": rpcReq})
	})
	log.Info("start webserver:", addr)
	r.Run(addr)
}
