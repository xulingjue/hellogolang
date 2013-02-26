/*
 *数据库操作类
 */
package database

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	hgConfig "hellogolang/system/helper"
)

var (
	db *sql.DB
)

func init() {
	if db == nil {
		dbName := hgConfig.GetConfig("db_name")
		dbHost := hgConfig.GetConfig("db_host")
		dbUser := hgConfig.GetConfig("db_user")
		dbPassword := hgConfig.GetConfig("db_password")

		db, dbErr := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+")/"+dbName+"?charset=utf8")
		if dbErr != nil {

		}
		//检查数据库连接
		_, dbErr = db.Query("SELECT 1")
		if dbErr != nil {

		}
	}
}

func Query(sql string) {

}
