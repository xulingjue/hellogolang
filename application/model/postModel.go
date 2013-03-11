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
	ReadNum     int
	ReplyNum    int
	Title       string

	Author People
	Class  PostClass
	Type   PostType
}

type PostModel struct {
	TableName string
}

func (pm *PostModel) Find(id int64) (Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.idpost_type,post_class.name," +
		"post_type.idpost_type,post_type.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople left join post_type on post_type.idpost_type=post_class.idpost_type where idpost=?"

	row := db.HgSql.QueryRow(sql, id)
	var post Post
	err := row.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum, &post.Title,
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

func (pm *PostModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.idpost_type,post_class.name," +
		"post_type.idpost_type,post_type.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople left join post_type on post_type.idpost_type=post_class.idpost_type "

	orderby := "order by post.create_time desc limit ?,?"

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

	var posts []Post
	if err == nil {
		for rows.Next() {
			var post Post
			err := rows.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum,
				&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
				&post.Class.IdPostClass, &post.Class.IdPostType, &post.Class.Name,
				&post.Type.IdPostType, &post.Type.Name,
			)
			if err == nil {
				posts = append(posts, post)
			}
		}
	} else {
		fmt.Println(err)
	}
	return posts, nil
}

func (pm *PostModel) Insert(post Post) {

}

func (pm *PostModel) FindAllReply(postId int64, page int, pageSize int) ([]Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.idpost_type,post_class.name," +
		"post_type.idpost_type,post_type.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople left join post_type on post_type.idpost_type=post_class.idpost_type where post.parentid=? order by create_time desc limit ?,? "

	rows, err := db.HgSql.Query(sql, postId, (page-1)*pageSize, pageSize)
	var posts []Post
	if err == nil {
		for rows.Next() {
			var post Post
			err := rows.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum,
				&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
				&post.Class.IdPostClass, &post.Class.IdPostType, &post.Class.Name,
				&post.Type.IdPostType, &post.Type.Name,
			)
			if err == nil {
				posts = append(posts, post)
			}
		}
	} else {
		fmt.Println(err)
	}
	return posts, nil
}
