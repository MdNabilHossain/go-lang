package main

import "fmt"

func main() {
	//   for x := 10; x<= 100; x+=2{
	//     fmt.Printf("%v \n",x)
	//   }

	//series -> 1 + 2 + 3 + 4 + 5
  var startNumber, endNumber int
  fmt.Printf("Enter Start Number: ")
  fmt.Scan(&startNumber)

  fmt.Printf("Enter End Number: ")
  fmt.Scan(&endNumber)
  
	sum := 0
	for i := startNumber; i <= endNumber; i++ {
		sum = sum + i
		 
	}

	fmt.Printf("Sum: %v \n", sum)
}
