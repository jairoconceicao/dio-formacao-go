package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%v é divisivel por 3\n", i)
		}
	}
}
