package auth

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"

	"github.com/S-KYUCHAN/backend/modules/db"
	"github.com/S-KYUCHAN/backend/models"
)

func Check(username, password string) (ok bool) {

	var user models.UserModel

	// Find user with username
	// conn := db.InitDB()
	
	query := fmt.Sprintf("select * from user_tb where `username`='%s'", username)

	// result := make(map[string]interface{})
	db.DbQueryRow(query, &user.Id, &user.Name, &user.UserName, 
						&user.Password)
	
	if user.IsEmpty() {
		ok = false
	} else {
		// if err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password)); err != nil {
		if password == user.Password {
			ok = true
		} else {
			ok = false
		}
	}
	return
}

// func EncodePassword(password []byte) string {
// 	hash, err := bcrypt.GererateFromPassword(pwd, bcrypt.DefaultCost)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(hash)
// }