package main

import ("fmt"
"os"
"log"
"io/ioutil")

func main() {
    file, err := os.Create("textoPrueba.txt")

    if err != nil{
        log.Fatal(err)
    }

    file.WriteString("Esto es un texto de prueba")

    file.Close()

    stream, err := ioutil.ReadFile("textoPrueba.txt")
    if err != nil{
        log.Fatal(err)
    }
    readString := string(stream)
    fmt.Println(readString)
}
