package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type Post struct {
	Idpost      int64
	Idpeople    int64
	Content     string
	Parentid    int64
	IdpostClass int64
	CreateTime  string
	ReprintFrom string
	ReprintUrl  string

	Author People
	Class  PostClass
	Type   PostType
}

type PostModel struct {
	TableName string
}

func (pm *PostModel) Find(id int64) (Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.idpost_type,post_class.name," +
		"post_type.idpost_type,post_type.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople left join post_type on post_type.idpost_type=post_class.idpost_type where idpost=?"

	row := db.HgSql.QueryRow(sql, id)
	var post Post
	err := row.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl,
		&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
		&post.Class.IdPostClass, &post.Class.IdPostType, &post.Class.Name,
		&post.Type.IdPostType, &post.Type.Name,
	)
	if err != nil {
		fmt.Print(err)
		return post, err
	}
	return post, nil
}

func (pm *PostModel) FindAll(page int, pageSize int) {

}

func (pm *PostModel) Insert(post Post) {

}

func (pm *PostModel) FindAllReply(page int, pageSize int) {

}
