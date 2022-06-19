package main

import "fmt"

func main(){
	// var num1, num2 int

	// fmt.Printf("Num1 = ")
	// fmt.Scan(&num1)

	// fmt.Printf("Num2 = ")
	// fmt.Scan(&num2)
	// result := num1 + num2
	// fmt.Printf("%v + %v = %v \n",num1, num2, result)
	
	// result = num1 - num2
	// fmt.Printf("%v - %v = %v \n",num1, num2, result)
	
	// result = num1 * num2
	// fmt.Printf("%v * %v = %v \n",num1, num2, result)
	

	// triangle
	// var width, height, area float32

	// fmt.Printf("Width = ")
	// fmt.Scan(&width)

	// fmt.Printf("Height = ")
	// fmt.Scan(&height)

	// area = 0.5 * width * height

	// fmt.Printf("Area = %v", area)
	
	// Radius
	var redius float32
	
	fmt.Print("Enter redius: ")
	fmt.Scan(&redius)
	area:= 3.1416 * redius * redius
	fmt.Printf(" Redius : %v", area)
}