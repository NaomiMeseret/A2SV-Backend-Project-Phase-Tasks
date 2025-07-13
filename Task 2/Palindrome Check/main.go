package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)
func palindrome(userInput string)bool{
	userInput = strings.ToLower(userInput)
	wordsOnly := ""
	for  _ , ch := range userInput{
		if unicode.IsLetter(ch) || unicode.IsDigit(ch){
			wordsOnly += string(ch)
		}
	}
	reversedString := ""
	for i := len(wordsOnly)-1; i>=0; i--{
		reversedString+=string(wordsOnly[i])
	}
	return reversedString == wordsOnly

}
type TestCase  struct{
	input string
	expected bool
}
func testing(){
	fmt.Println("=========== Here is the test result =========== ")
cases:=[]TestCase{
	{input:"Was it a car or a cat I saw?" ,  expected: true},
	{input:"1A Toyota! Race fast, safe car! A Toyota1" , expected: true},
	{input: "Hello, World!",  expected: false},
	{input: "12321",  expected: true},
	{input: "12345",  expected: false},
}
for i , tc := range cases{
	actual:=palindrome(tc.input)
	if actual == tc.expected{
		fmt.Printf("✅ Test %d passed: %q -> %v\n", i+1,tc.input, actual)
	}else{
		fmt.Printf("❌ Test %d failed: %q -> %v\n", i+1,tc.input, actual)

	}
}
}
func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text:")
	userInput , _:= reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	result := palindrome(userInput)
	if result{
		fmt.Println("It is a palindrome!")
	}else{
		fmt.Println("It is not a palindrome!")
	}

	testing()

}