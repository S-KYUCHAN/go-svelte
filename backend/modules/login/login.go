package login

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/S-KYUCHAN/backend/modules/auth"
	// "github.com/S-KYUCHAN/backend/modules/db"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	result := make(map[string]interface{})
	json.Unmarshal(reqBody, &result)
	
	fmt.Println(result["username"], result["password"])
	
	username := fmt.Sprintf("%s", result["username"])
	password := fmt.Sprintf("%s", result["password"])

	if auth.Check(username, password) {
		fmt.Println("Auth Complete")
	} else {
		fmt.Println("Wrong Password")
	}

	fmt.Println("Endpoint Hit: Signin")
}