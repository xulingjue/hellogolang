/*
 *数据库操作类
 */
package database

import (
	//_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	hgConfig "hellogolang/system/helper"
)

var (
	HgSql *sql.DB
	dbErr error
)

func init() {
	if HgSql == nil {
		dbName := hgConfig.GetConfig("db_name")
		dbHost := hgConfig.GetConfig("db_host")
		dbUser := hgConfig.GetConfig("db_user")
		dbPassword := hgConfig.GetConfig("db_password")

		HgSql, dbErr = sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+")/"+dbName+"?charset=utf8")
		fmt.Println(dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8")
		//检查数据库连接
		_, dbErr = HgSql.Query("SELECT 1")
		if dbErr != nil {
			panic(dbErr)
		} else {
			fmt.Println("db init success")
		}
	}
}
