package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"time"
)

var client *mongo.Client

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.InitDate = time.Now()

	collection := client.Database("myapp").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
