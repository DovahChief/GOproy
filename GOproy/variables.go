package main

import "fmt"

func main() {

    var age int = 21

    var numFav float64 = 117.343

    const pi float64 = 3.14159256

    var (
        varA = 2
        varB = 3
    )

    fmt.Println(varA + varB)

    numAzar := 1
    //Por default la coma pone espacios entre valores
    fmt.Println(age,numFav,numAzar)
    fmt.Println(age," ",numFav, " ",numAzar)

    fmt.Println("6 + 4 =", 6 + 4)
    fmt.Println("6 - 4 =", 6 - 4)
    fmt.Println("6 * 4 =", 6 * 4)
    fmt.Println("6 / 4 =", 6 / 4)
    fmt.Println("6 % 4 =", 6 % 4)


}
