package main
import "fmt"
	/*type define un nuevo tipo
	struct declara la estuctura*/
type User struct{
		edad int
		nombre string
		apellido string
}
func main(){
	josue:=new(User)
	josue.nombre="josue"
	fmt.Println(josue.nombre)
	//fmt.Println(User{nombre:"josue",apellido:"alvarez",edad:19})
}