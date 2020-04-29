package main
import(
	"fmt"
	"time"
	"strings"
)
func main(){
	go nombre_lento("josue")
	fmt.Println("ES el mejor")
	var wait string
	fmt.Scanln(&wait)
}
func nombre_lento(name string){
	letras:=strings.Split(name,"")
	for _,letras := range(letras){
		time.Sleep(1000*time.Millisecond)
		fmt.Println(letras)
	}
}