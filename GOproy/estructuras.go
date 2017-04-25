package main
import "fmt"
import "math"

func main()  {

    rect1 := rectangulo{0,50,10,10}
    fmt.Println("El rectangulo 1 es ", rect1.ancho, "de ancho")
    fmt.Println("Area es ", rect1.area())

    r := rect{20,50}
    c := circ{4}

fmt.Println("Area del circulo = ", getArea(c))
fmt.Println("Area del rectángulo = ", getArea(r))

}
type rectangulo struct {
    xizq float64
    ysup float64
    altura float64
    ancho float64
}
func (r *rectangulo) area() float64  {
    return r.altura * r.ancho
}

type Forma interface {
    area() float64
}
type rect struct {
    altura float64
    ancho float64
}
type circ struct {
    radio float64
}
func (r rect) area() float64  {
    return r.altura * r. ancho
}
func (c circ) area() float64  {
    return math.Pi * math.Pow(c.radio,2)
}
func getArea(f Forma) float64  {
    return f.area()
}
