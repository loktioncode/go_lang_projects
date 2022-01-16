package main

import "fmt"

func main() {
	// Array
	var fruitArr [2]string
	vegieArr := []string{"Tomato", "Rap"}

	//assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	fmt.Println(vegieArr)
	fmt.Println(len(vegieArr))
	fmt.Println(vegieArr[0:3])

}
