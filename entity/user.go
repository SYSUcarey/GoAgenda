package entity

//用户信息数据结构
type UserData struct{
	Name string      `json:"name"`
	Password string  `json:"password"`
	Email string     `json:"email"`
	Tel string 		 `json:"telephone"`
}
