package main

import "fmt"
func main(){
	/*1 es una direccion de memoria
	2 en lugar del valor tenemos la direccion el la que está el valor
	*/
	var x,y *int//tipo de dato
	entero := 5
	x = &entero
	y = &entero
	*x = 6//nombre de un puntero
	fmt.Println(x)
	fmt.Println(y)
}