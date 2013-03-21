package model

import (
	//"fmt"
	db "hellogolang/system/database"
)

type Comment struct {
	Idcomment  int64
	Idpost     int64
	CreateTime string
	Content    string
	Idpeople   int64
	Parent     int64
	Author     People
}

type CommentModel struct {
	TableName string
}

func (cm *CommentModel) FindAllCountByPostID(postId int64) (int, error) {
	sql := "select count(*) as total " +
		"from comment where comment.idpost=%d"

	rows, res, err := db.HgSql.Query(sql, postId)

	if err != nil {
		return 0, err
	}

	row, err := res.GetRow()
	num := row.Int(0)

	return num, nil
}

func (cm *CommentModel) FindAllByPostID(postId int64, page int, pageSize int) ([]Comment, error) {
	sql := "select comment.idcomment,comment.idpost,comment.create_time,comment.content,comment.Idpeople,comment.parent," +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople where comment.idpost=%d order by comment.create_time desc limit %d,%d"

	rows, res, err := db.HgSql.Query(sql, postId, (page-1)*pageSize, pageSize)

	var comments []Comment
	for _, row := range rows {

		var comment Comment
		// You can get converted value
		comment.Idcomment = row.Int64(0) // Zero value
		comment.Idpost = row.Int64(1)
		comment.CreateTime = row.Str(2)
		comment.Content = row.Str(3)
		comment.Idpeople = row.Int64(4)
		comment.Parent = row.Int64(5)

		comment.Author.Idpeople = row.Int64(6)
		comment.Author.Name = row.Str(7)
		comment.Author.Avatar = row.Str(8)

		comments = append(comments, comment)
	}

	return comments, nil
}

func (cm *CommentModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Comment, error) {
	sql := "select comment.idcomment,comment.idpost,comment.create_time,comment.content,comment.Idpeople,comment.parent," +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople "

	orderby := "order by comment.create_time desc limit %d,%d"

	if len(agrs) > 0 {
		sql = sql + " where "
	} else {
		sql = sql + " and "
	}

	for k, v := range agrs {
		sql = sql + " " + k + v + " and"
	}

	sql = sql + " 1=1 " + orderby

	rows, res, err := db.HgSql.Query(sql, (page-1)*pageSize, pageSize)

	var comments []Comment

	for _, row := range rows {

		var comment Comment
		// You can get converted value
		comment.Idcomment = row.Int64(0) // Zero value
		comment.Idpost = row.Int64(1)
		comment.CreateTime = row.Str(2)
		comment.Content = row.Str(3)
		comment.Idpeople = row.Int64(4)
		comment.Parent = row.Int64(5)

		comment.Author.Idpeople = row.Int64(6)
		comment.Author.Name = row.Str(7)
		comment.Author.Avatar = row.Str(8)

		comments = append(comments, comment)
	}

	return comments, nil
}

func (cm *CommentModel) FindAllCount(agrs map[string]string) (int, error) {
	sql := "select count(*) as total " +
		"from comment "

	if len(agrs) > 0 {
		sql = sql + " where "
	} else {
		sql = sql + " and "
	}

	for k, v := range agrs {
		sql = sql + " " + k + v + " and"
	}

	sql = sql + " 1=1 "

	rows, res, err := db.HgSql.Query(sql)

	if err != nil {
		return 0, err
	}

	row, err := res.GetRow()
	num := row.Int(0)

	return num, nil
}

func (cm *CommentModel) Insert(comment Comment) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT comment SET idpost=?,content=?,idpeople=?,parent=?,create_time=now()")
	rows, res, err := stmt.Exec(comment.Idpost, comment.Content, comment.Idpeople, comment.Parent)
	/*	id, err := rows.LastInsertId()
		if err != nil {
			return 0, err
		}*/
	return 0, nil
}
