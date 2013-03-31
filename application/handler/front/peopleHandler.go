package front

import (
	"encoding/json"
	"fmt"
	"hellogolang/application/model"
	"hellogolang/system/tmplfunc"
	"net/http"
	"text/template"
)

/*
 * 登录操作  计划改为ajax登录
 */
func Login(rw http.ResponseWriter, req *http.Request) { //ok
	if req.Method == "GET" {
		tmpl := template.New("people-login.tmpl")
		tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/people-login.tmpl",
			"template/front/footer.tmpl")

		siteInfo.CurrentNav = ""

		siteInfo.Js = []string{
			"js/front/people/people-login.js"}
		siteInfo.ExtraJs = []string{
			"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-login", map[string]interface{}{"siteInfo": siteInfo})
	} else if req.Method == "POST" {
		req.ParseForm()
		name := req.FormValue("name")
		password := req.FormValue("password")

		people := peopleModel.FindByName(name)
		if people == nil {
			people = peopleModel.FindByEmail(name)
		}

		if people != nil && people.Password == password {
			session, _ := store.Get(req, "hellogolang.org-user")
			session.Values["name"] = people.Name
			session.Values["email"] = people.Email
			session.Values["idpeople"] = people.Idpeople

			session.Save(req, rw)
			http.Redirect(rw, req, "/", http.StatusFound)
		} else {
			tmpl := template.New("people-login.tmpl")
			tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
			tmpl.ParseFiles(
				"template/front/header.tmpl",
				"template/front/people-login.tmpl",
				"template/front/footer.tmpl")

			siteInfo.Js = []string{
				"js/front/people/people-login.js"}
			siteInfo.ExtraJs = []string{
				"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

			errorMessage := "loginError"
			tmpl.ExecuteTemplate(rw, "people-login", map[string]interface{}{"errorMessage": errorMessage, "siteInfo": siteInfo})
		}
	}
}

func AjaxLogin(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Fprintf(rw, "false")
	} else if req.Method == "POST" {
		req.ParseForm()
		name := req.FormValue("name")
		password := req.FormValue("password")
		people := peopleModel.FindByName(name)
		if people == nil {
			people = peopleModel.FindByEmail(name)
		}

		result := make(map[string]interface{})

		if people != nil && people.Password == password {
			session, _ := store.Get(req, "hellogolang.org-user")
			session.Values["name"] = people.Name
			session.Values["email"] = people.Email
			session.Values["idpeople"] = people.Idpeople
			session.Save(req, rw)
			result["result"] = "success"
			result["people"] = people
		} else {
			result["result"] = "error"
		}

		b, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf("json err:", err)
		}
		fmt.Fprintf(rw, string(b))
	}

}

/*
 * 注册操作
 */
func Regist(rw http.ResponseWriter, req *http.Request) {
	//检测是否已经登录
	//people := isLogin(req)

	if req.Method == "GET" {
		tmpl := template.New("registView")
		tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/people-regist.tmpl",
			"template/front/footer.tmpl")

		siteInfo.Js = []string{
			"js/front/people/people-regist.js"}
		siteInfo.ExtraJs = []string{
			"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-regist", map[string]interface{}{"siteInfo": siteInfo})
	} else if req.Method == "POST" {
		req.ParseForm()
		var people model.People
		people.Name = req.FormValue("name")
		people.Email = req.FormValue("email")
		people.Password = req.FormValue("password")

		if checkRegistMess(people) {
			fmt.Println("start insert ...")
			people := peopleModel.Insert(people)
			if people != nil {
				session, _ := store.Get(req, "hellogolang.org-user")
				session.Values["name"] = people.Name
				session.Values["email"] = people.Email
				session.Values["idpeople"] = people.Idpeople

				session.Save(req, rw)
				http.Redirect(rw, req, "/", http.StatusFound)
			}
		}

	}
}

func checkRegistMess(people model.People) bool {
	//检测用户名、Email地址是否合法

	//检测是否重名
	if people.Name != "" {
		people := peopleModel.FindByName(people.Name)
		if people != nil {
			return false
		}
	}

	if people.Email != "" {
		people := peopleModel.FindByEmail(people.Email)
		if people != nil {
			return false
		}
	}

	return true
}

/*
 *退出
 */
func Logout(rw http.ResponseWriter, req *http.Request) { //ok
	session, _ := store.Get(req, "hellogolang.org-user")
	session.Values["name"] = nil
	session.Values["email"] = nil
	session.Values["idpeople"] = nil

	// Save it.
	session.Save(req, rw)

	http.Redirect(rw, req, "/", http.StatusFound)
}

func SessionSet(rw http.ResponseWriter, req *http.Request) { //ok
	//写入session
	session, _ := store.Get(req, "hellogolang.org-user")
	// Set some session values.
	session.Values["name"] = "xulingjue"
	session.Values["email"] = "xulingjue@email"
	session.Values["idpeople"] = 1

	// Save it.
	session.Save(req, rw)
}

func SessionGet(rw http.ResponseWriter, req *http.Request) { //ok
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
func PeopleAjaxIsExist(rw http.ResponseWriter, req *http.Request) { //ok
	req.ParseForm()
	name := req.FormValue("name")
	email := req.FormValue("email")

	if name != "" {
		people := peopleModel.FindByName(name)
		if people != nil {
			fmt.Fprintf(rw, "false")
			return
		}
	}

	if email != "" {
		people := peopleModel.FindByEmail(email)
		if people != nil {
			fmt.Fprintf(rw, "false")
			return
		}
	}

	fmt.Fprintf(rw, "true")
	return
}
