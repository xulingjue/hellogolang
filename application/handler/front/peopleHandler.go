package front

import (
	"fmt"
	"hellogolang/application/model"
	"net/http"
	"text/template"
)

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
		name := req.FormValue("name")
		password := req.FormValue("password")

		people, _ := peopleModel.FindByName(name)
		if people.Idpeople == 0 {
			people, _ = peopleModel.FindByEmail(name)
		}

		if people.Idpeople != 0 && people.Password == password {
			session, _ := store.Get(req, "hellogolang.org-user")
			session.Values["name"] = people.Name
			session.Values["email"] = people.Email
			session.Values["idpeople"] = people.Idpeople

			session.Save(req, rw)
			Index(rw, req)
		} else {
			tmpl, _ := template.ParseFiles("template/front/header.tmpl",
				"template/front/people-login.tmpl",
				"template/front/footer.tmpl")

			siteInfo.Js = []string{
				"js/front/people/people-regist.js"}
			siteInfo.ExtraJs = []string{
				"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

			errorMessage := "loginError"
			tmpl.ExecuteTemplate(rw, "people-login", map[string]interface{}{"errorMessage": errorMessage, "siteInfo": siteInfo})
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

		siteInfo.Js = []string{
			"js/front/people/people-regist.js"}
		siteInfo.ExtraJs = []string{
			"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-regist", map[string]interface{}{"siteInfo": siteInfo})
	} else {
		req.ParseForm()
		var people model.People
		people.Name = req.FormValue("name")
		people.Email = req.FormValue("email")
		people.Password = req.FormValue("password")
		peopleModel.Insert(people)

	}
}

/*
 *退出
 */
func Logout(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "hellogolang.org-user")
	session.Flashes()
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
	var people model.People
	people.Idpeople = session.Values["idpeople"].(int64)
}

/*
 * ajax 判断用户是否存在 
 */
func PeopleAjaxIsExist(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.FormValue("name")
	email := req.FormValue("email")

	if name != "" {
		people, _ := peopleModel.FindByName(name)
		if people.Idpeople != 0 {
			fmt.Println("no people")
		}
	}

	if email != "" {
		people, _ := peopleModel.FindByEmail(email)
		if people.Idpeople != 0 {
			fmt.Println("no people")
		}
	}

	fmt.Println("has people")
}