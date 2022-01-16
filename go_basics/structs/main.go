package main

import (
	"fmt"
	"strconv"
)

//define person struct
type Person struct {
	firstName, lastName, gender string
	age                         int
	height                      float64
	maritalStatus               string
}

//value reciever ----
//p identifies the
func (p Person) greeting() string {
	return p.firstName + " is " + strconv.Itoa(p.age) + "years old and is " + p.maritalStatus
	// return "Hello my name is " + p.firstName + " " + p.lastName
}

//pointer reciever
//allows to change struct values
func (p *Person) hasBirthday() {
	p.age++
}

//pointer reciever
func (p *Person) getMarried(maritalStatus string) {
	if p.maritalStatus == "single" {
		return
	} else {
		p.maritalStatus = maritalStatus
	}
}

func main() {

	//initialize object
	person1 := Person{
		firstName:     "Elisha",
		lastName:      "Bere",
		age:           3,
		gender:        "male",
		height:        50.5,
		maritalStatus: "single",
	}

	// fmt.Println(person1)
	// fmt.Println(person1.age)
	person1.hasBirthday()
	person1.getMarried("married")

	fmt.Println(person1.greeting())

}
