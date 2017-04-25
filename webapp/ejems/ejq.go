package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	s, err := ioutil.ReadFile("pagina/index.html")
	if err != nil {
		fmt.Println("algo salio mal")
	} else {
		fmt.Println(string(s))
	}

}
