package main

import "fmt"

//closure function
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	sum := adder()

	for i := 0; i < 10; i++ {
		//iterates and adds next number in the loop
		fmt.Println(sum(i))
	}

}
