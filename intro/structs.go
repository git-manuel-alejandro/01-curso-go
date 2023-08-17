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

	fmt.Printf("la edad de %s es %d", manuel.Nombre, manuel.Edad)

}
