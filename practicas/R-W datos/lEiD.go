package main
import(
	"fmt"
	"bufio"
	"os"
)
func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("dame tu nombre:")
	nombre,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("hola "+ nombre)
	}

	//var edad int
	//fmt.Printf("mi edad es%d",edad)
	//fmt.Println("coloca tu edad:")
	//fmt.Scanf("%d\n",&edad)
	//fmt.Printf("mi edad es %d\n",edad)
}