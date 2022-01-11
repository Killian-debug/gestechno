package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	affectations "josephdev.io/blog_api/Affectations"

	equipement "josephdev.io/blog_api/Equipments"
	fournisseurs "josephdev.io/blog_api/Fournisseurs"
	prestataires "josephdev.io/blog_api/Prestataires"
	problemes "josephdev.io/blog_api/Problemes"
	services "josephdev.io/blog_api/Services"
	solutions "josephdev.io/blog_api/Solutions"
	techniciens "josephdev.io/blog_api/Techniciens"
	users "josephdev.io/blog_api/Users"
)

//
// ──────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: G L O B A L   V A R I A B L E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────

var ctx context.Context

// LocalDB must be exported to the other package
var LocalDB *mongo.Database
var connectedToMongo bool

//
// ───────────────────────────────────────────────────── END GLOBAL VARIABLES ─────

func main() {

	// ─── MONGO ──────────────────────────────────────────────────────────────────────
	// We connect to mongo
	err := MongoConnect()
	if err != nil {
		connectedToMongo = false
		log.Fatalf(err.Error())
	}

	connectedToMongo = true
	log.Println("Successfully connected to mongoDB")

	//Verification if we are still connected to Mongo DB
	interval := time.NewTicker(time.Minute * 5)
	go VerifConnection(interval)

	// ─── ROUTER ─────────────────────────────────────────────────────────────────────
	router := mux.NewRouter().StrictSlash(true)

	// Users
	router.HandleFunc("/users", users.Add).Methods(http.MethodPost)
	router.HandleFunc("/users", users.List).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", users.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", users.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/search/{search}", users.SeekCustom).Methods(http.MethodGet)

	// Techniciens
	router.HandleFunc("/techniciens", techniciens.Add).Methods(http.MethodPost)
	router.HandleFunc("/techniciens", techniciens.List).Methods(http.MethodGet)
	router.HandleFunc("/techniciens/{id}", techniciens.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/techniciens/{id}", techniciens.Update).Methods(http.MethodPut)
	router.HandleFunc("/techniciens/search/{search}", techniciens.SeekCustom).Methods(http.MethodGet)

	// Solutions
	router.HandleFunc("/solutions", solutions.Add).Methods(http.MethodPost)
	router.HandleFunc("/solutions", solutions.List).Methods(http.MethodGet)
	router.HandleFunc("/solutions/{id}", solutions.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/solutions/{id}", solutions.Update).Methods(http.MethodPut)
	router.HandleFunc("/solutions/search/{search}", solutions.SeekCustom).Methods(http.MethodGet)

	// Services
	router.HandleFunc("/services", services.Add).Methods(http.MethodPost)
	router.HandleFunc("/services", services.List).Methods(http.MethodGet)
	router.HandleFunc("/services/{id}", services.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/services/{id}", solutions.Update).Methods(http.MethodPut)
	router.HandleFunc("/services/search/{search}", services.SeekCustom).Methods(http.MethodGet)

	// Problemes
	router.HandleFunc("/problemes", problemes.Add).Methods(http.MethodPost)
	router.HandleFunc("/problemes", problemes.List).Methods(http.MethodGet)
	router.HandleFunc("/problemes/{id}", problemes.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/problemes/{id}", problemes.Update).Methods(http.MethodPut)
	router.HandleFunc("/problemes/search/{search}", problemes.SeekCustom).Methods(http.MethodGet)

	// Prestataires
	router.HandleFunc("/prestataires", prestataires.Add).Methods(http.MethodPost)
	router.HandleFunc("/prestataires", prestataires.List).Methods(http.MethodGet)
	router.HandleFunc("/prestataires/{id}", prestataires.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/prestataires/{id}", prestataires.Update).Methods(http.MethodPut)
	router.HandleFunc("/prestataires/search/{search}", prestataires.SeekCustom).Methods(http.MethodGet)

	// Fournisseurs
	router.HandleFunc("/fournisseurs", fournisseurs.Add).Methods(http.MethodPost)
	router.HandleFunc("/fournisseurs", fournisseurs.List).Methods(http.MethodGet)
	router.HandleFunc("/fournisseurs/{id}", fournisseurs.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/fournisseurs/{id}", fournisseurs.Update).Methods(http.MethodPut)
	router.HandleFunc("/fournisseurs/search/{search}", fournisseurs.SeekCustom).Methods(http.MethodGet)

	// Equipement
	router.HandleFunc("/equipement", equipement.Add).Methods(http.MethodPost)
	router.HandleFunc("/equipement", equipement.List).Methods(http.MethodGet)
	router.HandleFunc("/equipement/{id}", equipement.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/equipement/{id}", equipement.Update).Methods(http.MethodPut)
	router.HandleFunc("/equipement/search/{search}", equipement.SeekCustom).Methods(http.MethodGet)
	router.HandleFunc("/equipement/filter/{searchBy}", equipement.ListWithFilter).Methods(http.MethodGet)
	router.HandleFunc("/equipement/horsPark/{id}", equipement.ExitOfThePark).Methods(http.MethodPut)

	// Affectations
	router.HandleFunc("/affectations", affectations.Add).Methods(http.MethodPost)
	router.HandleFunc("/affectations", affectations.List).Methods(http.MethodGet)
	router.HandleFunc("/affectations/{id}", affectations.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/affectations/{id}", affectations.Update).Methods(http.MethodPut)
	router.HandleFunc("/affectations/search/{search}", affectations.SeekCustom).Methods(http.MethodGet)
	router.HandleFunc("/desaffectations/{id}", affectations.Desafectation).Methods(http.MethodPut)

	//Cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "x-requested-with", "user-token", "security-token", "verification-token", "inmate-token"},
		ExposedHeaders:   []string{"verification-token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
	})

	handler := c.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":8090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Printf("%s Http server running", "Blog")
	log.Fatal(srv.ListenAndServe())

}

// MongoConnect establish a connection with mongo DB
func MongoConnect() error {

	// Set Client options
	ctx = context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb+srv://killian:killian2022@cluster0.onsls.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// We make sure we are connected to the DB
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	LocalDB = client.Database("Calculator")
	users.LocalDB = LocalDB
	techniciens.LocalDB = LocalDB
	solutions.LocalDB = LocalDB
	services.LocalDB = LocalDB
	problemes.LocalDB = LocalDB
	prestataires.LocalDB = LocalDB
	fournisseurs.LocalDB = LocalDB
	equipement.LocalDB = LocalDB
	affectations.LocalDB = LocalDB
	return nil
}

//
// ────────────────────────────────────────────────────────────────── II ──────────
//   :::::: A P I   F U N C T I O N S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────
//

//
// ──────────────────────────────────────────────────────── END API FUNCTIONS ─────

//VerifConnection control the continue access of the mongo DB
func VerifConnection(interval *time.Ticker) {

	if !connectedToMongo {
		MongoConnect()
	}
}
