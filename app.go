package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/werifu/touhou_bot/controller"
	"github.com/werifu/touhou_bot/util"
	"strconv"
)
//go:generate cqcfg -c .
// cqp: 名称: thbot
// cqp: 版本: 1.0.1:1
// cqp: 作者: Werifu
// cqp: 简介: 供给风蓝东方project群使用
const(
	th_group = 280569556
	my_qq = 1363195380
)
func main() { /*此处应当留空*/}

func init() {
	cqp.AppID = "me.cqp.werifu.thbot" // TODO: 修改为这个插件的ID
	cqp.Enable = controller.OnlyPushHY
	cqp.GroupMsg = GroupCMD
	cqp.PrivateMsg = controller.Repeat
}



func GroupCMD(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32{
	var str string

	if util.HasImg(msg){	//符合
		util.AddImg(fromGroup, fromQQ)
	}
	if util.HasKusa(msg){
		util.AddKusa(fromGroup, fromQQ)
	}
	if util.IsTeach(msg){
		controller.Teach(fromGroup,fromQQ,msg)
	}
	if msg[0] == '.'{
		cmd := msg[1:]
		//cqp.SendGroupMsg(fromGroup, cmd)
		if util.Is114514(cmd){
			cqp.SendGroupMsg(fromGroup, "这么臭的命令有什么执行的必要吗")
			return 0
		}else if cmd == "help" { //打印指令集
			controller.Help(fromGroup)
		}else if cmd == "at" { //at那些id带@一下的
			controller.Send_atAll(fromGroup)
		}else if cmd == "rollth" { //随机正作
			thyes, thno := controller.Roll_th()
			str = "今日宜：" + thyes + "\n今日不宜：" + thno
			cqp.SendGroupMsg(fromGroup, str)
		}else if cmd == "jrxy" {	//今日性欲
			cqp.SendGroupMsg(fromGroup, "[CQ:at,qq="+strconv.Itoa(int(fromQQ))+"] \n今日性欲："+strconv.Itoa(controller.Roll_xy(114,514,fromQQ))+"\n(114~514)")
		}else if cmd == "test" {	//测试代码
			cqp.SendPrivateMsg(my_qq, "测试用")
		}else if cmd == "img"{
			controller.SendImgNum(fromGroup, fromQQ)
		}else if cmd == "kusa"{
			controller.SendKusaNum(fromGroup,fromQQ)
		}else if cmd == "_clear" && fromQQ == my_qq{
			util.ClearDaily(fromGroup)
		}else if util.IsAtCMD(cmd){	//响应at命令
			cqp.SendGroupMsg(fromGroup, controller.At_times(cmd))
		}else{
			cqp.SendGroupMsg(fromGroup, "Unknown command")
		}
	}

	//自定义条件反射
	ref := util.GetRefRemsg(msg)
	if ref != ""{
		cqp.SendGroupMsg(fromGroup, ref)
	}

	return 0
}

