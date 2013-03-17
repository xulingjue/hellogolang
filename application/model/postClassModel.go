package model

import (
	//"fmt"
	db "hellogolang/system/database"
)

type PostClass struct {
	IdpostClass int64
	Name        string
	Parent      int64
}

type PostClassModel struct {
	TableName string
}

func (pcm *PostClassModel) FindAll(id int) []PostClass {

	rows, err := db.HgSql.Query("SELECT idpost_class,name FROM post_class", id)

	var postClass []PostClass
	if err == nil {
		for rows.Next() {
			var pc PostClass
			err = rows.Scan(&pc.IdpostClass, &pc.Name, &pc.Parent)
			if err == nil {
				postClass = append(postClass, pc)
			}
		}
	}

	return postClass
}
