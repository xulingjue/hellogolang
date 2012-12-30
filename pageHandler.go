/*
	加载静态文件
*/
package main

import (
	"net/http"
)

func sourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
