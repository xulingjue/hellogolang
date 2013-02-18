package people

import (
	"fmt"
	hgHelper "hellogolang/system/helper"
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
	} else {

	}
}

/*
 * 注册操作
 */
func Regist(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
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
		for k, v := range req.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", v)
		}
	}
}

/*
 * ajax 判断用户是否存在 
 */
func AjaxIsExist(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form["name"]
	email := req.Form["email"]

	if name != nil {
		fmt.Println("name")
	}

	if email != nil {
		fmt.Println("email")
	}
}
