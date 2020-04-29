package main
import "fmt"
/*func main(){
	matriz:=[]int{1,2,3,5,4}
	fmt.Println(matriz)
}*/
func main(){
	/*matriz:=[]int{1,3,5}
	if matriz==nil{
		fmt.Println("el valor es nulo")
	}else{
		fmt.Println(len(matriz))
	}*/
	//puntero al arreglo
	//longitud del arreglo al que se apunta
	//capacidad
	arreglo :=[3]int{1,1,2}
	slice:=arreglo[:3]//slicing
	fmt.Println(slice)
}