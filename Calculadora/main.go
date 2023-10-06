package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculadora (a, b, op int) int {
	if op == 1 {
		return a + b
	} else if op == 2 {
		return a - b
	} else if op == 3 {
		return a * b
	} else if op == 4 {
		return a / b
	} else {
		fmt.Printf ("Opção inválida")
		return 0
	}
}


func main() {
	file, err := os.Open("./ex.txt")
	scanner := bufio.NewScanner(file)
	fmt.Println(err)

	var list [436456]int
	var num1 int
	var num2 int
	var opr int

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

			numeros, err := strconv.Atoi(line)

			if err != nil {
				fmt.Println(err)
			}

			list[i] = numeros
			i++
		}


	fmt.Printf ("Escolha 2 posições da lista de números: ")
	fmt.Scanf ("%d" , &num1)
	fmt.Scanf ("%d" , &num2)

	fmt.Printf ("Os números escolhidos foram %d e %d" , list[num1] , list[num2])

	fmt.Printf ("\n1 - SOMA\n2 - SUBTRAÇÃO\n3 - MULTIPLICAÇÃO\n4 - DIVISÃO\n")
	fmt.Scanf ("%d" , &opr)

	resultado := calculadora(list[num1] , list[num2] , opr)
	fmt.Printf ("Resultado da operção: %d\n" , resultado)
}	
