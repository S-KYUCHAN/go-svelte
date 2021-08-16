package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"github.com/S-KYUCHAN/backend/db"
)

type Article struct {
	Id			string `json:"id"`
	Title 		string `json:"title"`
	Desc 		string `json:"description"`
	Content 	string `json:"content"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	var Articles []Article
	var query string = "select * from articles"

	rows := db.DbQuery(query)
	for rows.Next() {
		var article Article
		rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		Articles = append(Articles, article)
	}
	
	json.NewEncoder(w).Encode(Articles)
	
	fmt.Println("Endpoint Hit: returnAllArticles")
}

func returnSingleArticles(w http.ResponseWriter, r *http.Request) {
	var Articles []Article
	vars := mux.Vars(r)
	key := vars["id"]
	
	query := fmt.Sprintf("select * from articles where `id`=%s", key)
	
	rows := db.DbQuery(query)
	for rows.Next() {
		var article Article
		rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		Articles = append(Articles, article)
	}
	
	json.NewEncoder(w).Encode(Articles)
	
	fmt.Println("Endpoint Hit: returnSingleArticles")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	// Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	for index, article := range Articles {
// 		if article.Id == id {
// 			fmt.Fprintf(w, "This article!!")
// 		}
// 	}
// }

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]

	// for index, article := range Articles {
	// 	if article.Id == id {
	// 		Articles = append(Articles[:index], Articles[index+1:]...)
	// 	}
	// }
}

func welcomePage(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, "Welcome to my page")
	fmt.Println("Endpoint Hit: welcomePage")
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Methods:", "*")
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(CORS)

	myRouter.HandleFunc("/", welcomePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	// myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticles)
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	handleRequests()
}