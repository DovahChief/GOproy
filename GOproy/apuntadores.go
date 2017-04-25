package main
import "fmt"
///////////////
//VALORES POR REFERENCIA
func main()  {

    x := 0

    cambiaVal(&x)

    fmt.Println("x = ",x)
    fmt.Println("Direccion de memoria de x = ", &x)
}

func cambiaVal(x *int)  {
    *x += 2
}
