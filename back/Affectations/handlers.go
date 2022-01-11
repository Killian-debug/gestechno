package affectations

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/polo44/goutilities"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	equipement "josephdev.io/blog_api/Equipments"
	users "josephdev.io/blog_api/Users"
)

type EquipUser struct {
	// Default infos
	EquipID string `json:"_EquipID" bson:"_EquipID"`
	User    string `json:"_User" bson:"_User"`
}

type AffestationData struct {
	User      string `json:"_User" bson:"_User"`
	EquipMarq string `json:"_EquipMarq" bson:"_EquipMarq"`
	EquipMod  string `json:"_EquipMod" bson:"_EquipMod"`
}

type LIstAffectation struct {
	Affectation     Affectations    `json:"_affectation" bson:"_affectation"`
	AffectationData AffestationData `json:"_affectationData" bson:"_affectationData"`
}

type HistoriqueData struct {
}

// Add adds a new record in DB
func Add(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	var record Affectations
	err = json.Unmarshal(data, &record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusNotAcceptable, goutilities.ErrorsRender("Bad content.", err))
		return
	}

	if strings.TrimSpace(record.EquipId) == "" || strings.TrimSpace(record.UserId) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "Please fill in all the fields.")
		return
	}
	if strings.TrimSpace(record.DateFin) != "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "Le champ date de fin doit etre vide")
		return
	}

	// ─── WE ADD THE GROUP IN DB ───────────────────────────────────────────────────
	t, _ := time.Now().UTC().MarshalText()
	record.ID = uuid.Must(uuid.NewRandom()).String()
	record.CreatedAt = string(t)
	record.UpdatedAt = string(t)
	record.DateDebut = string(t)

	_, err = LocalDB.Collection(CollectionName).InsertOne(context.TODO(), record)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't create the Affectations, please retry", err))
		return
	}

	// We update the status of the equipment
	filter := bson.M{"_id": record.EquipId}
	_, err = LocalDB.Collection(equipement.CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"_EtatUti": "affecte"}})
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Affectations from db", err))
		return
	}
	list := []LIstAffectation{}
	records := []Affectations{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Affectations from db", err))
		return
	}
	for _, record := range records {
		data := LIstAffectation{}
		data.Affectation = record
		dataRecords, err := RegroupAffectationInfos(record)
		if err != nil {
			goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't regroup affectation data from db", err))
			return
		}
		data.AffectationData = dataRecords
		list = append(list, data)
	}
	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&list)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return Affectations data", err))
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
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't delete the Affectations", err))
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
	if update["_EquipId"] == "" || update["_UserId"] == "" || update["_DateDebut"] == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "You can only set the end of the affectation")
		return
	}

	// ─── WE UPDATE THE RECORD ───────────────────────────────────────────────────────
	vars := mux.Vars(r)
	sID := vars["id"]

	filter := bson.M{"_id": sID}
	_, err = LocalDB.Collection(CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the Affectations", err))
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
		bson.M{"_EquipId": bson.M{"$regex": ".*" + sSearch + ".*", "$options": "i"}},
	}}
	searchOptions := options.Find().SetSort(bson.M{"_CreatedAt": -1}).SetLimit(10000).SetProjection(bson.M{"_Password": 0})
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter, searchOptions)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Affectations from db", err))
		return
	}

	records := []Affectations{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Affectations from db", err))
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

func Desafectation(w http.ResponseWriter, r *http.Request) {
	t, _ := time.Now().UTC().MarshalText()
	update := bson.M{}
	update["_DateFin"] = string(t)

	// ─── WE UPDATE THE RECORD ───────────────────────────────────────────────────────
	vars := mux.Vars(r)
	sID := vars["id"]
	filter := bson.M{"_id": sID}
	var affec Affectations
	err := LocalDB.Collection(CollectionName).FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}).Decode(&affec)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the Affectations", err))
		return
	}

	update2 := bson.M{}
	update2["_EtatUti"] = "non-affecte"
	filter = bson.M{"_id": affec.EquipId}
	_, err = LocalDB.Collection(equipement.CollectionName).UpdateOne(context.TODO(), filter, bson.M{"$set": update2})
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't update the equipment", err))
		return
	}
	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	goutilities.APIBodyString(&w, `{"result": "ok"}`)
}

func EquipAffectedUser(w http.ResponseWriter, r *http.Request) {
	// ─── WE PULL LIST OF RECORDS ────────────────────────────────────────────────────
	filter := bson.M{"_DateFin": ""}
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Affectations from db", err))
		return
	}

	records := []Affectations{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Affectations from db", err))
		return
	}
	var results []EquipUser
	for i := 0; i < len(records); i++ {
		result := EquipUser{}
		result.EquipID = records[i].EquipId
		user := users.User{}
		filter := bson.M{"_id": records[i].UserId}
		err := LocalDB.Collection(users.CollectionName).FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Affectations from db", err))
			return
		}
		result.User = user.FirstName + " " + user.LastName
		results = append(results, result)
	}
	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(&records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return Affectations data", err))
		return
	}
	goutilities.APIBody(&w, &bufData)
}

func Historique(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sID := vars["id"]
	// We get the user wich use this equipment
	filter := bson.M{"_EquipId": sID}
	searchOptions := options.Find().SetSort(bson.M{"_CreatedAt": -1}).SetLimit(10000)
	cur, err := LocalDB.Collection(CollectionName).Find(context.TODO(), filter, searchOptions)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't pull Equipments from db", err))
		return
	}

	records := []Affectations{}
	err = cur.All(context.TODO(), &records)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't read Equipments from db", err))
		return
	}
}

func RegroupAffectationInfos(affectation Affectations) (AffestationData, error) {
	//We get the user data
	affectationData := AffestationData{}
	filter := bson.M{"_id": affectation.UserId}
	user := users.User{}
	err := LocalDB.Collection(users.CollectionName).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("user")
		return affectationData, err
	}
	filter = bson.M{"_id": affectation.EquipId}
	euip := equipement.Equipments{}
	err = LocalDB.Collection(equipement.CollectionName).FindOne(context.TODO(), filter).Decode(&euip)
	if err != nil {
		log.Println("equip")
		return affectationData, err
	}
	affectationData.User = user.FirstName + " " + user.LastName
	affectationData.EquipMarq = euip.Marque
	affectationData.EquipMod = euip.Modele
	return affectationData, nil
}

/*func RegroupHistoInfos() {
	//We get the user data
	filter := bson.M{"_id": affectation.UserId}
	user := users.User{}
	err := LocalDB.Collection(users.CollectionName).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("user")
		return
	}
}*/
