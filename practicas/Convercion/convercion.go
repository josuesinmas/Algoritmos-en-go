package main
import (
	"fmt"
	"strconv"
)
func main(){
	edad:="19"
	//edad_str:=strconv.Itoa(edad)
	edad_int,_:=strconv.Atoi(edad)
	fmt.Println(edad_int+10)
}