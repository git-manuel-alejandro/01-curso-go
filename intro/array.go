package intro

import (
	"fmt"
	"strconv"
	"strings"
)

func Array() {

	var arreglo [10]string

	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = strconv.Itoa(i)
	}

	fmt.Println(arreglo)
}

func Slice() {
	var slice = []int{}

	for i := 0; i <= 9; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)
}

func Slice1() {
	var slice = []string{"sofia", "manuel", "caty"}

	for _, v := range slice {
		fmt.Println(v)
	}
}

func IsPalindrome(word string) bool {
	word = strings.ToLower(word)
	var reverseWord string
	for i := len(word) - 1; i >= 0; i-- {
		reverseWord += string(word[i])
	}

	if reverseWord == word {
		return true
	}

	return false

}
