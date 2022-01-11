package services

import "go.mongodb.org/mongo-driver/mongo"

type Services struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	DateExp       string `json:"_DateExp" bson:"_DateExp"`
	Ref           string `json:"_Ref" bson:"_Ref"`
	Type          string `json:"_Type" bson:"_Type"`
	PrestataireID string `json:"_PrestataireID" bson:"_PrestataireID"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Services"
