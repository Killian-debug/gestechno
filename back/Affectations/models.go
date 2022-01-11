package affectations

import "go.mongodb.org/mongo-driver/mongo"

type Affectations struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	EquipId   string `json:"_EquipId" bson:"_EquipId"`
	UserId    string `json:"_UserId" bson:"_UserId"`
	DateDebut string `json:"_DateDebut" bson:"_DateDebut"`
	DateFin   string `json:"_DateFin" bson:"_DateFin"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Affectations"
