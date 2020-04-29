package main
import(
	"fmt"
	"math"
)
type Abser interface{
	abs()float64
}
func main(){
	var a Abser
	f:=Myfloat(-math.Sqrt2)
	v:= vertice{3,4}
	a=f
	a=&v
	a=v 
	fmt.Println(a.abs())
}
type Myfloat float64
func (f Myfloat)abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}
type vertice struct {
	x,y float64
}
func (v*vertice)abs()float64{
	return math.Sqrt(v.x*v.x+v.y*v.y)
}