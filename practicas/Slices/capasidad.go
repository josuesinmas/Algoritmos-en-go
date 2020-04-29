package main
import "fmt"
func main(){
	slice :=[]int{3,5,4,3,644,4}
	copia:=make([]int,len(slice),cap(slice)*2)
	//slice=append(slice,2,5)
	//copy(destino,fuente)
	copy(copia,slice)
	fmt.Println(/*cap,len*/slice)
	fmt.Println(copia)

}