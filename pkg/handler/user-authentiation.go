package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/amrali/golang-blog/pkg/db"
	"github.com/amrali/golang-blog/pkg/models"
	jwt "github.com/amrali/golang-blog/pkg/token"
	"github.com/amrali/golang-blog/pkg/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	db, _ := db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Remember to Close Body
	defer r.Body.Close()
	data := make(map[string]string)
	if r.Method == http.MethodPost {
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

		// Create User-Object
		var user = models.User{}

		// Check that UserName is not Taken
		res := db.Where("user_name = ?", data["username"]).Find(&user)
		if res.RowsAffected == 1 && data["username"] == user.UserName {
			utils.RespondWithError(w, http.StatusInternalServerError, "User Name is already taken")
			return
		}

		// Check that Email is not taken
		res = db.Where("email = ?", data["email"]).Find(&user)
		if res.RowsAffected == 1 && data["email"] == user.Email {
			utils.RespondWithError(w, http.StatusInternalServerError, "Email is already taken")
			return
		}

		// Create a models.User{} object
		user.Email = data["email"]
		user.UserName = data["username"]

		// SetUp Password
		user.Password, err = utils.EncryptPassword(data["password"])
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Incorrect Password")
			return
		}
		// Insert Record to DB
		db.Create(&models.User{
			UserName:  user.UserName,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		// Return Ok if and only if all the above is passed
		utils.RespondWithJson(w, http.StatusCreated, struct {
			Ok string `json:"ok"`
		}{
			Ok: "Success, Welcome to our website bro!",
		})
	}

}


func Login(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	db, _ := db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Remember to Close Body
	defer r.Body.Close()
	data := make(map[string]string)
	if r.Method == http.MethodPost {
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

		// Create User-Object
		var user = models.User{}

		// Assign map valus to variables
		username := data["username"]
		//password := data["password"]

		// Check UserName is alreay exist
		res := db.Where("user_name = ?", username).Find(&user)
		if res.Error != nil || res.RowsAffected == 0 {
			utils.RespondWithError(w, http.StatusInternalServerError, "username or password are wrong")
			return
		}

		// Compare Password
		if err = utils.ComparePasswords(user.Password, []byte(data["password"])); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "username or password are wrong")
			return
		}

		// Create JWT
		token, err := jwt.CreateToken(strconv.Itoa(int(user.ID)))
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't login")
			return
		}

		// Set a Cookie
		cookie := http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Minute * 20),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		// Respond With Ok Status if all above is passed
		utils.RespondWithJson(w, http.StatusOK, struct {
			Ok string `json:"ok"`
		}{
			Ok: "Success, Welcome back our friend !",
		})
	}

}


func Logout(w http.ResponseWriter, r *http.Request) {
	// Ckeck that He is not already LoggedIn
	Ok := alreayLoggedIn(r)
	if !Ok{
		utils.RespondWithError(w, 205, "You 're not already logged in")
		return
	}
	// Delete both Cookie and Token
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Minute * 60),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	// Return Ok status
	utils.RespondWithJson(w, http.StatusOK, struct {
		Ok string `json:"ok"`
	}{
		Ok: "Success, Good bye our friend !",
	})
}


func alreayLoggedIn(r *http.Request)(bool){
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return false
	}
	Ok := jwt.VerifyToken(cookie.Value)
	return Ok
}


