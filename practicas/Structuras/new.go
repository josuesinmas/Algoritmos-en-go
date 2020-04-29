package main
import "fmt"
type Vertice struct {
	X,Y int
}
func main(){
	v:= new(Vertice)
	fmt.Println(v)
	v.X,v.Y=11,9
	fmt.Println(v)
}