package solutions

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
	problemes "josephdev.io/blog_api/Problemes"
)

// Add adds a new record in DB
func Add(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	var record Solution
	err = json.Unmarshal(data, &record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	if strings.TrimSpace(record.ProblemID) == "" || strings.TrimSpace(record.Description) == "" || strings.TrimSpace(record.DateExit) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "Please fill in all the fields.")
		return
	}

	/*// ─── WE MAKE SURE THAT THE RECORD DOESN'T ALREADY EXIST ──────────────────────────────────────────────
	var recordCheck Solution
	filter := bson.M{"$or": bson.A{
		bson.M{"_Number": record.ProblemID},
	}}
	err = LocalDB.Collection(CollectionName).FindOne(context.TODO(), filter).Decode(&recordCheck)
	if err != mongo.ErrNoDocuments {
		goutilities.APIResult(&w, http.StatusBadRequest, goutilities.ErrorsRender("The technicien already exists", err))
		return
	}*/

	// ─── WE ADD THE GROUP IN DB ───────────────────────────────────────────────────
	t, _ := time.Now().UTC().MarshalText()
	record.ID = uuid.Must(uuid.NewRandom()).String()
	record.CreatedAt = string(t)
	record.UpdatedAt = string(t)

	_, err = LocalDB.Collection(CollectionName).InsertOne(context.TODO(), record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't create the solution, please retry", err))
		return
	}
	// We update the status of the equipment
	var problem problemes.Problemes
	filter := bson.M{"_id": record.ProblemID}
	err = LocalDB.Collection(problemes.CollectionName).FindOne(context.TODO(), filter).Decode(&problem)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't find the problem", err))
		return
	}
	var affec affectations.Affectations
	filter = bson.M{"_id": problem.AffecId}
	update := bson.M{}
	update["_DateFin"] = string(t)
	err = LocalDB.Collection(affectations.CollectionName).FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}).Decode(&affec)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't find the affectation", err))
		return
	}
	filter = bson.M{"_id": affec.EquipId}
	_, err = LocalDB.Collection(equipement.CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"_EtatUti": "non-affecte"}})
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull solution from db", err))
		return
	}

	records := []Solution{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read solution from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return solution data", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't delete the solution", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the solution", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull solution from db", err))
		return
	}

	records := []Solution{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read solution from db", err))
		return
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return solution data", err))
		return
	}

	goutilities.APIBody(&w, &bufData)
}
