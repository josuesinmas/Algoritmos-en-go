package main
import "fmt"
type user struct{//creamos la estructura que obtendra las funciones o metodos
	edad int
	nombre,apellido string

}
func (this user) nombre_completo()string{
	return this.nombre+" "+this.apellido
}
func (this *user) set_name(n string){
	this.nombre=n
}
func main(){
	josue:=new(user)

	josue.nombre="josue"
	josue.set_name("marcos")
	fmt.Println(josue.nombre)
}