/*
	一些公共函数
*/
package models

import (
	"fmt"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func init() {

}
