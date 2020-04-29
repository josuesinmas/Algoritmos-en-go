package main
import "fmt"
type Vertice struct{
	Lat,Long float64
}
var m map[string]Vertice
func main(){
	m=make(map[string]Vertice)
	m["bell labs"]=Vertice{40.54343,-74.12373,
	}
	fmt.Println(m["bell labs"])
}