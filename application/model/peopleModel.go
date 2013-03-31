package model

/*
	Find 开头的  返回单个值
	FindAll 开头的 返回所有值
*/

import (
	"fmt"
	db "hellogolang/system/database"
)

type People struct {
	Idpeople   int64
	Name       string
	Email      string
	Phone      string
	Avatar     string
	LastLogin  string
	CreateTime string
	Fansnum    int
	Favnum     int
	Postnum    int
	Questionum int
	Password   string
	QQ         string
}

type PeopleModel struct {
	TableName string
}

func (pm *PeopleModel) Find(id int64) *People {
	sql := "select people.idpeople,people.name,people.email,people.phone,people.avatar,people.lastlogin," +
		"people.create_time,people.fansnum,people.favnum,people.password,people.qq,people.postnum,people.quesnum from people where idpeople=?"

	stmt, err := db.HgSql.Prepare(sql)
	row := stmt.QueryRow(id)

	if err != nil || row == nil {
		fmt.Println(err)
		return nil
	}

	var people People
	row.Scan(&people.Idpeople, &people.Name, &people.Email, &people.Phone, &people.Avatar, &people.LastLogin,
		&people.CreateTime, &people.Fansnum, &people.Favnum, &people.Password, &people.Phone, &people.Postnum, &people.Questionum)

	return &people
}

func (pm *PeopleModel) Insert(people People) *People {
	stmt, err := db.HgSql.Prepare("INSERT people SET name=?,email=?,phone=?,avatar=?,create_time=CURDATE(),lastlogin=now(),fansnum=?,favnum=?,password=?,qq=?,postnum=?,quesnum=?")
	res, err := stmt.Exec(people.Name, people.Email, people.Phone, people.Avatar, people.Fansnum, people.Favnum, people.Password, people.QQ, people.Postnum, people.Questionum)

	insertId, err := res.LastInsertId()
	if err != nil {
		return nil
	}

	people.Idpeople = insertId
	return &people
}

func (pm *PeopleModel) FindByName(name string) *People {
	sql := "select people.idpeople,people.name,people.email,people.phone,people.avatar,people.lastlogin," +
		"people.create_time,people.fansnum,people.favnum,people.password,people.qq,people.postnum,people.quesnum from people where name=?"

	stmt, err := db.HgSql.Prepare(sql)
	row := stmt.QueryRow(name)

	var people People
	err = row.Scan(&people.Idpeople, &people.Name, &people.Email, &people.Phone, &people.Avatar, &people.LastLogin,
		&people.CreateTime, &people.Fansnum, &people.Favnum, &people.Password, &people.Phone, &people.Postnum, &people.Questionum)

	if err != nil || row == nil {
		fmt.Println(err)
		return nil
	}

	return &people
}

func (pm *PeopleModel) FindByEmail(email string) *People {
	sql := "select people.idpeople,people.name,people.email,people.phone,people.avatar,people.lastlogin," +
		"people.create_time,people.fansnum,people.favnum,people.password,people.qq,people.postnum,people.quesnum  from people where email=?"

	stmt, err := db.HgSql.Prepare(sql)
	row := stmt.QueryRow(email)

	var people People
	err = row.Scan(&people.Idpeople, &people.Name, &people.Email, &people.Phone, &people.Avatar, &people.LastLogin,
		&people.CreateTime, &people.Fansnum, &people.Favnum, &people.Password, &people.Phone, &people.Postnum, &people.Questionum)

	if err != nil || row == nil {
		fmt.Println(err)
		return nil
	}

	return &people
}

func (pm *PeopleModel) Update(people People) {
	stmt, err := db.HgSql.Prepare("update people SET name=?,email=?,phone=?,avatar=?,create_time=CURDATE(),lastlogin=now(),fansnum=?,favnum=?,password=?,qq=?,postnum=?,quesnum=? where idpeople=?")
	_, err = stmt.Exec(people.Name, people.Email, people.Phone, people.Avatar, people.Fansnum, people.Favnum, people.Password, people.QQ, people.Postnum, people.Questionum, people.Idpeople)
	if err != nil {
		fmt.Println(err)
	}
}
