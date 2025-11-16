package main


import "fmt"


/*
	Fibonaci
	0 - 1 - 1 - 2 - 3 - 5 - 8 - 13 - 21 - 34

	Se quiere hacer un sistema que permita a una funcion pasarle 
	cantidad de terms (esos terminos se desea que se ingresen por teclado)
	y queremos que muestre la sucesión por pantalla

	SI la cantidad de terminos es menor 2 deben informar
*/

func calcular_fibo(n int){
	serie := []int{0,1}
	for i:=2; i<n; i++{
		sig := serie[i-1] + serie[i-2]
		serie = append(serie,sig)
	}
	fmt.Println(serie)
}


func main(){

	calcular_fibo(8)

}