// a must to add before starting a new file
package main

// a standard module allows you to display msgs and texts in the console
import "fmt"

func main() {

	// Exercise 1 - Print hello world
	fmt.Println("Hello world!")

	// Exercise 2 - Greetings
	fmt.Println()
	var name,surname string

	fmt.Print("Please enter your name: ")
	fmt.Scan(&name, &surname)
	fmt.Println("My name is "+name+" "+surname)

	// Exercise 3 - Area of a rectangle
	fmt.Println()
	var length float32
	var width float32

	var area float32

	fmt.Print("Please enter the width in cm : ")
	fmt.Scan(&width)
	fmt.Print("Please enter the length in cm : ")
	fmt.Scan(&length)

	area = length * width

	fmt.Println("The area of the rectangle is : ",area,"cm")

	
}

