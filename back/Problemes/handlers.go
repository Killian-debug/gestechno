package problemes

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/polo44/goutilities"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	affectations "josephdev.io/blog_api/Affectations"
	equipement "josephdev.io/blog_api/Equipments"
)

// Add adds a new record in DB
func Add(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	var record Problemes
	err = json.Unmarshal(data, &record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	if strings.TrimSpace(record.DatePanne) == "" || strings.TrimSpace(record.Description) == "" || strings.TrimSpace(record.AffecId) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "Please fill in all the fields.")
		return
	}

	// ─── WE ADD THE GROUP IN DB ───────────────────────────────────────────────────
	t, _ := time.Now().UTC().MarshalText()
	record.ID = uuid.Must(uuid.NewRandom()).String()
	record.CreatedAt = string(t)
	record.UpdatedAt = string(t)

	_, err = LocalDB.Collection(CollectionName).InsertOne(context.TODO(), record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't create the probleme, please retry", err))
		return
	}

	// We update the status of the equipment
	var affec affectations.Affectations
	filter := bson.M{"_id": record.AffecId}
	update := bson.M{}
	update["_DateFin"] = string(t)
	err = LocalDB.Collection(affectations.CollectionName).FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}).Decode(&affec)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't find the affectation", err))
		return
	}
	filter = bson.M{"_id": affec.EquipId}
	_, err = LocalDB.Collection(equipement.CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"_EtatUti": "panne"}})
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the Equipments", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	goutilities.APIBodyString(&w, `{"result": "ok"}`)
}

// List list records in DB
func List(w http.ResponseWriter, r *http.Request) {

	// ─── WE PULL LIST OF RECORDS ────────────────────────────────────────────────────
	filter := bson.D{}
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull problemes from db", err))
		return
	}

	records := []Problemes{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read problemes from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return problemes data", err))
		return
	}

	goutilities.APIBody(&w, &bufData)
}

// Delete deletes a record in DB
func Delete(w http.ResponseWriter, r *http.Request) {

	// ─── WE DELETE THE RECORD ───────────────────────────────────────────────────────
	vars := mux.Vars(r)
	sID := vars["id"]

	filter := bson.M{"_id": sID}
	_, err := LocalDB.Collection(CollectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't delete the probleme", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	goutilities.APIBodyString(&w, `{"result": "ok"}`)
}

// Update updates a record in DB
func Update(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	t, _ := time.Now().UTC().MarshalText()
	update := bson.M{}
	err = json.Unmarshal(data, &update)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}
	update["_UpdatedAt"] = string(t)

	// ─── WE UPDATE THE RECORD ───────────────────────────────────────────────────────
	vars := mux.Vars(r)
	sID := vars["id"]

	filter := bson.M{"_id": sID}
	_, err = LocalDB.Collection(CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the probleme", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	goutilities.APIBodyString(&w, `{"result": "ok"}`)
}

// SeekCustom searches for an info in the db
func SeekCustom(w http.ResponseWriter, r *http.Request) {

	// We make sure the search is not empty
	vars := mux.Vars(r)
	sSearch, ok := vars["search"]
	if !ok {
		goutilities.APIResult(&w, http.StatusBadRequest, "The search is empty.")
		return
	}

	if strings.TrimSpace(sSearch) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "The search is empty.")
		return
	}

	// ─── WE MAKE SURE THAT THE GROUP DOESN'T ALREADY EXIST ──────────────────────────────────────────────
	filter := bson.M{"$or": bson.A{
		bson.M{"_Description": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_DatePanne": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
	}}
	searchOptions := options.Find().SetSort(bson.M{"_CreatedAt": -1}).SetLimit(10000).SetProjection(bson.M{"_Password": 0})
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter, searchOptions)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull problemes from db", err))
		return
	}

	records := []Problemes{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read problemes from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return services data", err))
		return
	}

	goutilities.APIBody(&w, &bufData)
}
