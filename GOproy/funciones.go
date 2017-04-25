package main

import "fmt"
func main() {

    listNum := []float64{1,2,3,4,5}

    fmt.Println("Suma = ",sumalos(listNum))
    num1, num2 := sigDosValores(5)
    fmt.Println(num1, num2)

    fmt.Println(restalos(1,2,3,4,5,6,7))

///////FUNCION DENTRO DE FUNCION (MAIN)

    num3 := 3
    doble := func () int  {
        num3 *= 2
        return num3
    }
    doble()

    fmt.Println(doble())
    fmt.Println(doble())

    fmt.Println(Factorial(5))

}

//funcion nombre de la funcion argumrntos y retorno

func sumalos(numeros []float64) float64{
    sum := 0.0;

    for _, val := range numeros{
        sum += val
    }
    return sum
    }
//funcion que regresa 2 valores
func sigDosValores (num int) (int, int)  {
    return num+1, num+2
}
// Funcion variatica sin numero definido de valores

func restalos(args... int) int  {
    valor := 0
    for _, val := range args{
    valor -= val
        }
return valor
}
///////////////////////////////////////
//Recursiva

func Factorial(num int) int{
    if num == 0{
        return 1
    }
    return num * Factorial(num -1)
}
