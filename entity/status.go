package entity

import (
	"os"
	"log"
	"encoding/json"
)

type status struct{
	UserName string	`json:"username"`
	Password string `json:"password"`
	Islogin bool 	`json:"islogin"`
}

var CurStatus status
var statusfile string = os.Getenv("GOPATH") + "/src/github.com/chenf99/GoAgenda/data/status.json"

func init() {
	CurStatus.readFromFile()
}

func (c *status)GetStatus() status{
	return CurStatus
}

func (c *status)LogIn(username, password string) {
	c = &status{
		UserName : username,
		Password : password,
		Islogin  : true,
	}
	c.saveToFile()
}

func (c *status)LogOut() {
	c = &status{
		UserName : "",
		Password : "",
		Islogin  : false,
	}
	c.saveToFile()
}

func (c *status)saveToFile() {
	//status转json格式数据
	data, err := json.Marshal(*c)
	if (err != nil) {
		log.Fatal(err)
	}
	fp, err := os.OpenFile(statusfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	//写入文件
	_, err = fp.Write(data)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *status)readFromFile() {
	//判断文件是否存在
	_, err := os.Stat(statusfile)
	if os.IsNotExist(err) {
		return 
	}
	fp, err := os.OpenFile(statusfile, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1000)
	//读取文件
	total, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	//解析json数据到status
	err = json.Unmarshal(data[:total], c)
	if err != nil {
        log.Fatal(err)
	}
}