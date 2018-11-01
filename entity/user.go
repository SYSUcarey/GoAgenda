package entity

import (
	"os"
	"log"
	"encoding/json"
)

//用户信息数据结构
type UserData struct{
	Name string      `json:"name"`
	Password string  `json:"password"`
	Email string     `json:"email"`
	Tel string 		 `json:"telephone"`
}

//用户数据结构
type UserList struct{
	Users []UserData `json:"users"`
}

var UserModel UserList
var fileName string = os.Getenv("GOPATH") + "/src/github.com/chenf99/GoAgenda/data/users.json"

//通过名字获得用户
func (u *UserList)GetUserByName(username string) UserData{
	for _, user := range u.Users {
		if user.Name == username {
			return user
		}
	}
	return UserData{}
}

//获取所有用户
func (u *UserList)GetAllUsers() []UserData{
	return u.Users
}

//用户是否存在
func (u *UserList)IsExist(username string) bool{
	for _, user := range u.Users {
		if user.Name == username {
			return true
		}
	}
	return false
}

//密码是否正确
func (u *UserList)MatchPass(username, password string) bool{
	for _, user := range u.Users {
		if user.Name == username {
			return user.Password == password
		}
	}
	return false
}

//添加一个用户
func (u *UserList)AddUser(userinfo UserData) {
	u.Users = append(u.Users, userinfo)
	u.saveToFile()
}

//删除一个用户
func (u *UserList)DeleteUser(username string) bool{
	for i, user := range u.Users {
		if user.Name == username {
			u.Users = append(u.Users[:i], u.Users[i+1:]...)
			u.saveToFile()
			return true
		}
	}
	return false
}

func init() {
	UserModel.readFromFile()
}

func (u *UserList)saveToFile() {
	//UserList转json格式数据
	data, err := json.Marshal(*u)
	if (err != nil) {
		log.Fatal(err)
	}
	fp, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
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

func (u *UserList)readFromFile() {
	//判断文件是否存在
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return 
	}
	fp, err := os.OpenFile(fileName, os.O_RDONLY, 0755)
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
	//解析json数据到UserList
	err = json.Unmarshal(data[:total], u)
	if err != nil {
        log.Fatal(err)
	}
}