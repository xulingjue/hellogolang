package people

import (
	"fmt"
	//hgHelper "hellogolang/core/helper"
	"net/http"
	//"text/template"
)

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
