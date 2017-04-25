package main

import ("fmt"
"strings"
"sort")

func main() {
    cadMuestra := "Hola Mundo"
    fmt.Println(strings.Contains(cadMuestra, "la"))
    fmt.Println(strings.Index(cadMuestra, "u"))
    fmt.Println(strings.Count(cadMuestra, "o"))

    csvString := "a,s,d,c,f,e,r,t,g"

    fmt.Println(strings.Split(csvString, ","))

    lista := []string{"v","b","u","y"}

    sort.Strings(lista)

    fmt.Println(lista)

}
/*
"os"
"log"
"io/ioutil"
"strconv" */
