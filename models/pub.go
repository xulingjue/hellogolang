/*
	一些公共函数
*/
package models

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
	hgHeplers "hellogolang/hghelpers"
)

var (
	db *sql.DB
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func init() {
	db, dbErr := sql.Open("mysql", "root:adzure1105@/tcp(192.168.1.151:3306)/hellogolang?charset=utf8")
	if dbErr != nil {
		//hgHeplers.logMessage("db error")
	}
}
