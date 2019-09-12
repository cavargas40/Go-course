package main

import (
	"fmt"
	//"time"
)

func main() {
	/*
		VARS
		time.Sleep(time.Second * 5)
	*/
	/*
		var sum int = 8 + 9
		var substraction int = 6 - 4
		var name string = "Carlos"
		var lastname string = "Vargas"
		//lastname = "Lopez"

		country := "Colombia"
		// country = 12      ///hello-world.go:19:10: cannot use 12 (type int) as type string in assignment

		var boolean bool = true
		var number float32 = 12

		const year int = 2019


				PRINT

			fmt.Println("Hello world from Go with " + name + " " + lastname + " from " + country)
			fmt.Println(sum)
			fmt.Println(substraction)
			fmt.Println(boolean)
			fmt.Println(number)
			fmt.Println(year)
	*/

	/*
			OPERATORS

		var number1 float32 = 10
		var number2 float32 = 6

		fmt.Print("Sum:")
		fmt.Println(number1 + number2)

		fmt.Print("Substraction:")
		fmt.Println(number1 - number2)

		fmt.Print("Multiplication:")
		fmt.Println(number1 * number2)

		fmt.Print("Division:")
		fmt.Println(number1 / number2)

		//fmt.Print("Mod:")
		//fmt.Println(number1 % number2)
	*/

	/*
		Structs
	*/
	var blackCup = Cup{"Nike", "Black", 25.50, false}

	fmt.Println(blackCup)
	fmt.Println(blackCup.brand)
}

type Cup struct {
	brand string
	color string
	price float32
	flat  bool
}
