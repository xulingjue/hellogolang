package datebase

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	hgConfig "hellogolang/core/config"
)

var (
	HgDB = *sql.DB
)

func init() {
	//初始化数据库
	dbName := hgConfig.GetValue("db_name")
	dbHost := hgConfig.GetValue("db_host")
	dbUser := hgConfig.GetValue("db_user")
	dbPassword := hgConfig.GetValue("db_password")

	HgDB, dbErr := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+")/"+dbName+"?charset=utf8")
	if dbErr != nil {
		//logMessage("db connect error")
		os.Exit(1)
	}
	//检查数据库连接
	_, dbErr = HgDB.Query("SELECT 1")
	if dbErr != nil {
		//panic(dbErr)
		//logMessage("db select error")
		os.Exit(1)
	} else {
		//logMessage("db select success")
	}
}
