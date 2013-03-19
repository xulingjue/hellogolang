package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type PostClass struct {
	IdpostClass int64
	Name        string
	Parent      int64
	Code        string
}

type PostClassModel struct {
	TableName string
}

func (pcm *PostClassModel) FindAll() []PostClass {

	rows, err := db.HgSql.Query("SELECT idpost_class,name,parent,code FROM post_class")

	var postClass []PostClass
	if err == nil {
		for rows.Next() {
			var pc PostClass
			err = rows.Scan(&pc.IdpostClass, &pc.Name, &pc.Parent, &pc.Code)
			if err == nil {
				postClass = append(postClass, pc)
			}
		}
	}

	return postClass
}

func (pcm *PostClassModel) Find(id int64) PostClass {
	sql := "select idpost_class,name,parent,code from post_class where idpost_class=?"
	row := db.HgSql.QueryRow(sql, id)
	var postClass PostClass
	err := row.Scan(&postClass.IdpostClass, &postClass.Name, &postClass.Parent, &postClass.Code)
	if err != nil {
		fmt.Print(err)
		return postClass
	}
	return postClass
}
