package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	if db == nil {
		ndb, _ := sql.Open("mysql", "root:xin6816839@tcp(127.0.0.1:3306)/blog")
		db = ndb
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

//数据库操作
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, e := db.Exec(sql, args...)
	if e != nil {
		log.Println(e)
		return 0, e
	}
	con, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return con, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

	ModifyDB(sql)

}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}
