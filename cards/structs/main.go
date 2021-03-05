package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// manthan := person{
	// 	firstName: "Manthan",
	// 	lastName:  "Surkar",
	// }
	// fmt.Println(manthan.firstName)
	manthan := person{
		firstName: "Manthan",
		lastName:  "Surkar",
		contactInfo: contactInfo{
			email:   "manthan.surkar@gmail.com",
			zipCode: 445001,
		},
	}
	fmt.Printf("%+v", manthan)
}
