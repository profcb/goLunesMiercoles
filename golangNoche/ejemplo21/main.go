package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	//Escribir un archivo
	file, err := os.Create("archivo.txt") //
	if err != nil {
		fmt.Println("Error al crear archivo", err)
		return
	}
	defer file.Close()
	
	fmt.Println("Archivo creado con exito")

	_, err = file.WriteString("Hola de nuevo")
	if err != nil {
		fmt.Println("Error al escribir en el archivo")
		return
	}
	fmt.Println("Archivo creado con exito")


	//agregar en el archivo
	file, err := os.OpenFile("archivo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\n Nueva línea agregada.\n")
	if err != nil {
		fmt.Println("Error escribiendo:", err)
		return
	}
		*/

	//leer

	data, err := os.ReadFile("archivo.txt")
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return
	}

	fmt.Println("Contenido:")
	fmt.Println(string(data))


}
