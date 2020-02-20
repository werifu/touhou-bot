package gotest

import (
	"github.com/werifu/touhou_bot/util"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func Test_114514(t *testing.T) {
	str := ".114514"
	if m, e := regexp.MatchString("^114514.|.114514.|.114514$", str); m == false || e != nil{
		t.Error("恶臭命令测试未通过")
	}else{
		t.Log("恶臭命令测试通过了")
	}
}

func Test_at(t *testing.T){
	str := "[CQ:at,qq=1000380]*1000"
	if m, e := regexp.MatchString("[CQ:at,qq=[1-9][0-9]{5,14}][ ]?\\*[0-9]*$", str);m ==false || e != nil{
		t.Error("未响应at")
	}else{
		t.Log("响应到at")
	}
}
func Test_at_times(t *testing.T){
	cmd := "[CQ:at,qq=1000380]*20"
	remsg, times_s := strings.Split(cmd, "*")[0], strings.Split(cmd,"*")[1]
	times, err := strconv.Atoi(times_s)
	if err != nil{
		t.Error("字符转数字错了（次数）")
		//return "打错还@锤子呢"
	}
	//把次数后的切了，保留@的内容
	once := remsg
	for i:=1;i < times;i++{
		remsg += "\n"+once
	}

	if remsg != "[CQ:at,qq=1000380]\n[CQ:at,qq=1000380]\n[CQ:at,qq=1000380]"{
		t.Error("at得不对："+remsg)
	} else{
		t.Log("at了好几次,通过")
	}
}
func Test_get_image(t *testing.T){
	msg := "[CQ:image,file=3135EFFF27E0128080176C64846554BE.gif]"

	if m,_ := regexp.MatchString("\\[CQ:image,file=[0-9a-zA-Z]{20,60}\\.(jpg|png|gif)",msg);m{
		reg,_ := regexp.Compile("=[0-9a-zA-Z]{20,60}\\.(jpg|png|gif)")
		file := reg.Find([]byte(msg))
		t.Log(string(file[1:]))
	}else{
		t.Error("匹配不到图片")
	}
}

func Test_null_struct(t *testing.T){
	if (util.SpeakInfo{}.QQ) == 0{
		t.Log("空结构值0")
	}else{
		t.Error("空结构值不为0")
	}
}

func Test_Kusa(t *testing.T){
	msg := "大草原"
	if m,_ := regexp.MatchString("草", msg);m{
		t.Log("草生")
	}else{
		t.Error("草未生")
	}
}

func Test_Teach(t *testing.T){
	if util.IsTeach("def 草 return 草"){
		t.Log("定义通过")
	}else{
		t.Error("失败")
	}
}
