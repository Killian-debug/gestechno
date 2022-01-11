package solutions

import "go.mongodb.org/mongo-driver/mongo"

type Solution struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	ProblemID   string `json:"_ProblemID" bson:"_ProblemID"`
	Description string `json:"_Description" bson:"_Description"`
	DateExit    string `json:"_DateExit" bson:"_DateExit"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Technicien in DB
const CollectionName string = "__Solution"
