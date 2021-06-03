package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/amrali/golang-blog/pkg/db"
	"github.com/amrali/golang-blog/pkg/models"
	"github.com/amrali/golang-blog/pkg/utils"
)

// Post-Handlers
func AddPost(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	db, _ := db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Remember to Close Body
	defer r.Body.Close()
	data := make(map[string]string)

	if r.Method == http.MethodPost {
		if Ok := alreayLoggedIn(r); !Ok{
			utils.RespondWithError(w, http.StatusUnauthorized, "Un Authenticated")
			return
		}

		// Get User
		userID, Ok := getUser(r)
		if !Ok{
			utils.RespondWithError(w, http.StatusUnauthorized, "Un Authenticated")
			return
		}
		// Read Request Body
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Decode Body to data
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		// Create Record
		db.Create(&models.Post{
			Description: data["desription"],
			UserID: utils.ParseStringToUint(userID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		// Respond with Json
		utils.RespondWithJson(w, http.StatusCreated, struct {
			Ok string `json:"ok"`
		}{
			Ok: "Success, Submitted Succefully Bro!",
		})
	}
	
}


func UpdatePost(w http.ResponseWriter, r *http.Request) {

}
func GetPost(w http.ResponseWriter, r *http.Request) {

}
func DeletePost(w http.ResponseWriter, r *http.Request) {

}