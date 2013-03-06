package front

import (
	"code.google.com/p/gorilla/sessions"
	"hellogolang/application/model"
)

var (
	peopleModel model.PeopleModel
	store       *sessions.CookieStore
	siteInfo    SiteInfo
)

func init() {
	peopleModel = model.PeopleModel{"people"}
	//hgHandler = HgHandler{}
	store = sessions.NewCookieStore([]byte("hellogolang.org"))
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
