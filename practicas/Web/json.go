package Web
import(
	"net/http"
	"encoding/json"
)
type Curso struct{
	Title string `json:"title`
	NumeroVideo int
}
type Cursos[]Cursos

func web(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
	
		cursos:=Cursos{
		Curso{"curso de go",30},
		Curso{"curso de Ruby",60},
		Curso{"curso de Pyton",40},
		Curso{"curso de html",20},
		}
		json.NewEnconder(W).Encode(cursos)
	})
	http.ListenAndServe(":8002",nil)
}