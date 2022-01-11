package problemes

import "go.mongodb.org/mongo-driver/mongo"

type Problemes struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	Description string `json:"_Description" bson:"_Description"`
	DatePanne   string `json:"_DatePanne" bson:"_DatePanne"`
	AffecId     string `json:"_AffecId" bson:"_AffecId"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Problemes"
