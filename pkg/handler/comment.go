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

// Comment-handlers
func AddComment(w http.ResponseWriter, r *http.Request) {
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

		if len(data["description"]) == 0 {
			utils.RespondWithError(w, http.StatusInternalServerError, "Description shouldn't be empty")
			return
		}
		res := db.Create(&models.PostComment{
			PostID: utils.ParseStringToUint(data["post_id"]),
			Description: data["description"],
			UserID: utils.ParseStringToUint(userID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if res.Error != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, res.Error.Error())
			return
		}

		// Respond with Json
		utils.RespondWithJson(w, http.StatusCreated, struct {
			Ok string `json:"ok"`
		}{
			Ok: "Success, Submitted Succefully Bro!",
		})
	}
}
func UpdateComment(w http.ResponseWriter, r *http.Request) {}
func GetComment(w http.ResponseWriter, r *http.Request) {}
func DeleteComment(w http.ResponseWriter, r *http.Request) {}