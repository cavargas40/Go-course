package main

import (
	"fmt"
	//"os"
	//"strconv"
)

func main() {
	fmt.Println("************ My Program with Golang ************")

	/*
	//go run program.go Carlos 27
	fmt.Println("Hello "+os.Args[1] + " Welcome to the program with Golang")
	//Conditionals
	// && AND
	// || OR
	// != DIFFERENT
	// == EQUAL
	age, _ := strconv.Atoi(os.Args[2])

	if age >= 18 && age <= 99 {
		fmt.Println("You're of age")
	} else {
		fmt.Println("You're MINOR or too old")
	}
	*/

	/*
	//go run program.go 1
	number,_ := strconv.Atoi(os.Args[1])
	if number%2 == 0 {
		fmt.Println("Pair number")
	} else {
		fmt.Println("Odd number")
	}
	*/

	//loops
	max := 26
	for i := 1; i <= max; i++{
		if i%2 == 0 {
			fmt.Println("Pair number:", i)
		} else {
			fmt.Println("Odd number:", i)
		}
		
	}

	fmt.Println("--------------------------------------")

	movies := []string{"Fast & Furious", "Fight Club", "Avengers", "Brazzers" } 

	for _, movie := range movies {
		fmt.Println(movie)
	}
}
