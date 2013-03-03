package people

import (
	"code.google.com/p/gorilla/sessions"
	"fmt"
	hgHelper "hellogolang/system/helper"
	"net/http"
	"text/template"
)

var (
	pm    PeopleModel
	store *sessions.CookieStore
)

func init() {
	pm = PeopleModel{"people"}
	store = sessions.NewCookieStore([]byte("hellogolang.org"))
}

/*
 * 登录操作
 */
func Login(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tmpl, _ := template.ParseFiles("template/front/header.tmpl",
			"template/front/people-login.tmpl",
			"template/front/footer.tmpl")

		tmpl.ExecuteTemplate(rw, "people-login", nil)
		tmpl.Execute(rw, nil)
	} else if req.Method == "POST" {
		req.ParseForm()
		for k, v := range req.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", v)
		}
	}
}

/*
 * 注册操作
 */
func Regist(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tmpl, _ := template.New("registView").ParseFiles(
			"template/front/header.tmpl",
			"template/front/people-regist.tmpl",
			"template/front/footer.tmpl")

		js := []string{
			"front/people/people-regist.js"}
		extra_js := []string{
			"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-regist", map[string]interface{}{"baseUrl": hgHelper.GetConfig("base_url"), "js": js, "extra_hs": extra_js})
	} else {
		req.ParseForm()
		var people People
		people.name = req.FormValue("name")
		people.email = req.FormValue("email")
		people.password = req.FormValue("password")
		pm.Insert(people)

	}
}

func SessionSet(rw http.ResponseWriter, req *http.Request) {
	//写入session
	session, _ := store.Get(req, "hellogolang.org-user")
	// Set some session values.
	session.Values["name"] = "xulingjue"
	session.Values["email"] = "xulingjue@email"
	session.Values["idpeople"] = 1

	// Save it.
	session.Save(req, rw)
}

func SessionGet(rw http.ResponseWriter, req *http.Request) {
	//写入session
	session, _ := store.Get(req, "hellogolang.org-user")
	// Set some session values.
	fmt.Println(session.Values["name"])
}

/*
 * ajax 判断用户是否存在 
 */
func AjaxIsExist(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.FormValue("name")
	email := req.FormValue("email")

	if name != "" {
		people, _ := pm.FindByName(name)
		if people.idpeople != 0 {
			fmt.Println("no people")
		}
	}

	if email != "" {
		people, _ := pm.FindByEmail(email)
		if people.idpeople != 0 {
			fmt.Println("no people")
		}
	}

	fmt.Println("has people")
}
