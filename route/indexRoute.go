package route

import (
	"fmt"
	"net/http"
)

func AddRoutes(m *http.ServeMux) {
	addIndexHandler(m)
	addTestHandler(m)
}

func addIndexHandler(m *http.ServeMux) {
	m.HandleFunc("/", indexHandler)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}
