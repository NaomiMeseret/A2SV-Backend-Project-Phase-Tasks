package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func getAverage(grades map[string]float64) float64 {
	var sum float64
	for _ ,g := range grades{
		sum+=g
	}
	return sum/float64(len(grades))
}
func testGetAverage(){
	fmt.Println("========= Testing the average grade ========= ")
	grades:=map[string]float64{
		"Math":88,
		"English":95,
		"Physics":80,
	}
	expected:=float64(88+95+80)/3.0
	actual:=getAverage(grades)
	if actual == expected{
		fmt.Println("Test Passed!")
	}else{
		fmt.Printf("Test failed .Expected %0.2f but got %0.2f\n", expected, actual)
	}

}
func main(){
	reader:= bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	nameInput , _ := reader.ReadString('\n')
	name:= strings.TrimSpace(nameInput)
	fmt.Print("How many subjects are you taking? ")
	numInput , _:=reader.ReadString('\n')
	numInput=strings.TrimSpace(numInput)
	numSubjects, err:=strconv.Atoi(numInput)
	if err!=nil || numSubjects<=0{
		fmt.Println("Invalid input . The number of subjects must be greater than 0.")
		return
	}
	subjects:=make([]string , numSubjects)
	grades:=make(map[string]float64 , numSubjects)
	for i:=0 ; i < numSubjects; i++{
		fmt.Printf("Enter the name of subject %d: ",i+1)
		subInput , _:=reader.ReadString('\n')
		subjects[i] = strings.TrimSpace(subInput)
		if subjects[i] == ""{
			fmt.Println("Subject name can't be empty.")
			i--
			continue
		}
		fmt.Printf("Enter your grade for %s: ",subjects[i])
		gradeInput, _:=reader.ReadString('\n')
		gradeInput = strings.TrimSpace(gradeInput)
		grade , err:=strconv.ParseFloat(gradeInput, 64)
		if err!=nil{
		fmt.Println("Invalid input.")
		i--
		continue
	}
		if grade<0 || grade>100{
			fmt.Println("Invalid input. Grade  must be between 0 and 100.")
			grade = 0
		}
		grades[subjects[i]] = grade
	}
	fmt.Printf("\n========= %s here are your results:=========\n", name )
	for _ , sub :=range subjects{
		fmt.Printf(" %s :%0.2f\n" , sub , grades[sub])
	}
	average:=getAverage(grades)
	fmt.Printf(" Your average grade is : %0.2f\n", average)

	testGetAverage()
}