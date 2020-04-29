package main
import "fmt"
type Vertice struct{
	x,y int
}
var(
	p=Vertice{1,2}
	q=&Vertice{1,2}
	r=Vertice{x:1}
	s=Vertice{}
)
func main() {
	fmt.Println(p,q,r,s)
}