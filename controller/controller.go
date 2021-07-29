package controller

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"netease.com/core"
	"netease.com/gen"
)

// func callEnterRoomReq1(enterRoomReq *gen.EnterRoomReq) {
// 	log.Info("------------ enterRoomReq ", enterRoomReq.RoomId)

// }

func callEnterRoomReq(message *proto.Message) {
	enterRoomMsg := (*message).(*gen.EnterRoomReq)
	log.Info("enterRoomMsg= ", enterRoomMsg.RoomId, " addr= ", message)
}

func Register(core *core.Core) {
	entroomType := new(gen.EnterRoomReq)
	core.RegisterMsgIdCall(int(gen.MsgId_ENTER_ROOM_REQ), entroomType, callEnterRoomReq)
}
