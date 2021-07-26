package tcpserver

import (
	"bytes"
	"encoding/binary"

	log "github.com/sirupsen/logrus"
	"netease.com/core"

	"github.com/panjf2000/gnet"
)

type TcpServer struct {
	*gnet.EventServer
	Core *core.Core
}

func (es *TcpServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	log.Info("========= server init", server)
	return gnet.None
}

func (es *TcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Info("new connectd ", c.RemoteAddr().String())
	return
}

func (es *TcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var msgId int16
	bytebuffer := bytes.NewBuffer(frame)
	binary.Read(bytebuffer, binary.BigEndian, &msgId)
	var checkSum uint32
	binary.Read(bytebuffer, binary.BigEndian, &checkSum)
	if msgId == 2 {
		pong := int16(1)
		sendBytebuffer := bytes.NewBuffer([]byte{})
		binary.Write(sendBytebuffer, binary.BigEndian, &pong)
		out = sendBytebuffer.Bytes()
		return
	}
	log.Info("msgId=  ", msgId, " len = ", bytebuffer.Len(), "receive check sum = ", checkSum)
	msgIdMethod := es.Core.MsgIdMethodMap[int(msgId)]
	if msgIdMethod == nil {
		log.Error("not foud msgId=", msgId)
		return
	}
	msgIdMethod.Call(bytebuffer)
	return
}

// OnShutdown fires when the server is being shut down, it is called right after
// all event-loops and connections are closed.
func (es *TcpServer) OnShutdown(server gnet.Server) {
	log.Info("XXXXXconnect closed", server)
}

func (es *TcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Info("XCXXX  socket closed ", c.RemoteAddr(), "   msg = ", err.Error())
	return gnet.None

}
func (es *TcpServer) Init(addr string) {
	ec := gnet.EncoderConfig{binary.BigEndian, 2, 2, true}
	dc := gnet.DecoderConfig{binary.BigEndian, 0, 2, 0, 2}
	err := gnet.Serve(es, addr, gnet.WithMulticore(true), gnet.WithCodec(gnet.NewLengthFieldBasedFrameCodec(ec, dc)))
	log.Fatal(err)
}
