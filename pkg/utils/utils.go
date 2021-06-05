package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

//var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9]+([_ - .]?[a-zA-Z0-9])*$")
var usernameRegex = regexp.MustCompile(`^[a-zA-Z]+([a-zA-Z0-9](_|-|.)[a-zA-Z0-9])*[a-zA-Z0-9]+$`)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func EncryptPassword(plain string) (HashedPassword []byte, err error) {
	if len(plain) <= 7 {
		return HashedPassword, errors.New("password length should be greater than seven")
	}
	HashedPassword, err = bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return
}

func ComparePasswords(hashedPassword, password []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	return
}

func ParseStringToUint(s string) uint {
	u64, _ := strconv.ParseUint(s, 10, 32)
	return uint(u64)
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsUserNameValid(username string) bool {
	if len(username) < 4 || len(username) > 20 {
		return false
	}
	return usernameRegex.MatchString(username)
}
