package main
import "fmt"

func main()  {
    fmt.Println(divSegura(3,0))
    fmt.Println(divSegura(4,2))
    demoPanic()

}

func divSegura(num1, num2 int) int  {
//Permite que siga la ejecucion si hay error
    defer func ()  {
        fmt.Println(recover())
    }()

    soluc := num1/num2
    return soluc

}

func demoPanic()  {
    defer func(){
            fmt.Println(recover())
    }()
    panic("PANIC")
}
