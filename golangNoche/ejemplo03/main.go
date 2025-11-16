package main

import "fmt"

/*

	Arrays son estaticos en su tamaño y en su tipo
	var arr [3]int = [3]int{10,20,30}

	Slices son 'arrays' dinamicos en su tamaño y estaticos en su tipo
	numeros := []int{10,20,30}
	numeros = append(numeros, 4, 5)



	slice vacio
	var numeros []int
	numeros := make([]int, 0)
	nombres := []string{}


	append solo en los slice
	len en todo , me permite contar sin contar

	map(Mapas - Mapeo)

	paises := map[string]string{
		"AR":"Argentina"
		"UY":"Uruguay"
		"BR":"Brasil"
		"PY":"Paraguay"
		"PE":"Peru"
	}

	capitales := make(map[string]string)
	capitales["Argentina"] = "Buenos Aires"
	capitales["Uruguay"] = "Montevideo"

*/

func main() {
	numeros := []int{10, 20, 30}
	numeros = append(numeros, 4, 5)
	fmt.Println(numeros)
	//for i := 0; i < len(numeros); i++ {
	//	fmt.Println(numeros[i])
	//}

	for j, valor := range numeros {
		fmt.Println(j, valor)
	}

	capitales := make(map[string]string)
	capitales["Argentina"] = "Buenos Aires"
	capitales["Uruguay"] = "Montevideo"

	for pais, capital := range capitales {
		fmt.Println(pais, "->", capital)
	}

}
