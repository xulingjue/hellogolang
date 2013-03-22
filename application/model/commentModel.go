package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type Comment struct {
	Idcomment  uint64
	Idpost     uint64
	CreateTime string
	Content    string
	Idpeople   uint64
	Parent     uint64
	Author     People
}

type CommentModel struct {
	TableName string
}

func (cm *CommentModel) FindAllCountByPostID(postId uint64) (int, error) {
	sql := "select count(*) as total " +
		"from comment where comment.idpost=%d"

	rows, _, err := db.HgSql.Query(sql, postId)

	if err != nil {
		return 0, err
	}

	for _, row := range rows {
		fmt.Println(row.Int(0))
		return row.Int(0), nil
	}

	return 0, nil
}

func (cm *CommentModel) FindAllByPostID(postId int64, page int, pageSize int) ([]Comment, error) {
	sql := "select comment.idcomment,comment.idpost,comment.create_time,comment.content,comment.Idpeople,comment.parent," +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople where comment.idpost=%d order by comment.create_time desc limit %d,%d"

	rows, _, err := db.HgSql.Query(sql, postId, (page-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err)
	}

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

	var comments []Comment
	rows, _, err := db.HgSql.Query(sql, (page-1)*pageSize, pageSize)

	if err != nil {
		return comments, err
	}

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

	_, res, err := db.HgSql.Query(sql)

	if err != nil {
		return 0, err
	}

	row, err := res.GetRow()
	num := row.Int(0)

	return num, nil
}

func (cm *CommentModel) Insert(comment Comment) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT comment SET idpost=?,content=?,idpeople=?,parent=?,create_time=now()")
	_, _, err = stmt.Exec(comment.Idpost, comment.Content, comment.Idpeople, comment.Parent)
	if err != nil {
		fmt.Println(err)
	}
	/*	id, err := rows.LastInsertId()
		if err != nil {
			return 0, err
		}*/
	return 0, nil
}
