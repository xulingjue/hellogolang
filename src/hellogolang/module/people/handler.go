package people

import (
	"net/http"
	"text/template"
)

/*
 * 加载登录界面
 */
func LoginView(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("template/front/header.tmpl",
		"template/front/people-login.tmpl",
		"template/front/footer.tmpl")

	t.ExecuteTemplate(rw, "people-login", nil)
	t.Execute(rw, nil)
}

/*
 * 执行登录操作
 */
func Login(rw http.ResponseWriter, req *http.Request) {

}

/*
 * 加载注册界面
 */
func RegistView(rw http.ResponseWriter, req *http.Request) {
}

/*
 * 执行注册操作
 */
func Regist(rw http.ResponseWriter, req *http.Request) {
}
