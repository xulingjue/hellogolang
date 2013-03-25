package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type Comment struct {
	Idcomment uint64
	Idpost    uint64
	Parent    uint64

	CreateTime string
	Content    string
	Author     People
}

type CommentModel struct {
	TableName string
}

func (cm *CommentModel) FindAllByPostID(postId uint64, page int, pageSize int) ([]Comment, int) {
	countSql := "select count(*) as total from comment where comment.idpost=%d"
	rowCount, _, err := db.HgSql.QueryFirst(countSql, postId)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	count := rowCount.Int(0)

	cmSql := "select comment.idcomment,comment.idpost,comment.parent,comment.create_time,comment.content" +
		" people.idpeople,people.name,people.avatar from comment " +
		" left join people on comment.Idpeople=people.idpeople where comment.idpost='%d' order by comment.create_time desc limit %d,%d"
	rowsCm, _, err := db.HgSql.Query(cmSql, postId, (page-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	var comments []Comment
	for _, row := range rowsCm {
		var comment Comment
		comment.Idcomment = row.Uint64(0)
		comment.Idpost = row.Uint64(1)
		comment.Parent = row.Uint64(2)
		comment.CreateTime = row.Str(3)
		comment.Content = row.Str(4)

		comment.Author.Idpeople = row.Uint64(5)
		comment.Author.Name = row.Str(6)
		comment.Author.Avatar = row.Str(7)
		comments = append(comments, comment)
	}

	return comments, count
}

func (cm *CommentModel) Insert(comment Comment) uint64 {
	stmt, err := db.HgSql.Prepare("INSERT comment SET idpost=?,content=?,idpeople=?,parent=?,create_time=now()")
	_, res, err := stmt.Exec(comment.Idpost, comment.Content, comment.Author.Idpeople, comment.Parent)
	if err != nil {
		return 0
	}
	return res.InsertId()
}
