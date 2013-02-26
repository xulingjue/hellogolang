package people

import (
	db "hellogolang/system/database"
)

type PeopleModel struct {
	tableName string
}

func (pm *PeopleModel) Find(id int) string {
	db.Query("select")
	return ""
}
