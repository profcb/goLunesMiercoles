/*
En Go, los términos marshal y unmarshal vienen del paquete encoding/json.
Se usan para convertir datos entre estructuras de Go ↔ JSON.

*/

package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	/*
		//ejemplo para pasar a json
		p := Persona{"Juan", 30}
		fmt.Println(p)

		jsonM, err := json.Marshal(p)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println(string(jsonM))
	*/
	jsonStr := `{"nombre":"Carlos","edad":30}`
	var p Persona
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("El nombre es",p.Nombre,"La edad es", p.Edad)

}
