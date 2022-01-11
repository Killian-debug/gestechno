package equipments

import "go.mongodb.org/mongo-driver/mongo"

type Equipments struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	CaracTech string `json:"_CaracTech" bson:"_CaracTech"`
	Marque    string `json:"_Marque" bson:"_Marque"`
	Modele    string `json:"_Modele" bson:"_Modele"`
	EtatAcha  string `json:"_EtatAcha" bson:"_EtatAcha"`
	PrixAcha  int    `json:"_PrixAcha" bson:"_PrixAcha"`
	DateAcha  string `json:"_DateAcha" bson:"_DateAcha"`
	EtatUti   string `json:"_EtatUti" bson:"_EtatUti"` // affecte -- non-affecte -- hors du parc -- panne
	FourniId  string `json:"_FourniId" bson:"_FourniId"`
}

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Equipments"
