package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("logs/ej.txt")

	if err != nil {
		fmt.Println("error fatal")
	}

	a := 0

	for a <= 5 {
		fmt.Println(a)
		log.SetOutput(file)
		log.Printf("vale = %d", a)
		//file.Write(log.Printf("valor = %d", a))
		a++
	}

}
