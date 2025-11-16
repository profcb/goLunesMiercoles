package main

import (
	"fmt"
	//"time"
)

type User struct {
	Nombre string
	Email  string
	Edad   int
	//Registro time.Time
}

func (x User) Saludar() {
	fmt.Println("Hola mi nombre es", x.Nombre)
}

func saludar() {
	fmt.Println("Hola")
}

func sumar(a int, b int) int {
	return a + b
}

func main() {
	// definidas como en cualquier lenguaje
	//minuscula -> edad
	//guiones bajos -> edad_uno
	//camel case -> edadUno

	//u:= Usuario{"Carlos","Carlos@gmail.com",26}
	u := User{ //""
		Edad:  32,
		Email: "carlos@gmail.com",
		//Registro: time.Now(),
	}
	u.Nombre = "Carlos"
	fmt.Println(u)
	saludar()
	r := sumar(5, 2)
	fmt.Println("El resultado de sumar es", r)

	u.Saludar()
}
