package main

import (
	"fmt"
)

func conversor(a float64 , opr int) float64 {
	var far float64
	var kel float64

	if (opr == 1) {
		far = (a * 1.8) + 32
		fmt.Printf ("Valor convertido para farenheit: %.2f\n" , far)
		return far
	} else if (opr == 2) {
		kel = a + 273
		fmt.Printf ("Valor convertido para kelvin: %.2f\n" , kel)
		return kel
	} else {
		fmt.Printf ("Opção inválida")
		return 0
	}
}

func main() {
	var num1 float64
	var op int

	fmt.Printf("Digite a temperatura em Celsius: ")
	fmt.Scanf("%f", &num1)

	fmt.Printf ("\n1 - CONVERSOR TO FAHRENHEIT\n2 - CONVERSOR TO KELVIN\n")
	fmt.Scanf ("%d" , &op)

	conversor(num1 , op)
	

}
