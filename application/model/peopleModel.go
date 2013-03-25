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
	Idpeople   uint64
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

func (pm *PeopleModel) Find(id uint64) *People {
	sql := "select * from people where idpeople=%d"

	var people People
	row, _, err := db.HgSql.QueryFirst(sql, id)

	if err != nil {
		return nil
	}

	people.Idpeople = row.Uint64(0)
	people.Name = row.Str(1)
	people.Email = row.Str(2)
	people.Phone = row.Str(3)
	people.Avatar = row.Str(4)
	people.LastLogin = row.Str(5)
	people.CreateTime = row.Str(6)
	people.Fansnum = row.Int(7)
	people.Favnum = row.Int(8)
	people.Password = row.Str(9)
	people.QQ = row.Str(10)

	return &people
}

func (pm *PeopleModel) Insert(people People) uint64 {
	stmt, err := db.HgSql.Prepare("INSERT people SET name=?,email=?,phone=?,avatar=?,create_time=CURDATE(),lastlogin=now(),fansnum=?,favnum=?,password=?,qq=?")
	_, res, err := stmt.Exec(people.Name, people.Email, people.Phone, people.Avatar, people.Fansnum, people.Favnum, people.Password, people.QQ)

	if err != nil {
		return 0
	}
	return res.InsertId()
}

func (pm *PeopleModel) FindByName(name string) *People {
	sql := "select * from people where name=?"

	var people People
	row, _, err := db.HgSql.QueryFirst(sql, name)

	if err != nil {
		return nil
	}

	people.Idpeople = row.Uint64(0)
	people.Name = row.Str(1)
	people.Email = row.Str(2)
	people.Phone = row.Str(3)
	people.Avatar = row.Str(4)
	people.LastLogin = row.Str(5)
	people.CreateTime = row.Str(6)
	people.Fansnum = row.Int(7)
	people.Favnum = row.Int(8)
	people.Password = row.Str(9)
	people.QQ = row.Str(10)

	return &people
}

func (pm *PeopleModel) FindByEmail(email string) *People {
	sql := "select * from people where email=?"

	var people People
	row, _, err := db.HgSql.QueryFirst(sql, email)

	if err != nil {
		return nil
	}

	people.Idpeople = row.Uint64(0)
	people.Name = row.Str(1)
	people.Email = row.Str(2)
	people.Phone = row.Str(3)
	people.Avatar = row.Str(4)
	people.LastLogin = row.Str(5)
	people.CreateTime = row.Str(6)
	people.Fansnum = row.Int(7)
	people.Favnum = row.Int(8)
	people.Password = row.Str(9)
	people.QQ = row.Str(10)

	return &people
}
