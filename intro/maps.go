package intro

import "fmt"

func Maps() {
	var mapa = make(map[string]int)
	mapa["manuel"] = 32
	mapa["nolo"] = 2
	mapa["sofia"] = 8

	for i, v := range mapa {
		fmt.Println(i, v)
	}

	caty, exist := mapa["caty"]
	fmt.Println(caty, exist)
}
