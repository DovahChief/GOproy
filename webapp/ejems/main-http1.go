package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Api struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		mesg := Api{"hola mundo soy api"}

		output, err := json.Marshal(mesg)

		if err != nil {
			fmt.Println("se presento un error")
		}

		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}
