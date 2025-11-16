// De una sola linea

/*
 comentario
 multilineal
*/

/*

Tipos de datos básicos
Tipo 											Descripción				Ejemplo
int, int8, int16, int32, int64					Enteros con signo		var edad int = 25
uint, uint8, uint16, uint32, uint64				Enteros sin signo		var cantidad uint = 100
float32, float64								Números decimales		var precio float64 = 9.99
bool											Verdadero o falso		var activo bool = true
string											Cadena de texto			var nombre string = "GoLang"
rune 											Caracter, una sola letra var letra rune = "h"

Operadores Aritmeticos

+,-,*,/,%

Operadores Auntoincr y otros

++
--
+=1
-=1


Relacionales

==, !=, <, >, <=, >=

Lógicos
Simbolo Significado
&& 		(AND)
|| 		(OR)
! 		(NOT)

*/

package main

import "fmt"

func main() {
	var nombre string = "Roberto"
	var edad int = 28
	var sube float64 = 1001.5
	amigo := "Juan"
	edad_amigo := 32
	sube_amigo := 2002.82
	fmt.Println(nombre, "tiene", edad, "años y en su sube tiene", sube)
	fmt.Println(amigo, "tiene", edad_amigo, "años y en su sube tiene", sube_amigo)

	if edad > 18 {
		fmt.Println("Roberto es mayor de edad")
	} else {
		fmt.Println("Roberto es menor de edad")

	}

	var dia string
	fmt.Print("Ingrese el día de hoy: ")
	fmt.Scanln(&dia)

	switch dia {
	case "lunes","Lunes":
		fmt.Println("Primer día laboral")
	case "viernes":
		fmt.Println("Ultimo día laboral")
	default:
		fmt.Println("En otro caso")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	/*
	for{
		//bucle infinito
		if flag < 5{
			break
		}
	}
		continue -> pasar a la siguente vuelta
		break -> corto totalmente el bucle
	*/


}
