package webserver

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"netease.com/core"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type AresWebapp struct {
}

func (webApp *AresWebapp) WebAppStart(addr string) {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		//inner rpc
		var rpcReq core.RpcReq
		if err := c.ShouldBind(&rpcReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		aresMethod := core.Core.GetCallFun(rpcReq.ServiceName, rpcReq.MethodName)
		if aresMethod == nil {
			c.JSON(403, "server or method not found")
			return
		}

		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		json.UnmarshalFromString(rpcReq.PayLoad, &realParam)
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})

	r.POST("/:serviceName/:methodName", func(c *gin.Context) {
		//inner restful or single server restful
		serviceName := strings.ToLower(c.Params.ByName("serviceName"))
		methodName := strings.ToLower(c.Params.ByName("methodName"))
		aresMethod := core.Core.GetCallFun(serviceName, methodName)
		if aresMethod == nil {
			c.JSON(403, "server or method not found")
			return
		}
		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		if err := c.ShouldBind(&realParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})

	r.POST("/game/:serviceName/:methodName", func(c *gin.Context) {
		//by proxy such as nginx

		serviceName := c.Params.ByName("serviceName")
		methodName := c.Params.ByName("methodName")
		aresMethod := core.Core.GetCallFun(serviceName, methodName)
		if aresMethod == nil {
			c.JSON(403, "server or method not found")
			return
		}
		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		if err := c.ShouldBind(&realParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})
	log.Info("start webserver:", addr)
	r.Run(addr)
}
