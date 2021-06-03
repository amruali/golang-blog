package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func EncryptPassword(password string) (HashedPassword []byte, err error) {
	HashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ComparePasswords(hashedPassword, password []byte)(err error){
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	return
}

func ParseStringToUint(s string) (uint){
	u64, _ := strconv.ParseUint(s, 10, 32)
	return uint(u64)
}
