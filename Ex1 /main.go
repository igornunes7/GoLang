package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./ex.txt")
	scanner := bufio.NewScanner(file)
	fmt.Println(err)

	var list [436456]int

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			linhaNumerica, err := strconv.Atoi(line)

			if err != nil {
				fmt.Println(err)
			}

			list[i] = list[i] + linhaNumerica

		} else {
			i++
		}
	}

	maiorElfo := list[0]
	indexElfo := 0

	tam := len(list)
	for i = 1; i < tam; i++ {
		if list[i] > maiorElfo {
			maiorElfo = list[i]
			indexElfo = i + 1
		}
	}

	fmt.Println("Maior elfo é o número", indexElfo, "e ele carrega", maiorElfo, "calorias")
}
