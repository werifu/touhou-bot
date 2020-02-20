package gotest

import (
	"github.com/werifu/touhou_bot/util"
	"io/ioutil"
	"testing"
)

const text_path = "D:\\werifu\\CQA-xiaoi\\酷Q Air\\dev\\me.cqp.werifu.thbot\\text\\"

func Test_readfile(t *testing.T){
	cmds, err := ioutil.ReadFile("../text/cmds.txt")
	if err != nil{
		t.Error(err)
		return
	}
	t.Log(string(cmds))
}

func Test_config(t *testing.T){
	a := util.GetGroup()
	t.Log(a)
}