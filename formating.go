package main  

import "fmt"
func main(){
	var decimalNumber int
	fmt.Printf ("Decimal Number: ")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("Binary Number = %b \n", decimalNumber)
	fmt.Printf("Decimal Number = %d \n ", decimalNumber)
	fmt.Printf("Hexa Decimal Number = %x \n", decimalNumber)
	fmt.Printf("Octal Number = %o \n", decimalNumber)
}