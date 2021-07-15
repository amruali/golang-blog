package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/amrali/golang-blog/pkg/utils"
)

type Test struct {
	name     string
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Expected string
}

func TestRegister(t *testing.T) {

	tests := getTestCases()

	// Loop Over parameters
	for _, tc := range tests {
		jsonVal, _ := json.Marshal(tc)
		req, err := http.NewRequest("POST", "/register", strings.NewReader(string(jsonVal)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Register)
		handler.ServeHTTP(rr, req)

		t.Run(tc.name, func(t *testing.T) {
			if rr.Body.String() != tc.Expected {
				t.Errorf("got %v want %v", rr.Body.String(), tc.Expected)
			}
		})

	}

}

func getTestCases() []Test {
	// Using Breaking Bad Actors' names
	tests := []Test{
		{
			name:     "Success TestCase 1",
			Username: "test26",
			Email:    "test26@gmail.com",
			Password: "12345678",
			Expected: `{"ok":"Success, Welcome to our website bro!"}`,
		},
		{
			name:     "User Name is not valid",
			Username: "Gus Fring",
			Email:    "Los Pollos Hermanos",
			Password: "",
			Expected: `{"error":"username is not valid"}`,
		},
		{
			name:     "Email is not valid",
			Username: "Skyler",
			Email:    "", Password: "9235",
			Expected: `{"error":"email is not valid"}`,
		},
		{
			name:     "success case 2",
			Username: "WalterJunior",
			Email:    "walterjunior@gmail.com",
			Password: "IamHungry",
			Expected: `{"ok":"Success, Welcome to our website bro!"}`,
		},
		{
			name:     "Success TestCase 3",
			Username: "jessie",
			Email:    "Jessie@street.com",
			Password: "Iam a cooker",
			Expected: `{"ok":"Success, Welcome to our website bro!"}`,
		},
		{
			name:     "Password length < 8",
			Username: "SaulGoodman",
			Email:    "SaulGoodman@gmail.com",
			Password: "123456",
			Expected: `{"error":"password length should be greater than seven"}`,
		},
	}
	return tests
}
