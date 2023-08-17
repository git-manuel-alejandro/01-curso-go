package intro

import "fmt"

func Punteros() {
	var nombre = "manuel"
	var direccion = &nombre
	fmt.Println(&nombre) // la direccion de memoria
	fmt.Println(*direccion)
	*direccion = "pedro"
	fmt.Println(nombre)

}
