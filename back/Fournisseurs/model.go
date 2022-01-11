package fournisseurs

import "go.mongodb.org/mongo-driver/mongo"

type Fournisseurs struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	Adresse string `json:"_Adresse" bson:"_Adresse"`
	Number  string `json:"_Number" bson:"_Number"`
	Name    string `json:"_Name" bson:"_Name"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Fournisseurs"
