package main

import (
	"github.com/jinzhu/gorm"
	//这一行需要保留，否则会报import _ "github.com/go-sql-driver/mysql"错误
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	conn *gorm.DB
}

func MysqlDefault() (db *Mysql) {
	mysql := &Mysql{}
	conn, err := gorm.Open("mysql", "root:123456@tcp(172.16.57.110:3306)/db_user")
	if err != nil {
		panic(err.Error() + "mysql")
	}
	conn = conn.Debug()
	mysql.conn = conn
	return mysql
}

func (this *Mysql) GetConn() *gorm.DB {
	return this.conn
}
