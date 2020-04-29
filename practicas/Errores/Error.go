package main
import (
	"fmt"
	"time"
)
type MyError struct{
	Cuando time.Time
	Que string 
}
func (e *MyError)Error()string{
	return fmt.Sprintf("at %v, %s", e.Cuando,e.Que)
}
func run()error{
	return &MyError{
		time.Now(),
		"no trabaja",
	}
}
func main(){
	if err:= run();err!=nil{
		fmt.Println(err)
	}
}