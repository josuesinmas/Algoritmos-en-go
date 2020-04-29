package main
import "fmt"
func main(){
	/*el segundo parametro es el buffer y son blokeadas 
	si el buffer esta lleno y se bloquean  cuando esta vacio*/
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}