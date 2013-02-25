/*
 *数据库操作类
 */
package database

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
	hgConfig "hellogolang/system/helper"
)

type Hgdb struct {
	db *sql.DB
}

/* 初始化数据库引擎 */
func Init() (*Hgdb, error) {
	hgDb := new(Hgdb)
	dbName := hgConfig.GetConfig("db_name")
	dbHost := hgConfig.GetConfig("db_host")
	dbUser := hgConfig.GetConfig("db_user")
	dbPassword := hgConfig.GetConfig("db_password")

	dbConnect, dbErr := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+")/"+dbName+"?charset=utf8")
	if dbErr != nil {
		//logMessage("db connect error")
		return nil, dbErr
	}
	//检查数据库连接
	_, dbErr = dbConnect.Query("SELECT 1")
	if dbErr != nil {
		fmt.Println("db error")
		return nil, dbErr
	}

	hgDb.db = dbConnect
	return hgDb, nil
}

/*
 *根据id查找
 */
func (h *Hgdb) Find(id int) {

}

/*
 *
 */
func (h *Hgdb) Find_all_by() {

}

/*
 *执行sql语句
 */
func (h *Hgdb) Query(sql string) {

}
