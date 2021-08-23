package posts

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	
	"github.com/gorilla/mux"

	"github.com/S-KYUCHAN/backend/modules/db"
)

func WelcomePage(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, "Welcome to my page")
	fmt.Println("Endpoint Hit: welcomePage")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	var query string = "select * from articles"

	rows := db.DbQuery(query)
	col, _ := rows.Columns()
	typeVal, _ := rows.ColumnTypes()

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		colVar := make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			typeName := typeVal[i].DatabaseTypeName()
			db.SetColVarType(&colVar, i, typeName)
		}
		
		result := make(map[string]interface{})
		rows.Scan(colVar...)
		for i := 0; i < len(col); i++ {
			typeName := typeVal[i].DatabaseTypeName()
			db.SetResultValue(&result, col[i], colVar[i], typeName)
		}

		results = append(results, result)
	}
	
	json.NewEncoder(w).Encode(results)

	fmt.Println("Endpoint Hit: returnAllArticles")
}

func ReturnSingleArticles(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	key := vars["id"]
	
	query := fmt.Sprintf("select * from articles where `id`=%s", key)
	
	rows := db.DbQuery(query)
	col, _ := rows.Columns()
	typeVal, _ := rows.ColumnTypes()

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		colVar := make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			typeName := typeVal[i].DatabaseTypeName()
			db.SetColVarType(&colVar, i, typeName)
		}
		
		result := make(map[string]interface{})
		rows.Scan(colVar...)
		for i := 0; i < len(col); i++ {
			typeName := typeVal[i].DatabaseTypeName()
			db.SetResultValue(&result, col[i], colVar[i], typeName)
		}

		results = append(results, result)
	}
	
	json.NewEncoder(w).Encode(results)

	fmt.Println("Endpoint Hit: returnSingleArticles")
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	result := make(map[string]interface{})
	json.Unmarshal(reqBody, &result)

	query := fmt.Sprintf("insert into articles(title,description,content)values('%s','%s','%s')",
						result["title"], result["description"], result["content"])

	_ = db.DbExec(query)
	
	fmt.Println("Endpoint Hit: createNewArticle")
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	query := fmt.Sprintf("delete from articles where `id`=%s", key)

	result := db.DbExec(query)

	fmt.Println(result)
}