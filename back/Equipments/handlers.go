package equipments

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
)

// Add adds a new record in DB
func Add(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	var record Equipments
	err = json.Unmarshal(data, &record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	if strings.TrimSpace(record.FourniId) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "Please fill the fournisseur ID.")
		return
	}
	record.EtatUti = "non-affecte"
	// ─── WE ADD THE GROUP IN DB ───────────────────────────────────────────────────
	t, _ := time.Now().UTC().MarshalText()
	record.ID = uuid.Must(uuid.NewRandom()).String()
	record.CreatedAt = string(t)
	record.UpdatedAt = string(t)

	_, err = LocalDB.Collection(CollectionName).InsertOne(context.TODO(), record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't create the Equipments, please retry", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Equipments from db", err))
		return
	}

	records := []Equipments{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Equipments from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return Equipments data", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't delete the Equipments", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the Equipments", err))
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
		bson.M{"_CaracTech": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_Marque": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_Modele": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_EtatAcha": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_PrixAcha": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_DateAcha": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
		bson.M{"_EtatUti": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
	}}
	searchOptions := options.Find().SetSort(bson.M{"_CreatedAt": -1}).SetLimit(10000).SetProjection(bson.M{"_Password": 0})
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter, searchOptions)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Equipments from db", err))
		return
	}

	records := []Equipments{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Equipments from db", err))
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

// ListWithFilter list equipement by a filter
func ListWithFilter(w http.ResponseWriter, r *http.Request) {
	// We make sure the search is not empty
	vars := mux.Vars(r)
	sSearch, ok := vars["searchBy"]
	if !ok {
		goutilities.APIResult(&w, http.StatusBadRequest, "The filter is empty.")
		return
	}

	if strings.TrimSpace(sSearch) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "The filter is empty.")
		return
	}

	if strings.TrimSpace(sSearch) != "non-affecte" && strings.TrimSpace(sSearch) != "affecte" && strings.TrimSpace(sSearch) != "panne" && strings.TrimSpace(sSearch) == "hors-parc" {
		goutilities.APIResult(&w, http.StatusBadRequest, "The filter is empty.")
		return
	}

	// ─── WE MAKE SURE THAT THE GROUP DOESN'T ALREADY EXIST ──────────────────────────────────────────────
	filter := bson.M{"_EtatUti": sSearch}
	searchOptions := options.Find().SetSort(bson.M{"_CreatedAt": -1}).SetLimit(10000).SetProjection(bson.M{"_Password": 0})
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter, searchOptions)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Equipments from db", err))
		return
	}

	records := []Equipments{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Equipments from db", err))
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

func ExitOfThePark(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sID := vars["id"]
	filter := bson.M{"_id": sID}
	_, err := LocalDB.Collection(CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"_EtatUti": "hors-du-park"}})
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the Equipments", err))
		return
	}
	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	goutilities.APIBodyString(&w, `{"result": "ok"}`)
}
