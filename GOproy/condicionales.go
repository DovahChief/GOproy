package main
import "fmt"
func main()  {

    tuEdad := 18

    if tuEdad >= 16 {
        fmt.Println("Puedes manejar")
    } else if tuEdad >= 18 {
        fmt.Println("Puedes votar")
    }else {
        fmt.Println("diviertete")
    }

}
