package util

import (
	"regexp"
)

//是否恶臭命令
func Is114514(cmd string) bool{
	m,_ := regexp.MatchString("^114514.|.114514.|.114514$", cmd);
	if m {
		return true
	}else {
		return false
	}
}

//是否.[CQ:at, ]命令
func IsAtCMD(cmd string) bool{
	m, _ := regexp.MatchString("[CQ:at,qq=[1-9][0-9]{5,14}][ ]?\\*[0-9]*$" , cmd)
	if m {
		return true
	}else{
		return false
	}
}

func HasImg(msg string) bool{
	if m,_ := regexp.MatchString("\\[CQ:image,file=[0-9a-zA-Z]{20,60}\\.(jpg|png|gif)",msg);m {
		return true
	}else{
		return false
	}
}

func HasKusa(msg string) bool{
	if m,_ := regexp.MatchString("草",msg);m {
		return true
	}else{
		return false
	}
}

func IsTeach(str string) bool {
	reg := regexp.MustCompile(`[^ ]{1,30}`)
	cmds := reg.FindAllString(str, -1)
	if len(cmds) < 4{
		return false
	}
	if cmds[0] == "def" && cmds[2] == "return"{
		return true
	}else {
		return false
	}
}
