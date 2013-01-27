package people

import (
	//hgHelper "hellogolang/core/helper"
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
	tmpl, _ := template.New("registView").Funcs(template.FuncMap{"emailDeal": lingjueTest}).ParseFiles(
		"template/front/header.tmpl",
		"template/front/people-regist.tmpl",
		"template/front/footer.tmpl")
	//t, _ := template.Funcs(template.FuncMap{"emailDeal": lingjueTest})
	tmpl.ExecuteTemplate(rw, "people-regist", map[string]interface{}{"form": "form"})
	tmpl.Execute(rw, nil)
}

/*
 * 执行注册操作
 */
func Regist(rw http.ResponseWriter, req *http.Request) {
}

func lingjueTest() string {
	return "-lingjue"
}
