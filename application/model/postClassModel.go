package model

import (
//"fmt"
//db "hellogolang/system/database"
)

type PostClass struct {
	IdPostClass int64
	Name        string
	IdPostType  int64
}

type PostClassModel struct {
	TableName string
}