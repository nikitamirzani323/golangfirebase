package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()

	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: "https://wigo-29d5f-default-rtdb.asia-southeast1.firebasedatabase.app/room",
	}

	// fetch service account key
	opt := option.WithCredentialsFile("wigo-29d5f-firebase-adminsdk-20j5g-dd81e93c5d.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	// add/update data to firebase DB
	SaveDataToFirebaseDB(client)

	// retrieve data from firebase DB
	GetDataFromFirebaseDB(client)

	// delete data from firebase DB
	// DeleteDataFromFirebaseDB(client)
}

func SaveDataToFirebaseDB(client *db.Client) {
	// create ref at path user_scores/:userId
	ref := client.NewRef("room")

	if err := ref.Set(context.TODO(), map[string]interface{}{
		"id": 1, "client": 3, "total": 5000}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("score added/updated successfully!")
}

func GetDataFromFirebaseDB(client *db.Client) {
	type UserData struct {
		Id     int `json:"id"`
		Client int `json:"client"`
		Total  int `json:"total"`
	}

	// get database reference to user score
	ref := client.NewRef("room")

	// read from user_scores using ref
	var s UserData
	if err := ref.Get(context.TODO(), &s); err != nil {
		log.Fatalln("error in reading from firebase DB: ", err)
	}
	fmt.Println("retrieved user's Id is: ", s.Id)
	fmt.Println("retrieved user's client is: ", s.Client)
	fmt.Println("retrieved user's Total is: ", s.Total)
}

func DeleteDataFromFirebaseDB(client *db.Client) {
	ref := client.NewRef("room/1")

	if err := ref.Delete(context.TODO()); err != nil {
		log.Fatalln("error in deleting ref: ", err)
	}
	fmt.Println("user's score deleted successfully:)")
}
