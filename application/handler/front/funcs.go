package front

import (
	"hellogolang/application/model"
	"net/http"
)

func isLogin(req *http.Request) *model.People {
	session, _ := store.Get(req, "hellogolang.org-user")
	var people model.People

	if session.Values["idpeople"] != nil {
		people.Idpeople = session.Values["idpeople"].(int64)
		people.Email = session.Values["email"].(string)
		people.Name = session.Values["name"].(string)
		return &people
	}
	return nil
}
