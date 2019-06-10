package main

import (
	"fmt"

	genpass "../../genpass"
)

func main() {
	fmt.Println("1: Generate password, 2: Decode password")
	var number int
	fmt.Scanf("%d\n", &number)

	switch number {
	case 1:
		if err := genpass.GeneratePassword(); err != nil {
			panic(err)
		}
	case 2:
		if err := genpass.DecodePassword(); err != nil {
			panic(err)
		}
	default:
		fmt.Println("Sorry, this number", number, "isn't valid")
	}
}
