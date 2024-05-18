// a must to add before starting a new file
package main

// a standard module allows you to display msgs and texts in the console
import "fmt"

func main() {

	// // Exercise 1 - Print hello world
	// fmt.Println("Hello world!")

	// // Exercise 2 - Greetings
	// fmt.Println()
	// var name,surname string

	// fmt.Print("Please enter your name: ")
	// fmt.Scan(&name, &surname)
	// fmt.Println("My name is "+name+" "+surname)

	// // Exercise 3 - Area of a rectangle
	// fmt.Println()
	// var length float32
	// var width float32

	// var area float32

	// fmt.Print("Please enter the width in cm : ")
	// fmt.Scan(&width)
	// fmt.Print("Please enter the length in cm : ")
	// fmt.Scan(&length)

	// area = length * width

	// fmt.Println("The area of the rectangle is : ",area,"cm")

	// Exercise 4 - Simple calculator 
	fmt.Println()
	var option, num1, num2, ans float32

	fmt.Println("_Please choose the operation you want to perform_")
	fmt.Println("addition - 1\nsubstraction - 2\nmultiplication - 3\ndivision - 4")
	fmt.Print("Please enter the option : ")
	fmt.Scan(&option)
	fmt.Print("Please enter the first number : ")
	fmt.Scan(&num1)
	fmt.Print("Please enter the second number : ")
	fmt.Scan(&num2)

	if (option == 1){
		ans = num1 + num2
		fmt.Print("Answer = ",ans)
	}else if(option == 2){
		ans = num1 - num2
		fmt.Print("Answer = ",ans)
	}else if(option == 3){
		ans = num1 * num2
		fmt.Print("Answer = ",ans)
	}else if(option == 4){
		ans = num1 / num2
		fmt.Print("Answer = ",ans)
	}else{
		fmt.Print("Invalid option")
	}
}

