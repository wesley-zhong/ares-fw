package controller

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"netease.com/core"
	"netease.com/gen"
)

// func callEnterRoomReq(enterRoomReq *gen.EnterRoomReq) {
// 	log.Info("------------ enterRoomReq ", enterRoomReq.RoomId)

// }

func funtest(param *interface{}) {
	// /req := new(gen.EnterRoomReq)
	//proto.Unmarshal( byte{"aaaa"}, req)

}

func callEnterRoomReq(message *proto.Message) {
	enterRoomMsg := (*message).(*gen.EnterRoomReq)
	log.Info("enterRoomMsg= ", enterRoomMsg.RoomId, " addr= ", message)
}

func Register(core *core.Core) {
	entroomType := new(gen.EnterRoomReq)
	core.RegisterMsgIdCall(int(gen.MsgId_ENTER_ROOM_REQ), entroomType, callEnterRoomReq)
}