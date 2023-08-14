package intro

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func Structs() {

	var manuel = Persona{
		Nombre: "manuel",
	}
	manuel.Edad = 32

	fmt.Println(manuel)

}
