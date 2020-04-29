package main
import(
	"io/ioutil"
	"fmt"
)
func main(){
	//ioutil.ReadFile("./hola.txt")
	file_data,err:=ioutil.ReadFile("./hola.txt")
	if err!= nil{
		fmt.Println("hubo un error")
	}
	fmt.Println(string(file_data))
}