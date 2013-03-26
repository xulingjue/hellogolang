package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type PostClass struct {
	IdpostClass int64
	Parent      int64
	Name        string
	Code        string
	Children    []PostClass //暂时未使用
}

type PostClassModel struct {
	TableName string
}

func (pcm *PostClassModel) FindAll() []PostClass {
	var postClasses []PostClass
	rows, err := db.HgSql.Query("SELECT idpost_class,name,parent,code FROM post_class")

	if err != nil {
		fmt.Println(err)
		return postClasses
	}

	for rows.Next() {
		var pc PostClass
		rows.Scan(&pc.IdpostClass, &pc.Name, &pc.Parent, &pc.Code)
		postClasses = append(postClasses, pc)
	}

	return postClasses
}

func (pcm *PostClassModel) Find(id int64) *PostClass {
	sql := "select idpost_class,name,parent,code from post_class where idpost_class=?"

	stmt, err := db.HgSql.Prepare(sql)
	row := stmt.QueryRow(id)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var postClass PostClass
	row.Scan(&postClass.IdpostClass, &postClass.Name, &postClass.Parent, &postClass.Code)

	return &postClass
}
