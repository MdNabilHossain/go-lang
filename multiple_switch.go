package main

import "fmt"

func class(){
	var  classNumber int
	fmt.Printf("Enter Your Class: ")
	fmt.Scan(&classNumber)

	switch classNumber {
	case 1,2,3,4,5:
		fmt.Printf("Primary \n")
	
	case 6,7,8,9,10:
		fmt.Printf("Secondary \n")

	case 11,12:
		fmt.Printf("Higher Secondary \n")
	default:
		fmt.Printf("Not a valid Class \n")
		}	
}