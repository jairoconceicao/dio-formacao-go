package main

import "fmt"

func soma(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func multiplicacao(numbers ...int) int {
	product := 1
	for _, n := range numbers {
		product *= n
	}
	return product
}

func divisao(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não é permitida")
	}
	return a / b, nil
}

func subtracao(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(soma(1, 2, 3))          // Output: 6
	fmt.Println(multiplicacao(1, 2, 3)) // Output: 6
	fmt.Println(divisao(10, 2))         // Output: 5
	fmt.Println(subtracao(10, 3))       // Output: 7
}
