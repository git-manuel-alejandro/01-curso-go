package ciclos

import "fmt"

// for while
// forever

func Ciclos() {
	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 0

	for counter < 10 {
		fmt.Println(counter)
	}

	for {
		fmt.Println("hola")
	}

}
