package main

import "fmt"
import "io/ioutil"
import "os"

func main(){
	filename := "file.txt"

	fmt.Println("Writer:")
	newText := os.Args[1]
	fmt.Println(string(newText))
	//go run reader.go "7. otro"

	//write := ioutil.WriteFile(filename, newText, 0777)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	showError(err)

	write, err := file.WriteString("\n" + newText)
	fmt.Println(write)
	showError(err)

	file.Close()
	
	fmt.Println("Reader:")
	readFile, err := ioutil.ReadFile(filename)
	showError(err)
	fmt.Println(string(readFile))
}

func showError(e error){
	if(e != nil){
		panic(e)
	}
}