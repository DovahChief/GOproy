package api

/*
	Api de golang con estructura para conectar a bd con diferentes endpoints
	----
	Routes.HandleFunc("/api.{format:json|xml|txt}/user",UsersRetrieve).Methods("GET")
	---- ejem

*/

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Database es la conexion a bd
var Database *sql.DB

//Format indica el tipo de request json|xml|txt
var Format string

//Users estructura de varios usuarios se llena al sacar filas de bd
type Users struct {
	Users []User `json:"users"`
}

//User captura usuarios a ingresar en la bd
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	First string `json:"first"`
	Last  string `json:"last"`
}

// StartServer inicia el servicio puerto 8080
func StartServer() {
	db, err := sql.Open("mysql", "root@/social_network")
	if err != nil {
	}
	Database = db
	routes := mux.NewRouter()
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}

//GetFormat saca formato de la URL json|xml|txt
func GetFormat(r *http.Request) {
	Format = r.URL.Query()["format"][0]
}

//SetFormat propporciona la salida de la api en el formato especificado
func SetFormat(data interface{}) []byte {
	var apiOutput []byte
	if Format == "json" {
		output, _ := json.Marshal(data)
		apiOutput = output
	} else if Format == "xml" {
		output, _ := xml.Marshal(data)
		apiOutput = output
	}
	return apiOutput
}

//UsersRetrieve saca a usuarios de la bd
func UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting retrieval")
	GetFormat(r)
	start := 0
	limit := 10
	next := start + limit

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", "<http://localhost:8080/api/users?start="+
		string(next)+"; rel=\"next\"")

	rows, _ := Database.Query("SELECT * FROM users LIMIT 10")
	Response := Users{}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)
		Response.Users = append(Response.Users, user)
	}
	output := SetFormat(Response)
	fmt.Fprintln(w, string(output))
}
