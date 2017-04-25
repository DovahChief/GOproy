//usando gorilla webkit tools

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func Hola(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]

	mesgHola := "Hola ," + name

	message := API{mesgHola}

	salida, err := json.Marshal(message)

	if err != nil {
		fmt.Println("hubo un error")
	}
	fmt.Fprintf(w, string(salida))

}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hola)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
