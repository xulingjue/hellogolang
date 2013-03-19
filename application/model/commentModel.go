package model

import (
	"fmt"
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
		"from comment where comment.idpost=?"
	var num int

	rows := db.HgSql.QueryRow(sql, postId)
	err := rows.Scan(&num)

	if err != nil {
		return num, err
	}

	return num, nil
}

func (cm *CommentModel) FindAllByPostID(postId int64, page int, pageSize int) ([]Comment, error) {
	sql := "select comment.idcomment,comment.idpost,comment.create_time,comment.content,comment.Idpeople,comment.parent," +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople where comment.idpost=? order by comment.create_time desc limit ?,?"

	rows, err := db.HgSql.Query(sql, postId, (page-1)*pageSize, pageSize)

	var comments []Comment
	if err == nil {
		for rows.Next() {
			var comment Comment
			err := rows.Scan(&comment.Idcomment, &comment.Idpost, &comment.CreateTime, &comment.Content, &comment.Idpeople, &comment.Parent,
				&comment.Author.Idpeople, &comment.Author.Name, &comment.Author.Avatar,
			)
			if err == nil {
				comments = append(comments, comment)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
	}
	return comments, nil
}

func (cm *CommentModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Comment, error) {
	sql := "select comment.idcomment,comment.idpost,comment.create_time,comment.content,comment.Idpeople,comment.parent," +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople "

	orderby := "order by comment.create_time desc limit ?,?"

	if len(agrs) > 0 {
		sql = sql + " where "
	} else {
		sql = sql + " and "
	}

	for k, v := range agrs {
		sql = sql + " " + k + v + " and"
	}

	sql = sql + " 1=1 " + orderby

	rows, err := db.HgSql.Query(sql, (page-1)*pageSize, pageSize)

	var comments []Comment
	if err == nil {
		for rows.Next() {
			var comment Comment
			err := rows.Scan(&comment.Idcomment, &comment.Idpost, &comment.CreateTime, &comment.Content, &comment.Idpeople, &comment.Parent,
				&comment.Author.Idpeople, &comment.Author.Name, &comment.Author.Avatar,
			)
			if err == nil {
				comments = append(comments, comment)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
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

	var num int

	rows := db.HgSql.QueryRow(sql)
	err := rows.Scan(&num)

	if err != nil {
		return num, err
	}

	return num, nil
}

func (cm *CommentModel) Insert(comment Comment) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT comment SET idpost=?,content=?,idpeople=?,parent=?,create_time=now()")
	res, err := stmt.Exec(comment.Idpost, comment.Content, comment.Idpeople, comment.Parent)
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
