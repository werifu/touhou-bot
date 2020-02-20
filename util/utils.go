package util

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"regexp"
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

	NoSeed := int64(daySeed - randNumY)
	rand.Seed(NoSeed)		//今日不适合
	randNumN := rand.Intn(max-min)
	randNumN = randNumN + min

	for ;randNumY==randNumN; {
		NoSeed -= int64(randNumY * 19)		//重新选种
		rand.Seed(NoSeed)
		randNumN = rand.Intn(max-min)
		randNumN = randNumN + min
	}
	return randNumY, randNumN
}


func IsIn(one interface{}, all...interface{}) bool {
	for _, b := range all{
		if one == b{
			return true
		}
	}
	return false
}

func ParseRef(str string)(cmd, remsg string){
	reg := regexp.MustCompile(`[^ ]{1,30}`)
	cmds := reg.FindAllString(str, -1)
	return cmds[1],cmds[3]
}


func GetRefRemsg(str string) string{
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SelfDeterminedCmd)

	ref := Reflection{}
	err = collection.Find(bson.M{"cmd":str}).One(&ref)
	if err != nil{
		log.Println(err)
		return ""
	}
	if ref.TeacherNum < 3{
		return ""
	}
	return ref.ReMsg
}