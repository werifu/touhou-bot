package gotest

import (
	"bufio"
	"github.com/werifu/touhou_bot/util"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

//打开文件的测试
func Test_open(t *testing.T) {
	num := util.Roll(0,12)

	gamesList, err := os.Open("../text/thgames.txt")
	if err != nil{
		log.Fatal("Can't open the file")
	}
	scanner := bufio.NewScanner(gamesList)
	line := 1
	for scanner.Scan(){
		if line == num{
			t.Log(scanner.Text())
			return
		}
		line++
	}
	err = gamesList.Close()
	if err != nil{
		log.Fatal("Can't close the file 'thgames.txt'")
	}
	return
}

func Test_randByDay(t *testing.T){
	y, n := util.RollByDay(1,12)
	y2, n2 := util.RollByDay(1,12)
	if y == y2 && n == n2 {
		t.Log(y,n)
		t.Log("今日占卜通过")
	}else{
		t.Log("今日占卜未通过")
	}

}


func Test_rollXY(t *testing.T){
	min,max := 114, 514
	fromQQ := 1363135380
	now := time.Now()
	daySeed := now.Year() * 10000 + int(now.Month()) *100 + now.Day() - int(fromQQ/100)
	rand.Seed(int64(daySeed))		//今日适合
	rp := rand.Intn(max - min)
	rp = rp + min
	t.Log("今日性欲为", rp)
	return
}

func Test_rollByDay(t *testing.T) {
	a, b := util.RollByDay(114, 514)
	t.Log(a, b)
	return
}

func Roll_th() (string, string) {
	//每天适宜和不适宜的飞机

	numYse, numNo := util.RollByDay(1,12)
	gamesList, err := os.Open(text_path+"thgames.txt")
	if err != nil{
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
func Test_TH(t *testing.T){
	a, b := Roll_th()
	t.Log(a, b)
}

