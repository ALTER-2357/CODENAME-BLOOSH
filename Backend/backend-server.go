package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"

	"github.com/bdwilliams/go-jsonify/jsonify"
	_ "github.com/go-sql-driver/mysql"
)

type userdata struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userdatabase []userdata = []userdata{}
var db *sql.DB
var err error
var ided int

var jsonFile string

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "NEXT"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/upload", uploadFile).Methods("POST")

	//	router.HandleFunc("/posts/", insert).Methods("POST")

	//	router.HandleFunc("/GET/", getPosts).Methods("GET")

	//	router.HandleFunc("/DELETE/", deletePost).Methods("POST")

	http.ListenAndServe(":3031", router)

	log.Println("backends up")

}

func getPosts(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	q := "SELECT * FROM KY WHERE ID = ?"
	ID := r.FormValue("ID")

	rows, err := db.Query(q, ID)
	if err != nil {
		log.Fatal(err)
	}

	jsonFile := (jsonify.Jsonify(rows))
	fmt.Println(jsonFile)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonFile)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	ID := r.FormValue("ID")

	insForm, err := db.Prepare("DELETE FROM KY WHERE ID = ?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(ID)

	defer db.Close()
	json.NewEncoder(w).Encode("done")
	///could make this form post
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		ID := r.FormValue("ID")
		name := r.FormValue("name")
		username := r.FormValue("username")
		insForm, err := db.Prepare("INSERT INTO KY (ID, name, username ) VALUES(?,?,? )")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(ID, name, username)
	}
	defer db.Close()
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("cvs", "*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "bloosh\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprint(w, "───────────────────█████\n───────────────────██████\n───────────────────███████\n──────────────────████████\n──────────────────████████\n─────────────────█████████\n────────────────█████████\n───────────────█████████\n──────────────█████████\n──────────────██████████████████\n────────────█████████████████████\n───────────███████████████████████\n████████─██████████████████████████\n████████─███████████████████████████\n████████─████████████████████████████\n████████─████████████████████████████\n████████─████████████████████████████\n████████─████████████████████████████\n████████─███████████████████████████\n████████─██████████████████████████\n████████─█████████████████████████\n████████─████████████████████████\n████████─███████████████████████\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "we got it!!!!!!!!")

}
