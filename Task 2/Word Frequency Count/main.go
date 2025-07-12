package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"reflect"
)
func wordFrequency(userInput string) map[string]int{
	userInput = strings.ToLower(userInput)
	wordsOnly := ""
	for _ , ch := range userInput{
		if unicode.IsLetter(ch) ||unicode.IsSpace(ch){
			wordsOnly += string(ch)
		}
	}
	words:=strings.Fields(wordsOnly)
	count := make(map[string]int)
	for _ , word := range words {
		count[word]+=1
	}
	return count

}
func testing(){
	fmt.Println("========= Here is the test result ========= ")
	input := "Hi, hi! This is a Test. Go is fun! go is fast. Just a test. "
	expected:=map[string]int{
		"a":2,
		"fast":1,
		"fun":1,
		"just":1,
		"hi":2,
		"is":3,
		"go":2,
		"test":2,
		"this":1,	
	}
	actual:=wordFrequency(input)
	if reflect.DeepEqual(expected , actual){
		fmt.Println(" ✅ Test Passed")
	}else{
		fmt.Println(" ❌ Test Failed")
		fmt.Println("Actual:",actual)
		fmt.Println("Expected:", expected)
	}
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text to count words: ")
	userInput , _:=reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	result:=wordFrequency(userInput)
	fmt.Println(" Word frequencies:")
	for word , count:=range result{
		fmt.Printf("%s : %d\n",word , count)
	}
	testing()
	
}
