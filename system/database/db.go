/*
 *数据库操作类
 */
package database

import (
	"fmt"
	"hellogolang/system/helper"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

var (
	HgSql mysql.Conn
)

func init() {
	if HgSql == nil {
		dbName := helper.GetConfig("db_name")
		dbHost := helper.GetConfig("db_host")
		dbUser := helper.GetConfig("db_user")
		dbPassword := helper.GetConfig("db_password")

		HgSql = mysql.New("tcp", "", dbHost, dbUser, dbPassword, dbName)
		err := HgSql.Connect()
		if err != nil {
			panic(err)
			fmt.Println("db connect error")
		} else {
			fmt.Println("db init success")
		}
	}

}
