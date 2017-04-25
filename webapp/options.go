package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var database *sql.DB

type users struct {
	Users []user `json:"users"`
}

type user struct {
	ID       int    `json:"id"`
	NomUsr   string `json:"username"`
	Email    string `json:"email"`
	Nombre   string `json:"first"`
	Apellido string `json:"last"`
}
type CreateResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"code"`
}
type ErrMsg struct {
	ErrCode    int
	StatusCode int
	Msg        string
}

func main() {

	db, err := sql.Open("mysql",
		"accusr:accusr@tcp(127.0.0.1:3306)/red_social")

	if err != nil {
		fmt.Println("Error al conectar con base de datos")
	}

	database = db

	ruta := mux.NewRouter()
	ruta.HandleFunc("/api/crea", creaUsr).Methods("POST")
	ruta.HandleFunc("/api/users", obtenUsr).Methods("GET")
	ruta.HandleFunc("/api/borra/{user:[0-9]+}", borraUsr)
	ruta.HandleFunc("/api/cambia/{user:[0-9]+}", borraUsr).Methods("POST")
	ruta.HandleFunc("/api", index)
	ruta.HandleFunc("/api/form", form)

	http.Handle("/", ruta)
	http.ListenAndServe(":8080", nil)

}

//funcion para index de creacion
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pagina/index.html")
}

//funcion de pagina
func form(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pagina/ccuenta.html")
}

//funcion maneja error de bd
func dbErrorParse(err string) (string, int64) {
	Parts := strings.Split(err, ":")
	errorMessage := Parts[1]
	Code := strings.Split(Parts[0], "Error ")
	errorCode, _ := strconv.ParseInt(Code[1], 10, 32)
	return errorMessage, errorCode
}

//error de mensaje
func ErrorMessages(err int64) ErrMsg {
	em := ErrMsg{}
	errorMessage := ""
	statusCode := 200
	errorCode := 0

	switch err {
	case 1062:
		errorMessage = "Duplicate entry"
		errorCode = 10
		statusCode = 409
	}
	em.ErrCode = errorCode
	em.StatusCode = statusCode
	em.Msg = errorMessage
	return em
}

//create
func creaUsr(w http.ResponseWriter, r *http.Request) {
	//funcion create saca valores de formulario en url e inserta en bd
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//
	fmt.Println(r)

	NuevoUsuario := user{}
	NuevoUsuario.Nombre = r.FormValue("first")
	NuevoUsuario.Email = r.FormValue("email")
	NuevoUsuario.Apellido = r.FormValue("last")
	NuevoUsuario.NomUsr = r.FormValue("user")

	sal, err := json.Marshal(NuevoUsuario)
	fmt.Println(string(sal))

	if err != nil {
		fmt.Println("Algo salio mal")
	} //ve que el usuario cumpla con lo establecido en la API

	sqlq := "INSERT INTO users set user_nickname= '" + NuevoUsuario.NomUsr +
		"', user_first= '" + NuevoUsuario.Nombre + "', user_last = '" +
		NuevoUsuario.Apellido + "', user_email ='" + NuevoUsuario.Email + "'"

	q, err := database.Exec(sqlq)

	if err != nil {
		Response := CreateResponse{}
		errorMessage, errorCode := dbErrorParse(err.Error())

		//error, httpCode, msg := ErrorMessages(errorCode)
		mserr := ErrorMessages(errorCode)

		Response.Error = mserr.Msg
		Response.ErrorCode = mserr.ErrCode

		//http.Error(w, "Conflict", mserr.StatusCode)
		http.ServeFile(w, r, "pagina/error.html")
		fmt.Println(mserr.StatusCode)

		fmt.Println("Algo salio mal al momento de insertar")
		fmt.Println(errorMessage)
		fmt.Println(err)
	} else {
		http.ServeFile(w, r, "pagina/creado.html")
	}
	fmt.Println(q)

}
func obtenUsr(w http.ResponseWriter, r *http.Request) {
	//saca y muestra un maximo de 10 usuarios de la bd
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println(r)
	filas, _ := database.Query("select * from users LIMIT 10")
	Respuesta := users{}

	for filas.Next() {
		usuario := user{}

		filas.Scan(&usuario.ID, &usuario.NomUsr, &usuario.Nombre,
			&usuario.Apellido, &usuario.Email)

		Respuesta.Users = append(Respuesta.Users, usuario)
	}
	sal, _ := json.Marshal(Respuesta)
	fmt.Fprintf(w, string(sal))
}

func borraUsr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	urlParams := mux.Vars(r)
	id := urlParams["user"]
	sqlq := "DELETE FROM users WHERE user_id = " + id + ";"
	fmt.Println(sqlq)
	q, err := database.Exec(sqlq)

	if err != nil {
		fmt.Println("ALgo salio mal")
		fmt.Println(err)
	}

	fmt.Println(q)

	fmt.Println("Usuario " + id + " eliminado")
}

func updateUsr(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	NuevoUsuario := user{}
	NuevoUsuario.Nombre = r.FormValue("first")
	NuevoUsuario.Email = r.FormValue("email")
	NuevoUsuario.Apellido = r.FormValue("last")
	NuevoUsuario.NomUsr = r.FormValue("user")

	urlParams := mux.Vars(r)
	id := urlParams["user"]

	sal, err := json.Marshal(NuevoUsuario)
	fmt.Println(string(sal))

	if err != nil {
		fmt.Println("Algo salio mal en update")
	}

	sqlq := "UPDATE users SET user_nickname= '" + NuevoUsuario.NomUsr +
		"', user_first= '" + NuevoUsuario.Nombre + "', user_last = '" +
		NuevoUsuario.Apellido + "', user_email ='" + NuevoUsuario.Email + "'" +
		"WHERE user_id = " + id + " ;"

	q, err := database.Exec(sqlq)
	if err != nil {
		fmt.Println("Algo salio mal al momento de actualizar")
		fmt.Println(err)
	}
	fmt.Println(q)
}
