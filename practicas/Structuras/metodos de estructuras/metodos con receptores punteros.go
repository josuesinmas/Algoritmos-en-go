package main
import(
	"fmt"
)
type vertice struct{
	x,y float64
}
func (v *vertice)Scale(f float64){
	v.x=v.x*f
	v.y=v.y*f
}
func (v *vertice)abs()float64{
	return math.Sqrt(v.x*v.x+v.y*v.y)
}
func main(){
	v:=&vertice{3,4}
	v.Scale(5)
	fmt.Println(v,v.abs())
}