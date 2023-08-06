package casel

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ValidateRequest(req *http.Request, res http.ResponseWriter, url, method string) (int, bool) {

	if req.URL.Path != url {
		res.WriteHeader(http.StatusNotFound)
		log.Println("404 ❌ - Page not found ", req.URL)
		return 404, false
	}

	if req.Method != method {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "%s", "Error - Method not allowed")
		log.Printf("405 ❌ - Method not allowed %s - %s on URL : %s\n", req.Method, method, url)
		return 405, false
	}

	return 0, true
}

func RenderPage(pagePath string, data any, res http.ResponseWriter) {
	files := []string{"templates/base.html", "templates/" + pagePath + ".html"}
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("🚨 " + err.Error())
	} else {
		tpl.Execute(res, data)
	}
}
