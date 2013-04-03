package front

import (
	"code.google.com/p/gorilla/sessions"
	"hellogolang/application/model"
)

var (
	peopleModel    model.PeopleModel
	postModel      model.PostModel
	postClassModel model.PostClassModel
	store          *sessions.CookieStore
	commentModel   model.CommentModel
)

func init() {
	peopleModel = model.PeopleModel{"people"}
	store = sessions.NewCookieStore([]byte("hellogolang.org"))
	postModel = model.PostModel{"post"}
	postClassModel = model.PostClassModel{"post_class"}
	commentModel = model.CommentModel{"comment"}
}
