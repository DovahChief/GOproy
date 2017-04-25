package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func maneja(w http.ResponseWriter, r *http.Request) {
	s, err := ioutil.ReadFile("pagina/index2.html")
	if err == nil {
		fmt.Fprint(w, string(s))
	} else {
		fmt.Println("hubo un error")
	}
}

func maneja2(w http.ResponseWriter, r *http.Request) {
	s, err := ioutil.ReadFile("pagina/index.html")
	if err == nil {
		fmt.Fprint(w, string(s))
	} else {
		fmt.Println("hubo un error")
	}
}

func main() {
	fmt.Println("iniciando")
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/a", maneja)
	gorillaRoute.HandleFunc("/b", maneja2)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
