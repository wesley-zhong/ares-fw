package core

import (
	"bytes"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"

	"google.golang.org/protobuf/proto"
)

type RpcReq struct {
	ServiceName string
	MethodName  string
	PayLoad     string
}

type AresMethod struct {
	MethodName string
	CallFun    reflect.Method
	ParamsType reflect.Type
	Param1     interface{}
}

func (aresMethod *AresMethod) Invoke(param interface{}) []reflect.Value {
	reValue := reflect.ValueOf(param)
	parmams := []reflect.Value{reflect.ValueOf(aresMethod.Param1), reValue}
	return aresMethod.CallFun.Func.Call(parmams)
}

type ServiceMethods struct {
	ServiceName    string
	AresMethodsMap map[string]*AresMethod
}

type MsgIdMethodFuc func(msg *proto.Message)

type MsgIdMethod struct {
	msgId   int
	param   proto.Message
	callFun MsgIdMethodFuc
}

func (msgIdMethod *MsgIdMethod) Call(bytebuffer *bytes.Buffer) {
	messageParam := msgIdMethod.param
	err := proto.Unmarshal(bytebuffer.Bytes(), messageParam)
	if err != nil {
		log.Error("proto parse failed", err)
		return
	}
	msgIdMethod.callFun(&messageParam)
}

type CoreWrap struct {
	ServiceMthodsMap map[string]*ServiceMethods
	MsgIdMethodMap   map[int]*MsgIdMethod
}

var Core = CoreWrap{}

func Init() {
	Core.Init()
}

func (core *CoreWrap) Init() {
	core.ServiceMthodsMap = map[string]*ServiceMethods{}
	core.MsgIdMethodMap = map[int]*MsgIdMethod{}
}

func (core *CoreWrap) RegisterMsgIdCall(msgId int, msg proto.Message, callFun MsgIdMethodFuc) {
	msgIdMethod := &MsgIdMethod{msgId, msg, callFun}
	core.MsgIdMethodMap[msgId] = msgIdMethod
	log.Info("------register msgId =", msgId)
}

func (core *CoreWrap) RegisterController(contoller interface{}) {
	rtype := reflect.TypeOf(contoller)
	methods := rtype.NumMethod()
	log.Info("============register service =", rtype.Elem().Name(), "  methods count = ", methods, "  ============== begin")
	for i := 0; i < methods; i++ {
		method := rtype.Method(i)
		var serviceName = strings.ToLower(rtype.Elem().Name())
		var methodName = strings.ToLower(method.Name)
		log.Info("---------- method_name = ", methodName, " service_name = "+serviceName)
		var val *ServiceMethods
		var ok bool
		if val, ok = core.ServiceMthodsMap[serviceName]; !ok {
			val = &ServiceMethods{}
			core.ServiceMthodsMap[serviceName] = val
		}
		aresMethod := AresMethod{}
		aresMethod.CallFun = method
		aresMethod.Param1 = contoller
		aresMethod.MethodName = methodName
		//method params
		mt := method.Type
		numIn := mt.NumIn()
		for j := 0; j < numIn; j++ {
			param := mt.In(j)
			aresMethod.ParamsType = param.Elem()
		}
		val.addMethod(&aresMethod)
	}
	log.Info("============register service =", rtype.Elem().Name(), "  methods count = ", methods, "  ============== end")
}

func (srviceMethods *ServiceMethods) addMethod(aresMethod *AresMethod) {
	if srviceMethods.AresMethodsMap == nil {
		srviceMethods.AresMethodsMap = map[string]*AresMethod{}
	}
	srviceMethods.AresMethodsMap[aresMethod.MethodName] = aresMethod
}

func (core *CoreWrap) GetCallFun(serviceName string, methodName string) *AresMethod {
	var val *ServiceMethods
	var ok bool
	if val, ok = core.ServiceMthodsMap[serviceName]; !ok {
		return nil
	}
	if val.AresMethodsMap == nil {
		return nil
	}
	return val.AresMethodsMap[methodName]
}
