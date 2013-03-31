package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type Post struct {
	Idpost      int64
	Content     string
	CreateTime  string
	ReprintFrom string
	ReprintUrl  string
	ReadNum     int //阅读数
	ReplyNum    int //回复数
	Title       string

	Author People
	Class  PostClass
}

type PostModel struct {
	TableName string
}

func (pm *PostModel) Find(id int64) *Post {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople where idpost=?"

	stmt, err := db.HgSql.Prepare(sql)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	row := stmt.QueryRow(id)

	var post Post
	err = row.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum, &post.Title,
		&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
		&post.Class.IdpostClass, &post.Class.Parent, &post.Class.Name)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &post
}

func (pm *PostModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Post, int) {
	countSql := "select count(*) as total " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople  "

	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople "

	orderby := "order by post.create_time desc limit ?,?"

	if len(agrs) > 0 {
		sql = sql + " where "
		countSql = countSql + " where "
	} else {
		sql = sql + " and "
		countSql = countSql + " and "
	}

	for k, v := range agrs {
		sql = sql + " " + k + v + " and"
		countSql = countSql + " " + k + v + " and"
	}

	sql = sql + " 1=1 " + orderby
	countSql = countSql + " 1=1 "

	countRow := db.HgSql.QueryRow(countSql)
	count := 0
	err := countRow.Scan(&count)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	stmt, err := db.HgSql.Prepare(sql)
	rows, err := stmt.Query((page-1)*pageSize, pageSize)

	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	var posts []Post
	for rows.Next() {

		post := Post{}
		err = rows.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum, &post.Title,
			&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
			&post.Class.IdpostClass, &post.Class.Parent, &post.Class.Name)
		if err == nil {
			posts = append(posts, post)
		} else {
			fmt.Println(err)
		}
	}

	return posts, count
}

func (pm *PostModel) Insert(post Post) *Post {
	stmt, err := db.HgSql.Prepare("INSERT post SET idpeople=?,content=?,idpost_class=?,reprint_from=?,reprint_url=?,read_num=?,reply_num=?,title=?,create_time=now()")
	res, err := stmt.Exec(post.Author.Idpeople, post.Content, post.Class.IdpostClass, post.ReprintFrom, post.ReprintUrl, post.ReadNum, post.ReplyNum, post.Title)

	insertId, err := res.LastInsertId()
	if err != nil {
		return nil
	}

	post.Idpost = insertId
	return &post
}

func (pm *PostModel) UpdateReadNum(post Post) {
	stmt, err := db.HgSql.Prepare("update post set read_num = read_num+1 where idpost=?")
	_, err = stmt.Exec(post.Idpost)
	if err != nil {
		fmt.Println(err)
	}
}

func (pm *PostModel) UpdateReplyNum(post Post) {
	stmt, err := db.HgSql.Prepare("update post set reply_num = reply_num+1 where idpost=?")
	_, err = stmt.Exec(post.Idpost)
	if err != nil {
		fmt.Println(err)
	}
}
