package models

import (
	"blog/utils"
	"fmt"
)

type User struct {
	Id         int
	UserName   string
	Password   string
	Status     int
	CreateTime int64
}

//插入数据
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values(?,?,?,?)", user.UserName, user.Password, user.Status, user.CreateTime)
}

//查询

func QueryUser(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUser(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUser(sql)
}
