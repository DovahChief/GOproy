package main

import ("fmt"
"net/http")

func main() {
    http.HandleFunc("/",handler)

    http.HandleFunc("/tierra",handler2)

    http.ListenAndServe(":8080",nil)
}
func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola mundo\n")
}
func handler2(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola tierra\n")
}
