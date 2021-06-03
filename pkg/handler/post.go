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
		if !alreayLoggedIn(r) {
			utils.RespondWithError(w, http.StatusUnauthorized, "Un Authenticated")
			return
		}

		// Get User
		userID, Ok := getUser(r)
		if !Ok {
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
		if len(data["description"]) == 0 {
			utils.RespondWithError(w, http.StatusInternalServerError, "Description shouldn't be empty")
			return
		}
		db = db.Create(&models.Post{
			Description: data["description"],
			UserID:      utils.ParseStringToUint(userID),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if db.Error != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	db, _ := db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Remember to Close Body
	defer r.Body.Close()
	data := make(map[string]string)
	if !alreayLoggedIn(r) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Un Authenticated")
		return
	}

	// Get User
	userID, Ok := getUser(r)
	if !Ok {
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

	// Get the record of Post that User Want to Update it
	var post models.Post
	res := db.Where("id", utils.ParseStringToUint(data["post_id"])).First(&post)
	if res.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, res.Error.Error())
		return
	}
	// Check that User that is trying to access post is the same owner of post
	if post.UserID != utils.ParseStringToUint(userID) {
		utils.RespondWithError(w, http.StatusInternalServerError, "UnAuthenticated, It's not even your post, bro")
		return
	}

	// Check that it's not left empty
	if len(data["description"]) == 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, "Don't leave your post empty, bro")
		return
	}

	// Update both post description & updateAt value
	post.Description = data["description"]
	post.UpdatedAt = time.Now()

	// Save to DB
	db.Save(&post)

	// Get Comments related to this Post whatever it is null or not
	var comments []models.PostComment
	db.Where("post_id", utils.ParseStringToUint(data["post_id"])).Find(&comments)
	post.PostComment = comments

	// Respond with Post Answer
	utils.RespondWithJson(w, http.StatusOK, post)

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	db, _ := db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Remember to Close Body
	defer r.Body.Close()
	data := make(map[string]string)
	if !alreayLoggedIn(r) {
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

	// Get All Comments over that Post
	var comments []models.PostComment
	res := db.Where("post_id", utils.ParseStringToUint(data["post_id"])).Find(&comments)
	if res.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, res.Error.Error())
		return
	}
	// Get the Post itself
	var post models.Post
	res = db.Where("id", utils.ParseStringToUint(data["post_id"])).First(&post)
	if res.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, res.Error.Error())
		return
	}
	// Assign Comments to Post
	post.PostComment = comments

	// Respond with Post Answer
	utils.RespondWithJson(w, http.StatusOK, post)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	db, _ := db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Remember to Close Body
	defer r.Body.Close()
	data := make(map[string]string)
	if !alreayLoggedIn(r) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Un Authenticated")
		return
	}

	// Get User
	userID, Ok := getUser(r)
	if !Ok {
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

	// Check that the post needed is already existed
	var post models.Post
	res := db.Where("id", data["post_id"]).First(&post)
	if res.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, res.Error.Error())
		return
	}

	// Check that the person trying to delete post is the same post owner
	if post.UserID != utils.ParseStringToUint(userID) {
		utils.RespondWithError(w, http.StatusInternalServerError, "UnAuthenticated, It's not even your post, bro")
		return
	}

	db.Delete(&post)

	// Respond with Json
	utils.RespondWithJson(w, http.StatusCreated, struct {
		Ok string `json:"ok"`
	}{
		Ok: "Success, Succefully Deleted Bro!",
	})

}
