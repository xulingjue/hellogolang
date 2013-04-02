package model

import (
	"fmt"
	db "hellogolang/HooGL/database"
)

type Comment struct {
	Idcomment int64
	Idpost    int64
	Parent    int64

	CreateTime string
	Content    string
	Author     People
}

type CommentModel struct {
	TableName string
}

func (cm *CommentModel) FindAllByPostID(postId int64, page int, pageSize int) ([]Comment, int) {
	countSql := "select count(*) as total from comment where comment.idpost=?"
	stmt, err := db.HgSql.Prepare(countSql)
	row := stmt.QueryRow(postId)

	count := 0
	row.Scan(&count)

	cmSql := "select comment.idcomment,comment.idpost,comment.parent,comment.create_time,comment.content," +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople where comment.idpost=? order by comment.create_time desc limit ?,?"

	stmt, err = db.HgSql.Prepare(cmSql)

	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	rows, err := stmt.Query(postId, (page-1)*pageSize, pageSize)

	if err != nil {
		return nil, 0
	}

	var comments []Comment

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.Idcomment, &comment.Idpost, &comment.Parent, &comment.CreateTime, &comment.Content,
			&comment.Author.Idpeople, &comment.Author.Name, &comment.Author.Avatar)
		if err == nil {
			comments = append(comments, comment)
		}
	}

	return comments, count
}

func (cm *CommentModel) Insert(comment Comment) *Comment {
	stmt, err := db.HgSql.Prepare("INSERT comment SET idpost=?,content=?,idpeople=?,parent=?,create_time=now()")
	fmt.Println("post id")
	fmt.Println(comment.Idpost)
	res, err := stmt.Exec(comment.Idpost, comment.Content, comment.Author.Idpeople, comment.Parent)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	comment.Idcomment = insertId
	return &comment
}
