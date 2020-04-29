package main
import "fmt"

type User interface{
	Permisos()int//1-5
	Name()string
}
type Admin struct{//se crea la estructura Admin con name como string
	name string

}
func (this Admin) Permisos()int{
	return 5
}
func (this Admin) Name()string{
	return this.name
}
type Editor struct{
	name string
}
func (this Editor) Permisos()int{
	return 3
}
func (this Editor) Name()string{
	return this.name
}
func auth(user User)string{
	permisos:=user.Permisos()

	if user.Permisos()>4 {
		return user.Name()+" tiene permiso de admin"
	}else  if permisos==3{
		return user.Name()+" tiene permisos de editor"
	}
	return "i"
}
func main(){
	usuarios:=[]User{Admin{"josue"},Editor{"juan"}}
	for _,usuario:= range usuarios{
		fmt.Println(auth(usuario))
	}
}