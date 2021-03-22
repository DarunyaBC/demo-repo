package main

import "fmt"

func add_values(a int, b int) int {

	var c = a + b

	return c
}

func main() {

	var a, b, c int

	fmt.Println("Enter first Number: ")

	fmt.Scan(&a)

	fmt.Println("Enter Second Number: ")

	fmt.Scan(&b)

	c = add_values(a, b)

	fmt.Printf("Sum: %d \n", c)

}
