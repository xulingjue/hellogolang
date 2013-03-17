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
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople where idpost=?"

	row := db.HgSql.QueryRow(sql, id)
	var post Post
	err := row.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum, &post.Title,
		&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
		&post.Class.IdpostClass, &post.Class.Parent, &post.Class.Name,
	)
	if err != nil {
		fmt.Print(err)
		return post, err
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

	var num int

	rows := db.HgSql.QueryRow(sql)
	err := rows.Scan(&num)

	if err != nil {
		return num, err
	}

	return num, nil
}

func (pm *PostModel) FindAll(page int, pageSize int, agrs map[string]string) ([]Post, error) {
	sql := "select post.idpost,post.content,post.create_time,post.reprint_from,post.reprint_url,post.read_num,post.reply_num,post.title," +
		"people.idpeople,people.name,people.avatar," +
		"post_class.idpost_class,post_class.parent,post_class.name " +
		"from post left join post_class on post_class.idpost_class=post.idpost_class left join people on people.idpeople=post.idpeople "

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
			err := rows.Scan(&post.Idpost, &post.Content, &post.CreateTime, &post.ReprintFrom, &post.ReprintUrl, &post.ReadNum, &post.ReplyNum, &post.Title,
				&post.Author.Idpeople, &post.Author.Name, &post.Author.Avatar,
				&post.Class.IdpostClass, &post.Class.Parent, &post.Class.Name,
			)
			if err == nil {
				posts = append(posts, post)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
	}
	return posts, nil
}

func (pm *PostModel) Insert(post Post) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT post SET idpeople=?,content=?,idpost_class=?,reprint_from=?,reprint_url=?,read_num=?,reply_num=?,title=?,create_time=now()")
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	res, err := stmt.Exec(post.Idpeople, post.Content, post.IdpostClass, post.ReprintFrom, post.ReprintUrl, post.ReadNum, post.ReplyNum, post.Title)

	if err != nil {
		fmt.Println(post.IdpostClass)
		fmt.Println(err)
		return 0, nil
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}
	return id, nil
}
