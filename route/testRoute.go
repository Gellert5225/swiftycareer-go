package route

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"swiftycareer-go/model"
	"swiftycareer-go/server/database"

	"go.mongodb.org/mongo-driver/bson"
)

func addTestHandler(m *http.ServeMux) {
	m.HandleFunc("/testDB", testDBHandler)
}

func testDBHandler(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := database.Database.Collection("User")
	cur, curErr := userCollection.Find(ctx, bson.M{})

	if curErr != nil {
		panic(curErr)
	}

	var users []model.User
	if err := cur.All(ctx, &users); err != nil {
		panic(err)
	}
	// j, _ := json.MarshalIndent(users, "", " ")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(users)
}
