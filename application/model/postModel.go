package model

import (
	"fmt"
	db "hellogolang/system/database"
)

type Post struct {
	Idpost      int64
	Idpeople    int64
	Content     string
	IdpostClass int64
	CreateTime  string
	ReprintFrom string
	ReprintUrl  string
	ReadNum     int
	ReplyNum    int
	Title       string

	Author People
	Class  PostClass
}

type PostModel struct {
	TableName string
}

func (pm *PostModel) Find(id int64) (Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople where idpost=%d"

	rows, _, err := db.HgSql.Query(sql, id)
	var post Post

	if err != nil {
		fmt.Println(err)
		return post, err
	}

	for _, row := range rows {
		post.Idpost = row.Int64(0)
		post.Content = row.Str(1)
		post.CreateTime = row.Str(2)
		post.ReprintFrom = row.Str(3)
		post.ReprintUrl = row.Str(4)
		post.ReadNum = row.Int(5)
		post.ReplyNum = row.Int(6)
		post.Title = row.Str(7)
		post.Author.Idpeople = row.Int64(8)
		post.Author.Name = row.Str(9)
		post.Author.Avatar = row.Str(10)
		post.Class.IdpostClass = row.Int64(11)
		post.Class.Parent = row.Int64(12)
		post.Class.Name = row.Str(13)

		return post, nil
	}

	return post, nil
}

func (pm *PostModel) FindAllCount(agrs map[string]string) (int, error) {
	sql := "select count(*) as total " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople  "

	if len(agrs) > 0 {
		sql = sql + " where "
	} else {
		sql = sql + " and "
	}

	for k, v := range agrs {
		sql = sql + " " + k + v + " and"
	}

	sql = sql + " 1=1 "

	rows, _, err := db.HgSql.Query(sql)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for _, row := range rows {
		return row.Int(0), nil
	}

	return 0, nil
}

func (pm *PostModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople "

	orderby := "order by post.create_time desc limit %d,%d"

	if len(agrs) > 0 {
		sql = sql + " where "
	} else {
		sql = sql + " and "
	}

	for k, v := range agrs {
		sql = sql + " " + k + v + " and"
	}

	sql = sql + " 1=1 " + orderby

	rows, _, err := db.HgSql.Query(sql, (page-1)*pageSize, pageSize)

	var posts []Post

	if err != nil {
		fmt.Println(err)
	}

	for _, row := range rows {
		var post Post
		// You can get converted value
		post.Idpost = row.Int64(0)
		post.Content = row.Str(1)
		post.CreateTime = row.Str(2)
		post.ReprintFrom = row.Str(3)
		post.ReprintUrl = row.Str(4)
		post.ReadNum = row.Int(5)
		post.ReplyNum = row.Int(6)
		post.Title = row.Str(7)
		post.Author.Idpeople = row.Int64(8)
		post.Author.Name = row.Str(9)
		post.Author.Avatar = row.Str(10)
		post.Class.IdpostClass = row.Int64(11)
		post.Class.Parent = row.Int64(12)
		post.Class.Name = row.Str(13)

		posts = append(posts, post)
	}

	return posts, nil
}

func (pm *PostModel) Insert(post Post) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT post SET idpeople=?,content=?,idpost_class=?,reprint_from=?,reprint_url=?,read_num=?,reply_num=?,title=?,create_time=now()")

	_, res, err := stmt.Exec(post.Idpeople, post.Content, post.IdpostClass, post.ReprintFrom, post.ReprintUrl, post.ReadNum, post.ReplyNum, post.Title)

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	return res.InsertId(), nil
}
