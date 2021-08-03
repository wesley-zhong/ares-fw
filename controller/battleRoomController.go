package controller

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"netease.com/gen"
)

type BattleRoomWrap struct {
}

func (battleRoom *BattleRoomWrap) callEnterRoomReq(message *proto.Message) {
	enterRoomMsg := (*message).(*gen.EnterRoomReq)
	log.Info("enterRoomMsg= ", enterRoomMsg.RoomId, " addr= ", message)
}

func (battleRoom *BattleRoomWrap) registerMe() {
	entroomType := new(gen.EnterRoomReq)
	RegisterCallByMsgId(int(gen.MsgId_ENTER_ROOM_REQ), entroomType, BattleRoom.callEnterRoomReq)
}
