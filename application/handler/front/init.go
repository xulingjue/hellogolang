package front

import (
	"code.google.com/p/gorilla/sessions"
	"hellogolang/application/model"
)

var (
	peopleModel   model.PeopleModel
	postModel     model.PostModel
	postTypeModel model.PostTypeModel
	store         *sessions.CookieStore
	siteInfo      SiteInfo
)

func init() {
	peopleModel = model.PeopleModel{"people"}
	store = sessions.NewCookieStore([]byte("hellogolang.org"))
	postTypeModel = model.PostTypeModel{"post_type"}
	postModel = model.PostModel{"post"}
}

/*
 *界面辅助类
 */
type SiteInfo struct {
	BaseUrl    string
	CurrentNav string

	ExtraJs  []string
	Js       []string
	ExtraCss []string
	Css      []string
}
