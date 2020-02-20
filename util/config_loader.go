package util
//
//import (
//	"encoding/json"
//	"io/ioutil"
//	"os"
//)
//const config = "./config.json"
//
//type MongoConfig struct {
//	URL string		`json:"url"`
//	DBname string 	`json:"DBname"`
//	SpeakInfoTable string 	`json:"speak_info_table"`
//	SelfDeterminedCMD string 	`json:"self_determined_cmd"`
//}
//type Path struct {
//	TextPath string		`json:"text_path"`
//}
//
//type Account struct {
//	THGroup int64	`json:"th_group"`
//	MyQQ	int64	`json:"my_qq"`
//}
//
//func GetAccount() Account{
//	f, err := os.Open("./config.json")
//	if err != nil{
//		panic(err)
//	}
//	defer f.Close()
//	data, _ := ioutil.ReadAll(f)
//
//	var account Account
//	json.Unmarshal(data, account)
//	return account
//}
//
//func GetMongo() MongoConfig{
//	f, err := os.Open("./config.json")
//	if err != nil{
//		panic(err)
//	}
//	defer f.Close()
//	data, _ := ioutil.ReadAll(f)
//
//	var mongo MongoConfig
//	json.Unmarshal(data, account)
//	return mongo
//}
//
//func GetPath() Path{
//	f, err := os.Open("./config.json")
//	if err != nil{
//		panic(err)
//	}
//	defer f.Close()
//	data, _ := ioutil.ReadAll(f)
//
//	var path Path
//	json.Unmarshal(data, account)
//	return path
//}