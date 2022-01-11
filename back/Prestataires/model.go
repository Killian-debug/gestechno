package prestataires

import "go.mongodb.org/mongo-driver/mongo"

type Prestataires struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	Name   string `json:"_Name" bson:"_Name"`
	Number string `json:"_Number" bson:"_Number"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Prestataires"
