package main

import (
	"bufio"
	"fmt"
	"os"

	"math/rand"
	"strings"
)

// Functions for exercise 6 - Temperature converter
func f_to_c(farenheit, answer float32) {
	answer = (farenheit - 32) * 5/9
	fmt.Println("The temparature in Celcius is : ",answer,"°C")
}
func c_to_f(celcius,answer float32) {
	answer = (celcius * 9/5)+32
	fmt.Println("The temparature in Farenheit is : ",answer,"°F")
}

// Function for exercise 7 - Loan calculator  
func loan_calculation(amount, rate, time float64){
	var monthly_payment float64

	monthly_payment = (amount * rate)/100
	fmt.Print("\nMonthly loan payment : ",monthly_payment/time)
}

// Function for exercise 8 - Text analyzer
func counting_numbers(text string){
	var word_count int
	var words [] string

	// number of words in the text
	words = strings.Fields(text)
	word_count = len(words)
	fmt.Println(text)

	fmt.Println("words count = ", word_count)

	// number of characters in the text
	fmt.Println("characters count = ", len(text))
	fmt.Println(text[0])

	// find a specific letter in the text
	//var letter string

	// fmt.Print("Please enter the letter you need to find : ")
	// fmt.Scan(&letter)
	

	// for i=0; i<word_count; i++{

	// }


}

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

	// Exercise 5 - Number guessing game
	fmt.Println()
	var random_num, guess int
	random_num = rand.Intn(10)

	fmt.Print("You got 3 chances to guess the number between 0 and 10\n")

	for i:=0; i<3; i++{
		fmt.Print("Please enter your guess no ",i+1," : ")
		fmt.Scan(&guess)
		if (guess==random_num){
			fmt.Println("Congratualations! You guessed correctly.")
			break
		}else{
			fmt.Println("Incorrect guess. Let's try again.")
		}
	}
	fmt.Print("\nThe correct number is : ", random_num)

	// Exercise 6 - Temperature converter
	fmt.Println()

	var farenheit, celcius, answer float32
	var choice int

	fmt.Println("Please choose the conversion you want to perform\nFarenheit to Celcius - 1\nCelcius to Farenheit - 2")
	fmt.Print("Please enter the option : ")
	fmt.Scan(&choice)
	if (choice == 1){
		fmt.Print("\nPlease enter the temparature in Fahrenheit : ")
		fmt.Scan(&farenheit)
		f_to_c(farenheit,answer)
	} else if (choice == 2){
		fmt.Print("\nPlease enter the temparature in Celcius : ")
		fmt.Scan(&celcius)
		c_to_f(celcius,	answer)
	} else{
		fmt.Print("\nInvalid option")
	}

	// Exercise 7 - Loan calculator
	fmt.Println()

	var amount, rate, time float64
	fmt.Print("Please enter the pincipal amount of your loan : ")
	fmt.Scan(&amount)
	fmt.Print("Please enter the rate of interest percentage : ")
	fmt.Scan(&rate)
	fmt.Print("Please enter the time period(months): ")
	fmt.Scan(&time)

	loan_calculation(amount, rate, time)

	// Exercise 8 - Text analyzer
	fmt.Println()

	var text string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a text : ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)

	counting_numbers(text)

	// //Exercise 9 - Simple web server
	// fmt.Println()

	// // Exercise 10 - Password checker
	// fmt.Println()

	// // Exercise 11 - random password checker
	// fmt.Println()

	// // Exercise 12 - File I/O
	// fmt.Println()

}

