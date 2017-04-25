package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//conexion a la bd
	database, er := sql.Open("mysql",
		"accusr:accusr@tcp(127.0.0.1:3306)/red_social")
	if er != nil {
		log.Fatal(er)
	}
	defer database.Close()

	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Algo salio mal!")
	}

	sql := "INSERT INTO users set user_nickname='" + NewUser.Name +
		"', user_first='" + NewUser.First + "', user_last='" +
		NewUser.Last + "', user_email='" + NewUser.Email + "'"

	q, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	//conexion a la bd
	database, er := sql.Open("mysql",
		"accusr:accusr@tcp(127.0.0.1:3306)/red_social")
	if er != nil {
		log.Fatal(er)
	}
	defer database.Close()

	urlParams := mux.Vars(r)
	id := urlParams["id"]
	ReadUser := User{}

	err := database.QueryRow("select * from users where user_id=?", id).
		Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last,
			&ReadUser.Email)

	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "no existe el usuario que buscas")
	case err != nil:
		log.Fatal(err)
		fmt.Fprintf(w, "Error")
	default:
		output, _ := json.Marshal(ReadUser)
		fmt.Fprintf(w, string(output))
	}
}

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	routes.HandleFunc("/api/user/get/{id:[0-9]+}", GetUser)
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}
