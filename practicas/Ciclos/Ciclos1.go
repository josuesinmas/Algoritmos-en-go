package main 
import (
	"fmt"
	"math"
)
func main(){
	var x,a,v float64//se declaran variables flotantes 
	var year1,year2 bool//se declaran boleanos
	var Dinero int//se declaran enteros
	x=5.6
	a=x*2.4
	v=(a*x)/3.2
	f:=math.Sqrt(v)

	for f< 1000{
		f+=f
		fmt.Println(f)
	}
	if (year1==true||year2==true){
		for i:=0;i<10;i++{
			for {//esto es un bucle infinito
				fmt.Println(i)
				Dinero+=4
				if i==2{
					fmt.Println("paso por 2")
					//continue hace que se inicie 
				}
				if (i>10){
					break
				}
			}
		}
	}else if (year1==true ||year2==false){
		fmt.Println("aqui no a pasado na ")
	}else{
		fmt.Println("nada importante")
	}
}