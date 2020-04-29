package main
import(
	"net/http"
	"io"
	"fmt"

)
func main(){
	http.HandleFunc("/holamundo",func(w http.ResponseWriter,r*http.Request){
		io.WriteString(w,"hooola")
	})
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8000",nil)
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("hay una nueva peticion")
	io.WriteString(w,"hola mundo")
}