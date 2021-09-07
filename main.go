package main

import (
	"context"
	"swiftycareer-go/server"
	"swiftycareer-go/server/database"
)

func main() {
	database.ConnectToMongo()
	server.CreateServer()

	// disconnect the database
	defer func() {
		if err := database.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}
