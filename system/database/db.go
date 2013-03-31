/*
 *数据库操作类
 */
package database

import (
	"fmt"
	//"hellogolang/system/helper"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var (
	HgSql *sql.DB
)

func init() {
	var err error
	if HgSql == nil {
		HgSql, err = sql.Open("mysql", "root:hg6688@tcp(localhost:3306)/hellogolang?charset=utf8")
		if err != nil {
			fmt.Println("db connect error")
		} else {
			fmt.Println("db init success")
		}
	}
}

//github.com/Go-SQL-Driver/MySQL
