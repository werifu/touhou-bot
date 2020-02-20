package controller

import (
	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/werifu/touhou_bot/util"
)




func Repeat(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	if util.IsTeach(msg){
		Teach(122,fromQQ,msg)
		cqp.SendPrivateMsg(fromQQ, "过了过了")
	}
	return 0
}