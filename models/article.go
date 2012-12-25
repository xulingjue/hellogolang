package models

import (
	"fmt"
	"time"
)

//用户
type People struct {
	Id         string
	Name       string
	Email      string
	Phone      string
	Avatar     string
	Lastlogin  time.Time
	CreateTime time.Time
	Fansunm    int
	Favnum     int
}

func findPeopleByName(name string) *People {
	stmt, err := db.Prepare("select * from people where name=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		fmt.Println("hello world!")
	}

}
