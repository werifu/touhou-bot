package util

import (
	"math/rand"
	"time"
)

//生成范围内随机数
func Roll(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	return randNum
}

func RollByDay(min, max int)(int,int){
	now := time.Now()
	daySeed := now.Year() * 10000 + int(now.Month()) *100 + now.Day()
	rand.Seed(int64(daySeed))		//今日适合
	randNumY := rand.Intn(max - min)
	randNumY = randNumY + min
	rand.Seed(int64(daySeed-now.Year() * 10000))		//今日不适合
	randNumN := rand.Intn(max-min)
	randNumN = randNumN + min
	return randNumY, randNumN
}