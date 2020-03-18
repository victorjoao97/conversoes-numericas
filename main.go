package main

import (
	"fmt"
)

func main() {
	decBin()
}

func decBin() {
	fmt.Println("Decimal para Binário")
	fmt.Println("Digite um número decimal")
	var num int
	_, err := fmt.Scan(&num)
	numOriginal := num
	if err != nil {
		fmt.Println(err)
	}
	var resto []int
	var dividendo int
	cont := 0

	fmt.Println("")

	for num >= 2 {
		fmt.Println(num, "/", 2)
		dividendo = num / 2
		resto = append(resto, num%2)
		num = dividendo
		fmt.Println(num, "sobra", resto[cont])
		// fmt.Println(dividendo, "-", resto)
		cont++
	}
	resto = append(resto, num)
	fmt.Println(numOriginal, "p/ Binário", reverseInts(resto))
}

// func binDec() {
// 	fmt.Println("Binário para Decimal")
// 	fmt.Println("Digite um número binário")
// 	var num int
// 	_, err := fmt.Scan(&num)
// 	numOriginal := num
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	cont := 0
// 	var soma int
// 	fmt.Println("")

// 	for cont < len(num) {
// 		soma = num[cont] * (2 * *cont)

// 		cont++
// 	}
// 	resto = append(soma)
// 	fmt.Println(numOriginal)
// }

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
