package front

import (
	"encoding/json"
	"fmt"
	hgQiniu "hellogolang/HooGL/qiniu"
	hgTemplate "hellogolang/HooGL/template"
	"hellogolang/application/model"
	"io"
	"net/http"
	"os"
	"text/template"
)

func PeopleAvatarEdit(rw http.ResponseWriter, req *http.Request) {
	people := isLogin(req)

	if people == nil {
		errorMessage(rw, req)
		return
	}

	if req.Method == "GET" {
		tmpl := template.New("people-ucenter-avatar")
		tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/people-ucenter-avatar.tmpl",
			"template/front/footer.tmpl",
			"template/front/people-ucenter-sidebar.tmpl")

		tmplInfo := hgTemplate.TmplInfo{}
		tmplInfo.AddData("people", people)
		tmplInfo.Js = []string{
			"js/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-ucenter-avatar", map[string]interface{}{"tmplInfo": tmplInfo})

	} else if req.Method == "POST" {

		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f, _ := os.OpenFile("d://"+handler.Filename, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		f.Seek(0, 0)

		ext := hgQiniu.GetExt(handler.Filename)
		fk := hgQiniu.GetFk(people.Idpeople)

		err = hgQiniu.UploadAvatar(f, fk+"."+ext)
		if err == nil {
			people.Avatar = fk + "." + ext
			peopleModel.Update(*people)
			http.Redirect(rw, req, "/people/ucenter/", http.StatusFound)
		}
	}
}

func PeoplePasswdEdit(rw http.ResponseWriter, req *http.Request) {

}

func PeopleMessageEdit(rw http.ResponseWriter, req *http.Request) {
	people := isLogin(req)

	if people == nil {
		errorMessage(rw, req)
		return
	}

	if req.Method == "GET" {
		tmpl := template.New("people-ucenter")
		tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/people-ucenter-edit.tmpl",
			"template/front/footer.tmpl",
			"template/front/people-ucenter-sidebar.tmpl")

		tmplInfo := hgTemplate.TmplInfo{}
		tmplInfo.AddData("people", people)
		tmplInfo.Js = []string{
			"js/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-ucenter-edit", map[string]interface{}{"tmplInfo": tmplInfo})
	} else if req.Method == "POST" {
		req.ParseForm()

		people.Phone = req.FormValue("name")
		people.QQ = req.FormValue("qq")
		people.Homepage = req.FormValue("homepage")
		people.Company = req.FormValue("company")
		people.Signature = req.FormValue("signature")
		people.Resume = req.FormValue("resume")

		peopleModel.Update(*people)

		http.Redirect(rw, req, "/people/ucenter/", http.StatusFound)
	}
}

func PeopleUcenter(rw http.ResponseWriter, req *http.Request) {
	people := isLogin(req)
	tmpl := template.New("people-ucenter")
	tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/people-ucenter.tmpl",
		"template/front/footer.tmpl",
		"template/front/people-ucenter-sidebar.tmpl")

	tmplInfo := hgTemplate.TmplInfo{}
	tmplInfo.AddData("people", people)
	tmplInfo.Js = []string{
		"js/jquery.validate.js"}

	tmpl.ExecuteTemplate(rw, "people-ucenter", map[string]interface{}{"tmplInfo": tmplInfo})

}

func PeopleAjaxLogin(rw http.ResponseWriter, req *http.Request) {
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
			return
		}
		fmt.Fprintf(rw, string(b))
	}

}

/*
 *退出
 */
func PeopleLogout(rw http.ResponseWriter, req *http.Request) { //ok
	session, _ := store.Get(req, "hellogolang.org-user")
	session.Values["name"] = nil
	session.Values["email"] = nil
	session.Values["idpeople"] = nil

	// Save it.
	session.Save(req, rw)

	http.Redirect(rw, req, "/", http.StatusFound)
}

/*
 * 注册操作
 */
func PeopleRegist(rw http.ResponseWriter, req *http.Request) {
	//检测是否已经登录
	//people := isLogin(req)

	if req.Method == "GET" {
		tmpl := template.New("registView")
		tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/people-regist.tmpl",
			"template/front/footer.tmpl")

		tmplInfo := hgTemplate.TmplInfo{}
		tmplInfo.Js = []string{
			"js/jquery.validate.js"}

		tmpl.ExecuteTemplate(rw, "people-regist", map[string]interface{}{"tmplInfo": tmplInfo})
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

/*
 *判断用户是否登录
 */
func isLogin(req *http.Request) *model.People {
	session, _ := store.Get(req, "hellogolang.org-user")
	if session.Values["idpeople"] != nil {
		idpeople := session.Values["idpeople"].(int64)
		people := peopleModel.Find(idpeople)

		return people
	}
	return nil
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
