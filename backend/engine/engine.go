package engine

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"

	"github.com/S-KYUCHAN/backend/modules/posts"
	"github.com/S-KYUCHAN/backend/modules/login"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(CORS)

	myRouter.HandleFunc("/", posts.WelcomePage)
	myRouter.HandleFunc("/articles", posts.ReturnAllArticles)
	myRouter.HandleFunc("/article", posts.CreateNewArticle).Methods("POST")
	// myRouter.HandleFunc("/article/{id}", posts.UdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", posts.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", posts.ReturnSingleArticles)
	
	myRouter.HandleFunc("/login", login.Signin).Methods("POST")
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
