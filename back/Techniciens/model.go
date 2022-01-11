package techniciens

import "go.mongodb.org/mongo-driver/mongo"

type Technicien struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	FirstName string `json:"_FirstName" bson:"_FirstName"`
	LastName  string `json:"_LastName" bson:"_LastName"`
	Login     string `json:"_Login" bson:"_Login"`
	Password  string `json:"_Password" bson:"_Password"`
	Number    string `json:"_Number" bson:"_Number"`
	Level     int    `json:"_Level" bson:"_Level"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Technicien in DB
const CollectionName string = "__Technicien"
