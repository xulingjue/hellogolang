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

	rows, res, err := db.HgSql.Query("SELECT idpost_class,name,parent,code FROM post_class")

	var postClass []PostClass

	for _, row := range rows {
		var pc PostClass
		// You can get converted value
		pc.IdpostClass = row.Int64(0)
		pc.Name = row.Str(1)
		pc.Parent = row.Int64(2)
		pc.Code = row.Str(3)

		postClass = append(postClass, pc)
	}

	return postClass, nil
}

func (pcm *PostClassModel) Find(id int64) PostClass {
	sql := "select idpost_class,name,parent,code from post_class where idpost_class=%d"
	rows, res, err := db.HgSql.Query(sql, id)

	var postClass PostClass
	err := row.Scan(&postClass.IdpostClass, &postClass.Name, &postClass.Parent, &postClass.Code)
	if err != nil {
		fmt.Print(err)
		return postClass
	}
	return postClass
}
