/*
 *数据库操作类
 */
package database

import (
	"fmt"
	hgConfig "hellogolang/HooGL/config"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var (
	HgSql *sql.DB
)

func init() {
	var err error
	if HgSql == nil {

		dbUser := hgConfig.GetConfig("db_user")
		dbPassword := hgConfig.GetConfig("db_password")
		dbHost := hgConfig.GetConfig("db_host")
		dbName := hgConfig.GetConfig("db_name")


		HgSql, err = sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+")/"+dbName+"?charset=utf8")
		if err != nil {
			fmt.Println("db connect error")
		} else {
			fmt.Println("db init success")
		}
	}
}

//github.com/Go-SQL-Driver/MySQL
