package main
import "fmt"
func main(){
	//if 
	//else if
	//else
	/*
	== igual a
	!= diferente de
	> mayor que
	< menor que
	>= mayor o igual
	<= menor o igual
	&& and
	|| OR 
	*/
	x:=10
	y:=11
	if (x>y){
		fmt.Printf("%d es mayor que %d\n ",x,y)
	}else if(x<y){
		fmt.Printf("%d es menor que %d ",x,y)
	}else{
		fmt.Println("son inguales")
	}
}