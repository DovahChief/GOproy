package main

import "fmt"

func main() {

    const pi float64 = 3.14159256

    var myName string = "José Antonio"

    fmt.Println(len(myName))

    fmt.Println(myName+ " es la verga \nsi wey es mejor que tu ")

    var tieneMasDe40 bool = true
    if(tieneMasDe40){

        fmt.Printf("%f \n", pi)
        fmt.Printf("%T \n", pi)

    }
    fmt.Printf("Normal = %d \n", 100)
    fmt.Printf(" en binario = %b \n", 100)
    fmt.Printf(" en caractér = %c \n", 100)
    fmt.Printf(" en hex = %x \n", 100)
    fmt.Printf(" en not cientifica = %e \n", 100)




}
