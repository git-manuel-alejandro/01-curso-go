package intro

import "fmt"

func Retorna() int {
	return 2 * 2
}

func Saluda(saludo string) {
	fmt.Println(saludo)
}

func Print() {
	const nombre string = "manuel"
	var edad int = 32

	fmt.Printf("hello world !! %v tiene %v años \n", nombre, edad)
}
