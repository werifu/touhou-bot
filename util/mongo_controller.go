package util

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type SpeakInfo struct {
	Group int64					`bson:"group"`
	QQ int64					`bson:"qq"`
	KusaNum int					`bson:"kusaNum"`
	ImgNum int					`bson:"imgNum"`
}

//SpeakInfo 发言记录相关内容
const (
	URL = "mongo:27017"
	DBname = "hustth"
	SpeakInfoTable = "speakinfo"
)

//插入词条
func InsertSI(group, qq int64) SpeakInfo{
	session, err := mgo.Dial(URL)
	if err != nil{
		panic(err)
	}
	defer session.Close()
	c := session.DB(DBname).C(SpeakInfoTable)		//切换到info表
	spi := SpeakInfo{Group:group, QQ:qq, KusaNum:0, ImgNum:0}
	err = c.Insert(spi)
	if err != nil{
		panic(err)
	}
	return spi
}


//用qq寻找群友发言记录
func FindByQQ(qq int64) SpeakInfo{
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SpeakInfoTable)

	result := SpeakInfo{}
	err = collection.Find(bson.M{"qq":qq}).One(&result)
	if err != nil {
		return SpeakInfo{}
	}
	return result
}

func AddKusa(group, qq int64){
	session, err := mgo.Dial(URL)
	if err != nil {
		error.Error(err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SpeakInfoTable)

	person := FindByQQ(qq)
	if person == (SpeakInfo{}){
		person = InsertSI(group, qq)
	}
	person.KusaNum++
	kusaNum := bson.M{"$set": bson.M{"kusaNum": person.KusaNum}}
	selector := bson.M{"qq": qq}
	err = collection.Update(selector, kusaNum)
	if err != nil {
		panic(err)
	}
}

func AddImg(group, qq int64){
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SpeakInfoTable)

	person := FindByQQ(qq)
	if person == (SpeakInfo{}){
		person = InsertSI(group, qq)
	}
	person.ImgNum++
	imgNum := bson.M{"$set": bson.M{"imgNum": person.ImgNum}}
	selector := bson.M{"qq": qq}
	err = collection.Update(selector, imgNum)
	if err != nil {
		panic(err)
	}
}

func ClearDaily(group int64){
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SpeakInfoTable)

	kusa := bson.M{"$set":bson.M{"kusaNum":0}}
	img := bson.M{"$set":bson.M{"imgNum":0}}
	_, err = collection.UpdateAll(bson.M{"group": group}, kusa)
	if err != nil{
		panic(err)
	}
	_, err = collection.UpdateAll(bson.M{"group": group}, img)
	if err != nil{
		panic(err)
	}

}