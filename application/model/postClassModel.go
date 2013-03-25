package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type PostClass struct {
	IdpostClass uint64
	Parent      uint64
	Name        string
	Code        string
	Children    []PostClass //暂时未使用
}

type PostClassModel struct {
	TableName string
}

func (pcm *PostClassModel) FindAll() []PostClass {
	var postClasses []PostClass
	rows, _, err := db.HgSql.Query("SELECT idpost_class,name,parent,code FROM post_class")

	if err != nil {
		fmt.Println(err)
		return postClasses
	}

	for _, row := range rows {
		var pc PostClass
		// You can get converted value
		pc.IdpostClass = row.Uint64(0)
		pc.Name = row.Str(1)
		pc.Parent = row.Uint64(2)
		pc.Code = row.Str(3)

		postClasses = append(postClasses, pc)
	}

	return postClasses
}

func (pcm *PostClassModel) Find(id uint64) *PostClass {
	sql := "select idpost_class,name,parent,code from post_class where idpost_class=%d"

	_, res, err := db.HgSql.Query(sql, id)
	row, err := res.GetRow()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var postClass PostClass
	postClass.IdpostClass = row.Uint64(0)
	postClass.Name = row.Str(1)
	postClass.Parent = row.Uint64(2)
	postClass.Code = row.Str(3)

	return &postClass
}
