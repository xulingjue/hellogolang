package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type Post struct {
	Idpost      uint64
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

func (pm *PostModel) Find(id uint64) *Post {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople where idpost=%d"

	row, _, err := db.HgSql.QueryFirst(sql, id)

	var post Post
	if err != nil {
		return nil
	}

	post.Idpost = row.Uint64(0)
	post.Content = row.Str(1)
	post.CreateTime = row.Str(2)
	post.ReprintFrom = row.Str(3)
	post.ReprintUrl = row.Str(4)
	post.ReadNum = row.Int(5)
	post.ReplyNum = row.Int(6)
	post.Title = row.Str(7)
	post.Author.Idpeople = row.Uint64(8)
	post.Author.Name = row.Str(9)
	post.Author.Avatar = row.Str(10)
	post.Class.IdpostClass = row.Uint64(11)
	post.Class.Parent = row.Uint64(12)
	post.Class.Name = row.Str(13)

	return &post
}

func (pm *PostModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Post, int) {
	countSql := "select count(*) as total " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople  "

	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople "

	orderby := "order by post.create_time desc limit %d,%d"

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

	countRow, _, err := db.HgSql.QueryFirst(countSql)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	count := countRow.Int(0)

	rows, _, err := db.HgSql.Query(sql, (page-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	var posts []Post
	for _, row := range rows {
		var post Post
		// You can get converted value
		post.Idpost = row.Uint64(0)
		post.Content = row.Str(1)
		post.CreateTime = row.Str(2)
		post.ReprintFrom = row.Str(3)
		post.ReprintUrl = row.Str(4)
		post.ReadNum = row.Int(5)
		post.ReplyNum = row.Int(6)
		post.Title = row.Str(7)
		post.Author.Idpeople = row.Uint64(8)
		post.Author.Name = row.Str(9)
		post.Author.Avatar = row.Str(10)
		post.Class.IdpostClass = row.Uint64(11)
		post.Class.Parent = row.Uint64(12)
		post.Class.Name = row.Str(13)

		posts = append(posts, post)
	}

	return posts, count
}

func (pm *PostModel) Insert(post Post) (uint64, error) {
	stmt, err := db.HgSql.Prepare("INSERT post SET idpeople=?,content=?,idpost_class=?,reprint_from=?,reprint_url=?,read_num=?,reply_num=?,title=?,create_time=now()")

	_, res, err := stmt.Exec(post.Author.Idpeople, post.Content, post.Class.IdpostClass, post.ReprintFrom, post.ReprintUrl, post.ReadNum, post.ReplyNum, post.Title)

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	return res.InsertId(), nil
}
