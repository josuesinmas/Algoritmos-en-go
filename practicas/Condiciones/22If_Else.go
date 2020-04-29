package main
import(
	"fmt"
	"math"
)
func pow(x,n,lim float64) float64{
	if v := math.Por(x,n)v<lim{
		return v
	}else{
		fmt.Printf("%g>=%g\n",v,lim)
	}
	//ne se puede usar v aqui, though
	return lim
}
func main(){
	fmt.Printl(
		pow(2,4,5),
		pow(3,2,10),
		pow(3,3,1¡)
	)
}