package main

import (
	"fmt"
	//"os"
	"strconv"
)

func main() {
	fmt.Println("ok")

	str := "123.1"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error:", err)
		//os.Exit(1)   // ➜ Finaliza la ejecución con código de salida 1
	}

	fmt.Println(num)
}