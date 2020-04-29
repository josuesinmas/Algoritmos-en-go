package main
import "fmt"
type Vertice struct{
	Lat,Long float64
}
/*
m[key]=elem insertar o actualizarun elemento en un mapa
elem=m[key]obtener un elemento
delet(m,key)borrar un elemento
elm,ok=m[key] determinar si la llave esta presente con una asignacion multiple
si la llave key esta en m ok sera true

*/
var m =map[string]Vertice{
	"bell labs":Vertice{43.22342,-74.23424,},
	"google":Vertice{37.34534,-123.34334,},
}
func main(){
	fmt.Println(m)
}