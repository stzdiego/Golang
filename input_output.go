package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//-------------------------------------------
	//fmt
	//-------------------------------------------
	var nombre string

	fmt.Println("Digite su nombre:")
	fmt.Scanf("%v", &nombre)
	fmt.Printf("Bienvenido %v\n", nombre)

	//-------------------------------------------
	//bufio - os
	//-------------------------------------------
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese su apellido:")
	apellido, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println("Bienvenido Sr." + apellido)
	} else {
		fmt.Println(err)
	}

	//-------------------------------------------
	//
	//-------------------------------------------

}
