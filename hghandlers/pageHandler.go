/*
	加载静态文件
*/
package hghandlers

import (
	"net/http"
)

func SourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
