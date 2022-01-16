package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// variables
	// var name  = "Bere" works also
	var name string = "Elisha"
	var age int = 45
	surname := "Bere" //only allowed inside function main
	size := 1.3
	email, address := "elis@eli.com", "22 XXXX, CA"

	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", surname)
	fmt.Println("size is: ", size)
	fmt.Println(email, address)
}
