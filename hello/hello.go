package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	for i := 9999; i < 2000000; i++ {
		if i%2 == 0 {
			fmt.Println(i, " jest parzysta")
		} else if i%i == 0 && i%1 == 0 {
			fmt.Println(i, " jest liczba pierwsza")
		}
	}
}
