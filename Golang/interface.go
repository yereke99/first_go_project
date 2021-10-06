package main
import(
  "fmt"
  "math"
)

type Shape interface{
  Area() float32
}

type Square struct{
  sideL float32
}

func (s Square) Area() float32{
  return s.sideL * s.sideL
}

type Circle struct{
  radius float32
}

func (c Circle) Area() float32{
  return float32(math.Pi) * c.radius*c.radius
}

func main(){
  square := Square{15.6}
  circle := Circle{5.6}
  printShapeArea(square)
  printShapeArea(circle)

}

func printShapeArea(s Shape){
  fmt.Println("Area figure: %0.2f sm\n", s.Area())
}
