package main

import (
	"swiftycareer-go/server"
	"swiftycareer-go/server/database"
)

func main() {
	database.ConnectToMongo()
	server.CreateServer()
}
