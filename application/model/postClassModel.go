package model

import (
	//"fmt"
	db "hellogolang/system/database"
)

type PostClass struct {
	IdPostClass int64
	Name        string
	IdPostType  int64
}

type PostClassModel struct {
	TableName string
}

func (pcm *PostClassModel) FindAll(id int) []PostClass {

	rows, err := db.HgSql.Query("SELECT idpost_class,name,idpost_type FROM post_class where idpost_type = ?", id)

	var postClass []PostClass
	if err == nil {
		for rows.Next() {
			var pc PostClass
			err = rows.Scan(&pc.IdPostClass, &pc.Name, &pc.IdPostType)
			if err == nil {
				postClass = append(postClass, pc)
			}
		}
	}

	return postClass
}
