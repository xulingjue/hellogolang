package model

/*
	Find 开头的  返回单个值
	FindAll 开头的 返回所有值
*/

import (
	//"fmt"
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
	Password   string
	QQ         string
}

type PeopleModel struct {
	TableName string
}

func (pm *PeopleModel) Find(id int64) (People, error) {
	row := db.HgSql.QueryRow("select * from people where idpeople=?", id)
	var people People
	err := row.Scan(&people.Idpeople, &people.Name, &people.Email, &people.Phone, &people.Avatar, &people.LastLogin, &people.CreateTime, &people.Fansnum, &people.Favnum, &people.Password, &people.QQ)
	if err != nil {
		return people, err
	}
	return people, nil
}

func (pm *PeopleModel) Insert(people People) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT people SET name=?,email=?,phone=?,avatar=?,create_time=CURDATE(),lastlogin=now(),fansnum=?,favnum=?,password=?,qq=?")
	res, err := stmt.Exec(people.Name, people.Email, people.Phone, people.Avatar, people.Fansnum, people.Favnum, people.Password, people.QQ)
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pm *PeopleModel) FindByName(name string) (People, error) {
	row := db.HgSql.QueryRow("select * from people where name=?", name)
	var people People
	err := row.Scan(&people.Idpeople, &people.Name, &people.Email, &people.Phone, &people.Avatar, &people.LastLogin, &people.CreateTime, &people.Fansnum, &people.Favnum, &people.Password, &people.QQ)
	if err != nil {
		return people, err
	}
	return people, nil
}

func (pm *PeopleModel) FindByEmail(email string) (People, error) {
	row := db.HgSql.QueryRow("select * from people where email=?", email)
	var people People
	err := row.Scan(&people.Idpeople, &people.Name, &people.Email, &people.Phone, &people.Avatar, &people.LastLogin, &people.CreateTime, &people.Fansnum, &people.Favnum, &people.Password, &people.QQ)
	if err != nil {
		return people, err
	}
	return people, nil
}
