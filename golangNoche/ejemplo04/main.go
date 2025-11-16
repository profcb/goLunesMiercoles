/*

	Tenemos las notas de 
	Juan se saco un 9.5
	Tomas un 10
	Elvira un 5
	Yonathan 3

	Necesitamos saber el promedio del curso, usar lo aprendido

	Curso se desea que sea un MAP

*/

package main

import (
	"fmt"
)

//nil
// [] corchetes  {} llaves 
func main (){
	curso := map[string]float64{
		"Juan":9.5,
		"Tomas":10,
		"Elvira":5,
		"Yonathan":3,
	}

	var sumatoria float64 = 0

	for _,nota := range curso{
		sumatoria += nota 
	}
	promedio := sumatoria / float64(len(curso))
	fmt.Println(promedio)

}