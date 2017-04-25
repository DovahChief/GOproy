package main
import "fmt"

func main()  {
    var numFav[5] float64

    numFav[0]= 117
    numFav[1]= 117343
    numFav[2]= 343
    numFav[3]= 11
    numFav[4]= 21

    fmt.Println(numFav[3])

    for i, value := range numFav{
        fmt.Println(value, i)
    }
    for _, value := range numFav{
        fmt.Println(value)
    }

    // Slice

    numSlice := []int {23,2,2,112,4,43,55,76}

    slice2 := numSlice[3:6]

    fmt.Println("slice2 [0] = ", slice2[0])



}
