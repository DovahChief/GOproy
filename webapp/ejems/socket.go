package main

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

/*
	ejemplo de un web socket
*/

var addr = ":12345"

func main() {
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.
		Request) {
		http.ServeFile(w, r, "socket/websocket.html")
	})
	websocketListen()
}

//EchoLengthServer regresa largo de mensaje
func EchoLengthServer(ws *websocket.Conn) {
	var mensaje string

	for {
		websocket.Message.Receive(ws, &mensaje)
		fmt.Println("Got message", mensaje)
		length := len(mensaje)
		if err := websocket.Message.Send(ws,
			strconv.FormatInt(int64(length), 10)); err != nil {
			fmt.Println("No se puede enviar el largo del mensaje")
			break
		}
	}
}

// funcion que escucha se inicia en main
func websocketListen() {
	http.Handle("/length", websocket.Handler(EchoLengthServer))
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
