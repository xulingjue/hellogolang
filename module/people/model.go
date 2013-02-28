package people

import (
	//"fmt"
	db "hellogolang/system/database"
)

type PeopleModel struct {
	tableName string
}

func (pm *PeopleModel) Find(id int) (People, error) {
	row := db.HgSql.QueryRow("select * from people where idpeople=?", id)
	var people People
	err := row.Scan(&people.idpeople, &people.name, &people.email, &people.phone, &people.avatar, &people.lastLogin, &people.createTime, &people.fansnum, &people.favnum, &people.password, &people.qq)
	if err != nil {
		return people, err
	}
	return people, nil
}

func (pm *PeopleModel) Insert(people People) (int64, error) {
	stmt, err := db.HgSql.Prepare("INSERT people SET name=?,email=?,phone=?,avatar=?,create_time=now(),fansnum=?,favnum=?,password=?,qq=?")
	res, err := stmt.Exec(people.name, people.email, people.phone, people.avatar, people.fansnum, people.favnum, people.password, people.qq)
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pm *PeopleModel) FindByName(name string) (People, error) {
	row := db.HgSql.QueryRow("select * from people where name=?", name)
	var people People
	err := row.Scan(&people.idpeople, &people.name, &people.email, &people.phone, &people.avatar, &people.lastLogin, &people.createTime, &people.fansnum, &people.favnum, &people.password, &people.qq)
	if err != nil {
		panic(err)
		return people, err
	}
	return people, nil
}

func (pm *PeopleModel) FindByEmail(email string) (People, error) {
	row := db.HgSql.QueryRow("select * from people where email=?", email)
	var people People
	err := row.Scan(&people.idpeople, &people.name, &people.email, &people.phone, &people.avatar, &people.lastLogin, &people.createTime, &people.fansnum, &people.favnum, &people.password, &people.qq)
	if err != nil {
		panic(err)
		return people, err
	}
	return people, nil
}
