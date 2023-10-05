package main

import (
	"fmt"
)

func soma(a, b int) int {
	return a + b
}

func main() {
	var num1, num2 int
	var resultado int

	fmt.Printf("Digite 2 valores: ")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	resultado = soma(num1, num2) 
	fmt.Printf("A soma é: %d\n", resultado) 
}
