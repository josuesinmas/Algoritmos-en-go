package web
import (
	"net/http"
)
func main(){
/*http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		http.ServeFile(w,r,r.URL.Path[1:])r
	})*/
	fileServer:=http.FileServer(http.Dir("public"))
	http.Handle("/",http.StripPrefix("/",fileServer))
	http.ListenAndServe(":8000",nil)
}