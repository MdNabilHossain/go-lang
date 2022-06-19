package main
import "fmt"

func main(){


	// var fullName string
	// var age int
	// var gpa float32

	// fullName := "MD Nabil Hossain"
	// age := 24
	// gpa := 3.56
	// const COUNTRY = "Bangladesh"

	// fullName = "MD Nabil Hossain"
	// age = 24
	// gpa = 3.56
	
	// fmt.Printf("%v is a Business Developer \n",fullName)
	// fmt.Println(age, "Years Old")
	// fmt.Println("GPA: ",gpa)
	// fmt.Printf("%v live in %v", fullName, COUNTRY)

	var fullName string
	var num1, num2 int
	
	fmt.Print("Enter Your Name: ")
	fmt.Scan(&fullName)

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)
	// fmt.Scanf("%v %v",&num1, &num2)

	fmt.Printf("%v is a student \n", fullName)
	fmt.Printf("num1 = %v, num2 = %v", num1, num2)





}