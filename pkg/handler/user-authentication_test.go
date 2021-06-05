package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/amrali/golang-blog/pkg/utils"
)

func TestRegister(t *testing.T) {

	// Using Breaking Bad Actors' names
	parameters := []struct {
		Id       string
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Expected string
	}{
		{Id: "1", Username: "test26", Email: "test26@gmail.com", Password: "12345678", Expected: `{"ok":"Success, Welcome to our website bro!"}`},
		{Id: "2", Username: "Gus Fring", Email: "Los Pollos Hermanos", Password: "", Expected: `{"error":"username is not valid"}`},
		{Id: "3", Username: "Skyler", Email: "", Password: "9235", Expected: `{"error":"email is not valid"}`},
		{Id: "4", Username: "WalterJunior", Email: "walterjunior@gmail.com", Password: "IamHungry", Expected: `{"ok":"Success, Welcome to our website bro!"}`},
		{Id: "5", Username: "jessie", Email: "Jessie@street.com", Password: "Iam a cooker", Expected: `{"ok":"Success, Welcome to our website bro!"}`},
		{Id: "6", Username: "SaulGoodman", Email: "SaulGoodman@gmail.com", Password: "123456", Expected: `{"error":"password length should be greater than seven"}`},
	}

	// Loop Over parameters
	for _, val := range parameters {
		jsonVal, _ := json.Marshal(val)
		req, err := http.NewRequest("POST", "/register", strings.NewReader(string(jsonVal)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Register)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		/*
			if status := rr.Code; status != http.StatusCreated {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		*/
		t.Run(val.Expected, func(t *testing.T) {
			if rr.Body.String() != val.Expected {
				t.Errorf("test id %v, handler returned unexpected body: got %v want %v",
					val.Id, rr.Body.String(), val.Expected)
			}
		})

		/*
			if rr.Body.String() != val.Expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), val.Expected)
			}
		*/
	}

}
