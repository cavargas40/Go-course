package main

import (
	"fmt"
	//"time"
)

func main() {
	var number1 float32 = 10
	var number2 float32 = 6
	fmt.Println("Calculator 1")
	calculator(number1, number2)

	fmt.Println("----------------")

	var number3 float32 = 44
	var number4 float32 = 7
	fmt.Println("Calculator 2")
	calculator(number3, number4)

	fmt.Println("----------------")
	fmt.Println(returnText())

	fmt.Println("----------------")
	fmt.Println("Order 1")
	fmt.Println(caps(40, "$"))

	fmt.Println("----------------")
	fmt.Println("Order 2")
	fmt.Println(caps(4, "€"))

	fmt.Println("----------------")
	pants("Red", "long", "without pockets", "Nike")

	/*
		fmt.Println("----------------")
		fmt.Println("-----TRADITIONAL ARRAYS------")
		var movies [3]string
		 movies[0] = "Truth hurts"
		 movies[1] = "Avengers"
		 movies[2] = "Superman"
		movies := [3]string{ "Truth hurts", "Avengers", "Superman" }
		fmt.Println(movies)


		fmt.Println("-----MULTI DIMENSION ARRAYS------")
		var movies [3][2]string
		movies[0][0] = "Truth hurts"
		movies[0][1] = "Avengers"
		movies[1][0] = "The Lord of the Rings"
		movies[1][1] = "The Hobbit"
		movies[2][0] = "Fast and Furious"
		movies[2][1] = "Irene me and myself"

		fmt.Println(movies)
	*/

	fmt.Println("-----SLICES------")
	movies := []string{
		"Truth hurts",
		"Avengers",
		"The Lord of the Rings",
		"The Hobbit",
		"Fast and Furious",
		"Irene me and myself"}

	movies = append(movies, "Limitless")
	movies = append(movies, "Camp Rock")

	fmt.Println(movies)
	fmt.Println(len(movies))
	fmt.Println(movies[0:3])
}

func operation(number1 float32, number2 float32, operator string) float32 {
	var res float32
	if operator == "+" {
		res = number1 + number2
	}

	if operator == "-" {
		res = number1 - number2
	}

	if operator == "*" {
		res = number1 * number2
	}

	if operator == "/" {
		res = number1 / number2
	}

	return res
}

func calculator(number1 float32, number2 float32) {
	fmt.Print("Sum:")
	fmt.Println(operation(number1, number2, "+"))

	fmt.Print("Substraction:")
	fmt.Println(operation(number1, number2, "-"))

	fmt.Print("Multiplication:")
	fmt.Println(operation(number1, number2, "*"))

	fmt.Print("Division:")
	fmt.Println(operation(number1, number2, "/"))
}

func returnText() (data1 string, data2 int) {
	data1 = "Carlos"
	data2 = 99

	return
}

//CLOSURES
func caps(order float32, currency string) (string, string, float32) {

	price := func() float32 {

		return order * 7
	}

	return "The price of the order is", currency, price()
}

func pants(attritutes ...string) {
	for _, attritute := range attritutes {
		fmt.Println(attritute)
	}
}
