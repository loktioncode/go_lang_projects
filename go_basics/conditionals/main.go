package main

import "fmt"

func main() {
	// x := 5
	// y := 10

	// if x > y {
	// 	fmt.Printf("%d & %d are not equal\n", x, y)
	// } else if x == y {
	// 	fmt.Printf("%d & %d are equal\n", x, y)
	// }

	//range with map
	emails := map[string]string{"Bob": "b@gmail.com", "Sharon": "shaz@gmail.com"}

	for k, v := range emails {
		fmt.Printf("Users are %s\n", k)
		fmt.Printf(" %s email is %s\n", k, v)
	}

	//we can define a function

}
