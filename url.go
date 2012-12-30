/*
URL和Handler的Mapping
*/
package main

import (
	"net/http"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	/*soruce*/
	"/static/": sourceHandler,

	/*article*/
	"/":        articlePageHandler,
	"/article": articleItemHandler,
}
