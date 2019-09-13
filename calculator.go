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

}

func operation(number1 float32, number2 float32, operator string) float32 {
	var res float32
	if(operator == "+"){
		res = number1 + number2
	}

	if(operator == "-"){
		res = number1 - number2
	}

	if(operator == "*"){
		res = number1 * number2
	}

	if(operator == "/"){
		res = number1 / number2
	}

	return res
}

func calculator(number1 float32, number2 float32){
	fmt.Print("Sum:")
	fmt.Println(operation(number1, number2, "+"))

	fmt.Print("Substraction:")
	fmt.Println(operation(number1, number2, "-"))

	fmt.Print("Multiplication:")
	fmt.Println(operation(number1, number2, "*"))

	fmt.Print("Division:")
	fmt.Println(operation(number1, number2, "/"))
}

func returnText() (data1 string, data2 int){
	data1 = "Carlos"
	data2 = 99

	return
}


//CLOSURES
func caps(order float32, currency string) (string, string, float32){

	price := func() float32{

		return order*7
	}

	return "The price of the order is", currency, price()
}