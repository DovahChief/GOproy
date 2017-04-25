//usando paquete routes
//mas rapido que gorilla y especializado

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drone/routes"
)

type API struct {
	Message string `json:"message"`
}

func Hola(w http.ResponseWriter, r *http.Request) {

	urlParams := r.URL.Query()
	name := urlParams.Get(":name")
	HelloMessage := "Hola , " + name
	message := API{HelloMessage}

	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Algo salio mal!")
	}
	fmt.Fprintf(w, string(output))

}

func main() {
	mux := routes.New()
	mux.Get("/api/:name([a-z]+)", Hola)
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
