package users

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// LocalDB must be exported to package main to take the data base info
var LocalDB *mongo.Database

//CollectionName is the collection name __Users in DB
const CollectionName string = "__Users"

type User struct {
	// Default infos
	ID        string `json:"_id" bson:"_id"`
	CreatedAt string `json:"_CreatedAt" bson:"_CreatedAt"`
	UpdatedAt string `json:"_UpdatedAt" bson:"_UpdatedAt"`

	// Other Infos
	FirstName string `json:"_FirstName" bson:"_FirstName"`
	LastName  string `json:"_LastName" bson:"_LastName"`
	Number    string `json:"_Number" bson:"_Number"`
}

/*/ Login login a user an give them auth token and user info
func Login(w http.ResponseWriter, r *http.Request) {

	// ─── WE GET THE DATA SENT ───────────────────────────────────────────────────────
	reqData := User{}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		goutilities.APIResult(&w, http.StatusBadRequest, goutilities.ErrorsRender("Can't retrieve user infos", err))
		return
	}

	if strings.TrimSpace(reqData.Email) == "" || strings.TrimSpace(reqData.Password) == "" {
		goutilities.APIResult(&w, http.StatusBadRequest, "Please fill in all the fields")
		return
	}

	if checkRequestBanned(reqData.Email) {
		goutilities.APIResult(&w, http.StatusBadRequest, "You can't login now please try again later")
		return
	}

	// We first check user existence
	theUser, userAlreadyExist := CheckUserExistenceEmailAndReturnUser(reqData.Email)
	if !userAlreadyExist {
		goutilities.APIResult(&w, http.StatusNotFound, "Invalid credentials")
		return
	}

	// We check if the user is active
	if !theUser.Active {

		goutilities.APIResult(&w, http.StatusUnauthorized, "Your account is not active")
		return
	}

	// Check password
	if goutilities.MyHash([]byte(reqData.Password)) != theUser.Password {
		goutilities.APIResult(&w, http.StatusNotFound, "Invalid credentials")
		return
	}

	// user already exist we unban him
	allow(reqData.Email)

	// anonymity confidential infos
	theUser.Password = ""
	theUser.Signature = ""

	d, _ := time.ParseDuration(fmt.Sprintf("%vm", gp.PConfig.JWTTokenDuration))
	sToken, err := gp.CreateToken(reqData.Email, d)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Error while generating token", err))
		return
	}

	repData := struct {
		Token string `json:"Token"`
		User  Users  `json:"User"`
	}{
		Token: sToken, User: *theUser,
	}

	// ─── OK ─────────────────────────────────────────────────────────────────────────
	w.Header().Set("Content-Type", "application/json")
	bufData, err := json.Marshal(repData)
	if err != nil {
		goutilities.APIResult(&w, http.StatusInternalServerError, goutilities.ErrorsRender("Can't return user data", err))
		return
	}

	goutilities.APIBody(&w, &bufData)

}*/
