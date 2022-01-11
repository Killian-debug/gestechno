package techniciens

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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add adds a new record in DB
func Add(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	var record Technicien
	err = json.Unmarshal(data, &record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	if strings.TrimSpace(record.FirstName) == "" || strings.TrimSpace(record.LastName) == "" || strings.TrimSpace(record.Number) == "" || strings.TrimSpace(record.Login) == "" || strings.TrimSpace(record.Password) == "" /*|| record.Level == 0*/ {
		goutilities.APIResult(&w, http.StatusBadRequest, "Please fill in all the fields.")
		return
	}

	// ─── WE MAKE SURE THAT THE RECORD DOESN'T ALREADY EXIST ──────────────────────────────────────────────
	var recordCheck Technicien
	filter := bson.M{"$or": bson.A{
		bson.M{"_Number": record.Number},
	}}
	err = LocalDB.Collection(CollectionName).FindOne(context.TODO(), filter).Decode(&recordCheck)
	if err != mongo.ErrNoDocuments {
		goutilities.APIResult(&w, http.StatusBadRequest, goutilities.ErrorsRender("The technicien already exists", err))
		return
	}

	// ─── WE ADD THE GROUP IN DB ───────────────────────────────────────────────────
	t, _ := time.Now().UTC().MarshalText()
	record.ID = uuid.Must(uuid.NewRandom()).String()
	record.CreatedAt = string(t)
	record.UpdatedAt = string(t)

	_, err = LocalDB.Collection(CollectionName).InsertOne(context.TODO(), record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't create the technicien, please retry", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull technicien from db", err))
		return
	}

	records := []Technicien{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read technicien from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return technicien data", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't delete the technicien", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the technicien", err))
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
		bson.M{"_FirstName": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_LastName": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_Login": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_Level": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
	}}
	searchOptions := options.Find().SetSort(bson.M{"_CreatedAt": -1}).SetLimit(10000).SetProjection(bson.M{"_Password": 0})
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter, searchOptions)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull technicien from db", err))
		return
	}

	records := []Technicien{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read technicien from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return technicien data", err))
		return
	}

	goutilities.APIBody(&w, &bufData)
}
