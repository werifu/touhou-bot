package controller

import (
	"bufio"
	"fmt"
	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/werifu/touhou_bot/util"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)
const(
	text_path = "./dev/me.cqp.werifu.thbot/text/"
)




func OnlyPushHY() int32{		//定时单推滑音爹
	for{
		now := time.Now()
		if now.Minute() == 0{
			remsg := "现在是系统时间:"+ now.Format("2006-01-02 15:04:05") + "\n让我们一起单推滑音爹"
			cqp.SendGroupMsg(THGroup, remsg)
			if now.Hour() == 0 {
				cqp.SendGroupMsg(THGroup, "今日统计已归零")
				util.ClearDaily(THGroup)
			}
			time.Sleep(time.Minute *59)
		}else{
			time.Sleep(time.Second *1)
		}
	}
	return 0
}

func Send_atAll(fromGroup int64)int32{
	members := cqp.GetGroupMemberList(fromGroup)
	var remsgs []string
	var remsg string
	for _, member := range members {
		if m, _ := regexp.MatchString("^@一下", member.Card); m { //匹配到了
			remsgs = append(remsgs, "[CQ:at,qq="+strconv.FormatInt(member.QQ, 10)+"]")
			remsg = strings.Join(remsgs, "\n")
		}
	}
	cqp.SendGroupMsg(fromGroup, remsg)
	return 0
}

func At_times(cmd string)string {
	remsg, times_s := strings.Split(cmd, "*")[0], strings.Split(cmd,"*")[1]//把命令切成[at]和次数
	times, err := strconv.Atoi(times_s)
	if err != nil{
		return "次数错误"
	}
	if times < 1 || times > 24{
		return "兄弟，一天才24小时"
	}
	once := remsg
	for i:=1;i < times;i++{
		remsg += "\n"+ once
	}
	return remsg
}

func Roll_th() (string, string) {
	//每天适宜和不适宜的飞机

	numYse, numNo := util.RollByDay(1,12)
	gamesList, err := os.Open(text_path+"thgames.txt")
	if err != nil{
		pwd, _ := os.Getwd()
		cqp.SendPrivateMsg(MyQQ, pwd)
		log.Fatal("Can't open the file")
	}
	scanner := bufio.NewScanner(gamesList)
	line := 1
	var yes,no string
	for scanner.Scan(){
		if line == numYse{
			yes = scanner.Text()
		}else if line == numNo{
			no = scanner.Text()
		}
		//如果yes和no都赋上值了
		if yes!="" && no!= ""{
			break
		}
		line++
	}
	err = gamesList.Close()
	if err != nil{
		log.Fatal("Can't close the file 'thgames.txt'")
	}
	return yes,no
}

func Roll_ufo() string{
	var result []string
	var ufo = make(chan string, 1)
	for i:=0;i<3;i++ {
		select {
		case ufo<-"红":
			result = append(result, "红")
		case ufo<-"蓝":
			result = append(result, "蓝")
		case ufo<-"绿":
			result = append(result, "绿")
		}
		<-ufo
	}
	return strings.Join(result, " ")
}


func Roll_xy(min int, max int, fromQQ int64) int{
	now := time.Now()
	daySeed := now.Year() * 10000 + int(now.Month()) *100 + now.Day() - int(fromQQ/100)
	rand.Seed(int64(daySeed))		//今日适合
	rp := rand.Intn(max - min)
	rp = rp + min
	return rp
}

func Ana_img(msg string){
	reg,_ := regexp.Compile("=[0-9a-zA-Z]{20,60}\\.jpg")
	file := reg.Find([]byte(msg))[1:]
	url := cqp.GetImage(string(file))
	cqp.SendPrivateMsg(MyQQ, url)
}

func SendKusaNum(fromGroup, fromQQ int64) int32 {
	person := util.FindByQQ(fromQQ)
	if person == (util.SpeakInfo{}){
		person = util.SpeakInfo{Group:fromGroup, QQ:fromQQ, KusaNum:0, ImgNum:0}
		util.InsertSI(fromGroup, fromQQ)
		cqp.SendGroupMsg(fromGroup, "[CQ:at,qq="+strconv.Itoa(int(fromQQ))+"]\n今日生草数为:0")
	}else{
		cqp.SendGroupMsg(fromGroup, "[CQ:at,qq="+strconv.Itoa(int(fromQQ))+"]\n今日生草数为:"+strconv.Itoa(person.KusaNum))
	}

	return 0
}

func SendImgNum(fromGroup, fromQQ int64) int32 {
	person := util.FindByQQ(fromQQ)
	if person == (util.SpeakInfo{}){
		person = util.SpeakInfo{Group:fromGroup, QQ:fromQQ, KusaNum:0, ImgNum:0}
		util.InsertSI(fromGroup, fromQQ)
		cqp.SendGroupMsg(fromGroup, "[CQ:at,qq="+strconv.Itoa(int(fromQQ))+"]\n今日发图数为:0")
	}else{
		cqp.SendGroupMsg(fromGroup, "[CQ:at,qq="+strconv.Itoa(int(fromQQ))+"]\n今日发图数为:"+strconv.Itoa(person.ImgNum))
	}
	return 0
}

func Help(fromGroup int64) int32 {
	cmds,err := ioutil.ReadFile(text_path+"cmds.txt")
	if err != nil {
		return 0
	}
	cqp.SendGroupMsg(fromGroup, string(cmds))
	return 0
}

//调教功能
func Teach(fromGroup, fromQQ int64, str string) int32 {
	cmd, remsg := util.ParseRef(str)

	newRef := util.Reflection{
		Cmd:        cmd,
		Group:      fromGroup,
		TeacherQQs: []int64{fromQQ},
		TeacherNum: 1,
		ReMsg:      remsg,
	}
	ref := util.FindRef(cmd)

	if ref.Cmd != cmd{		//说明是空的
		util.InsertCmd(newRef)
	}else{
		if util.IsRefTheSame(ref, newRef){	//如果是同一个反射
			if util.IsIn_int64(newRef.TeacherQQs[0], ref.TeacherQQs){	//一个人重复设置，跳过不管
				return 0
			}
			err := util.UpdateRef(newRef)
			if err != nil{
				err.Error()
			}
			if ref.TeacherNum == 3{
				cqp.SendGroupMsg(fromGroup, "*新词条更新完成*")
			}
		}else{	//不能更新（比如同条件不同回复(要覆盖
			err := util.CoverRef(newRef)
			if err != nil{
				err.Error()
			}
			cqp.SendGroupMsg(fromGroup, "*有词条被覆盖*")
		}
	}

	return 0
}

func SleepBan(fromGroup, fromQQ int64) int32 {
	var sleepTime = int64(util.Roll(6*60*60, 8*60*60))
	var hours = sleepTime / 3600
	var min = (sleepTime - (hours*3600))/60
	var sec = sleepTime % 60
	cqp.SetGroupBan(fromGroup, fromQQ, sleepTime)
	cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%s]\n获得%d小时%d分钟%d秒精致睡眠\n晚安！",
		strconv.FormatInt(fromQQ,10), hours, min, sec))
	return 0
}