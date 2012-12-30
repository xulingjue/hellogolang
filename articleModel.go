package main

import (
	"fmt"
)

type article struct {
	Idarticle    int
	Content      string
	CreateTime   string
	Readnum      int
	Favnum       int
	articleClass articleClass
}

type articleClass struct {
	IdarticleClass int
	Name           string
}

//查询数据
func findArticleList() {
	//rows, err := 
	//hgDb.Query("SELECT 1")
	fmt.Println(config["port"])
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	//articleItem := new(article)
	// 	fmt.Println("test")
	// 	//fmt.Println(articleItem.Idarticle)
	// 	//fmt.Println(articleItem.CreateTime)
	// }

}
