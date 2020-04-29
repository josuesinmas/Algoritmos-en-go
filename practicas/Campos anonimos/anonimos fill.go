package main
import "fmt"
type human struct{
	name string
}

func (this tutor )hablar()string{
	return "hola que tal como estamos"
}
type dumy struct{
	name string
}

type tutor struct {
	human
}

func main (){
	tutor:=tutor{human{"josue"}}

	fmt.Println(tutor.hablar())
}