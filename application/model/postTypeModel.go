package model

import (
	//"fmt"
	db "hellogolang/system/database"
)

type PostType struct {
	IdPostType int64
	Name       string
}

type PostTypeModel struct {
	TableName string
}

func (ptm *PostTypeModel) FindAll() []PostType {
	//查询数据
	rows, err := db.HgSql.Query("SELECT * FROM post_type")

	var postTypes []PostType
	if err == nil {
		for rows.Next() {
			var pt PostType
			err = rows.Scan(&pt.IdPostType, &pt.Name)
			if err == nil {
				postTypes = append(postTypes, pt)
			}
		}
	}
	return postTypes
}
