package main
import "fmt"
func sum(a[]int, c chan int){
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // envia sum a c
}
func main(){
	a := []int{7,2,8,-9,4,0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // recive de c. los datos fluyen en la direccion de las flechas y tienen que crearce los canales antes de usarlos
	fmt.Println(x, y, x + y)
}

