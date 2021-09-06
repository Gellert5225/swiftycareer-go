package server

import (
	"fmt"
	"net/http"
)

func CreateServer() {
	fmt.Println("Server listening on 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
