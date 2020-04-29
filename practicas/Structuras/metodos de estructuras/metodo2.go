package main
import (
	"fmt"
	"math"
)
type Vertice struct{
	x,y float64
}
func (v*Vertice)abs()float64{
	return math.Sqrt(v.x*v.x+v.y*v.y)
}
func main(){
	v:=&Vertice{3,4}
	fmt.Println(v.abs())
}