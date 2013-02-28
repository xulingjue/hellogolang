package people

import (
	"fmt"
	hgHelper "hellogolang/system/helper"
	"net/http"
	"text/template"
)

var (
	pm PeopleModel
)

func init() {
	pm = PeopleModel{"people"}
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

		//写入session
	}
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
