package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/sethvargo/go-password/password"
)

// Functions

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
func word_count(text string){
	var word_count int
	var words [] string

	// number of words in the text
	words = strings.Fields(text)
	word_count = len(words)
	fmt.Println(text)

	fmt.Println("words count = ", word_count)

	// number of characters in the text
	fmt.Println("characters count = ", len(text))
}

// find a specific letter in the text
func find_character(letter,text string){
	// Trim any newline characters and get the first rune
	letter = strings.TrimSpace(letter)
	char := []rune(letter)[0]

	char_count := 0
	for _, c := range text {
		if c == char {
			char_count++
		}
	}

	fmt.Printf("Character '%c' count = %d\n", char, char_count)
}	

// Function for exercise 9 - Simple web server
func handler(w http.ResponseWriter, r *http.Request) {

	// HTML template
	tmpl, err := template.ParseFiles("page.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// Function for exercise 10 - Password checker
func password_checker(password string){
	var upper,lower,number bool
	
	if len(password) < 8 {
		fmt.Println("Password must contain atleast 8 characters")
	} else {
		for _, r := range password {
			if unicode.IsUpper(r){
				upper = true
			} else if unicode.IsLower(r){
				lower = true
			} else if r >= '0' && r <= '9'{
				number = true
			}
		}
	}

	if upper && lower && number {
		fmt.Print("Password is valid.")
	}else{
		fmt.Print("Password is invalid.")
	}
}

func random_password(length,symbols,digits int, upper,repeat string){
	var setUpper,setRepeat bool
	if length <= symbols + digits {
		fmt.Println("The number of symbols and digits should be atleast equal to the length of the password.")
	} else {
		if upper == "yes" {
			setUpper = true
		}else{
			setUpper = false
		}

		if repeat == "yes" {
			setRepeat = true
		}else {
			setRepeat = false
		}
	}

	rand_pass := password.MustGenerate(length, symbols, digits, setUpper, setRepeat)
	fmt.Println(setUpper , setRepeat)
	fmt.Println("\nRandom password: ", rand_pass)
}

// Function for exercise 12 - count number of lines in the data retrieved from the file
func data_processing(data []byte, fileName string){

	fmt.Printf("\nFile Name: %s", fileName)
    fmt.Printf("\nSize: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)

	new_data:= string(data)
	
	// Number of lines in the data
	lines := strings.Split(new_data, ".")
	fmt.Printf("Number of lines: %d\n", len(lines))

	// Find a specific word in the data
	var word string

	fmt.Print("Please enter the word you need to find : ")
	fmt.Scan(&word)

	if strings.Contains(new_data, word) {
		fmt.Printf("The word '%s' is found in the file\n", word)
	}else{
		fmt.Printf("The word '%s' is not found in the file\n", word)
	}
}

// Function for exercise 12 - write to a file
func write_to_file(){
	fmt.Println("\nWriting to a file in Go lang")

	file, err := os.Create("new_file.txt")
     
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

	defer file.Close()

}

// Main method

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
	var length1 float32
	var width float32

	var area float32

	fmt.Print("Please enter the width in cm : ")
	fmt.Scan(&width)
	fmt.Print("Please enter the length in cm : ")
	fmt.Scan(&length1)

	area = length1 * width

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

	var text,letter string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a text : ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	word_count(text)

	fmt.Print("Please enter the letter you need to find : ")
	letter, _ = reader.ReadString('\n')
	find_character(letter,text)

	//Exercise 9 - Simple web server
	fmt.Println()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))


	// Exercise 10 - Password checker
	fmt.Println()

	var password string

	fmt.Println("Make sure the password has\n- Atleast 8 characters\n- Atleast 1 uppercase and lowercase letter\n- Atleast 1 number")
	fmt.Print("Please enter a password : ")
	fmt.Scan(&password)

	password_checker(password)

	// Exercise 11 - random password checker
	fmt.Println()

	// using must Generate package shortcut
	var length,symbols,digits  int
	var upper, repeat string

	fmt.Println("To generate a random password according to your need , please fill the below sections ;")
	fmt.Print("Please enter the length of the password : ")
	fmt.Scan(&length)
	fmt.Print("Please enter the number of digits you need in the password : ")
	fmt.Scan(&digits)
	fmt.Print("Please enter the bnumber of symbols you need in the password : ")
	fmt.Scan(&symbols)
	fmt.Print("Please enter whether you need any uppercase letters in your password (yes/no) : ")
	fmt.Scan(&upper)
	fmt.Print("Please enter whether you allow to repeat the characters in the password (yes/no) : ")
	fmt.Scan(&repeat)

	random_password(length,symbols,digits,upper,repeat)

	// Exercise 12 - File I/O
	fmt.Println()

	// Reading a file
	fmt.Printf("Reading a file in Go lang\n")
    fileName := "test.txt"

	data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
	}

	data_processing(data,fileName)
	write_to_file()
}

