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

func (pcm *PostClassModel) FindAll() ([]PostClass, error) {

	var postClass []PostClass
	rows, _, err := db.HgSql.Query("SELECT idpost_class,name,parent,code FROM post_class")

	if err != nil {
		fmt.Println(err)
		return postClass, nil
	}

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

func (pcm *PostClassModel) Find(id int64) (PostClass, error) {
	sql := "select idpost_class,name,parent,code from post_class where idpost_class=%d"

	var postClass PostClass
	_, res, err := db.HgSql.Query(sql, id)
	row, err := res.GetRow()

	if err != nil {
		fmt.Println(err)
		return postClass, err
	}

	postClass.IdpostClass = row.Int64(0)
	postClass.Name = row.Str(1)
	postClass.Parent = row.Int64(2)
	postClass.Code = row.Str(3)

	if err != nil {
		fmt.Print(err)
		return postClass, err
	}

	return postClass, nil
}
