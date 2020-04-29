package main
import(
	"fmt"
	"bufio"
	"os"
)
func main(){
	ejecucion:=readFile()
	fmt.Println(ejecucion)
}
	func readFile()bool{
		archivo,error:=os.Open("./hola.txt")
		defer func(){
			archivo.Close()
			fmt.Println("Defer")
		}()
		if error !=nil{
			fmt.Println("hubo un error")
			panic(error)
		}
		scanner:=bufio.NewScanner(archivo)
		var i int
		for scanner.Scan(){
			i++
			linea:= scanner.Text()
			fmt.Println(i)
			fmt.Println(linea)
		}
	if true{
	return false
	}
	fmt.Println("hola esto nunca se ejeculta")
	archivo.Close()
	return true 
}