package controller

import (
	"google.golang.org/protobuf/proto"
	"netease.com/core"
)

var BattleRoom = &BattleRoomWrap{}

func Init() {
	BattleRoom.registerMe()
	RegisterController()
}

/**
msgId controller register
*/
func RegisterCallByMsgId(msgId int, msg proto.Message, callFun core.MsgIdMethodFuc) {
	core.Core.RegisterMsgIdCall(msgId, msg, callFun)
}

/**
rpc controller register
*/

func RegisterController() {
	//------------userLogin controller-----------------
	webc := &UserLogin{}
	core.Core.RegisterController(webc)
}
