package util

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Reflection struct {
	Cmd string				`bson:"cmd"`
	Group int64				`bson:"group"`
	TeacherQQs []int64		`bson:"teacherQQs"`
	TeacherNum int			`bson:"teacher_num"`
	ReMsg string			`bson:"remsg"`
}

const (
	SelfDeterminedCmd = "reflections"
)
//命令格式
//def "命令" return "回复"


//插入词条
func InsertCmd(ref Reflection) Reflection{
	session, err := mgo.Dial(URL)
	if err != nil{
		panic(err)
	}
	defer session.Close()
	c := session.DB(DBname).C(SelfDeterminedCmd)		//切换到调教表


	err = c.Insert(ref)
	if err != nil{
		log.Println("insert:")
		return Reflection{}
	}
	return ref
}

//查看当前已有反射
func FindRef(cmd string) Reflection{
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SelfDeterminedCmd)

	ref := Reflection{}
	err = collection.Find(bson.M{"cmd":cmd}).One(&ref)
	if err != nil{
		log.Println("read:")
	}
	return ref
}

//检查新旧ref是否相同
func IsRefTheSame(oldRef,newRef Reflection) bool {
	if newRef.Group == oldRef.Group &&
			newRef.ReMsg == oldRef.ReMsg{
		return true
	}else{
		return false
	}
}


func UpdateRef(newRef Reflection)error{
	session, err := mgo.Dial(URL)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SelfDeterminedCmd)

	ref := FindRef(newRef.Cmd)
	ref.TeacherNum++
	ref.TeacherQQs = append(ref.TeacherQQs, newRef.TeacherQQs[0])

	selector := bson.M{"cmd": newRef.Cmd,"group":newRef.Group}
	newdata := bson.M{"$set": bson.M{"teacher_num": ref.TeacherNum, "teacherQQs":ref.TeacherQQs, "remsg":ref.ReMsg}}

	err = collection.Update(selector, newdata)
	if err != nil {
		return err
	}
	return nil
}

//覆盖原来词条
func CoverRef(newRef Reflection)error {
	session, err := mgo.Dial(URL)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(DBname).C(SelfDeterminedCmd)

	selector := bson.M{"cmd": newRef.Cmd,"group":newRef.Group}
	newdata := bson.M{"$set": bson.M{"teacher_num": newRef.TeacherNum, "teacherQQs":newRef.TeacherQQs, "remsg":newRef.ReMsg}}
	err = collection.Update(selector, newdata)
	if err != nil {
		return err
	}
	return nil
}