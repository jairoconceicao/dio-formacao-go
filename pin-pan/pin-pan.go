package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 {
				fmt.Print("Pin! ")
			}
			if i%5 == 0 {
				fmt.Print("Pan! ")
			}
			fmt.Println("")
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}
