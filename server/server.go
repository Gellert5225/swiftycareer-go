package server

import (
	"fmt"
	"net/http"

	"swiftycareer-go/route"
)

func CreateServer() {
	fmt.Println("Server listening on 3000")

	mux := newRouter()

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		panic(err)
	}
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	route.AddRoutes(mux)

	return mux
}
