/*
URL和Handler的Mapping
*/

package main

import (
	hgHadlers "hellogolang/hghandlers"
	"net/http"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	/*soruce*/
	"/static/": hgHadlers.SourceHandler,

	/*article*/
	"/":        hgHadlers.ArticlePageHandler,
	"/article": hgHadlers.ArticleItemHandler,
}
