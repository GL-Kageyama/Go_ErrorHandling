package main

import (
	"fmt"
	"log"
)

// Functions for division
func calcDiv(molec int, denomi int) (result int, err error) {
	result = molec / denomi

	// Returns including error types
	return result, err
}

// Main function
func main() {

	// Normal computed division
	result1, err := calcDiv(128, 4)
	if err != nil { // In case of normal processing, no value is stored in err.
		log.Fatal(err)
	}
	fmt.Printf("Answer : %d \n", result1)

	// Divide by 0 which results in an error
	result2, err := calcDiv(128, 0)
	if err != nil { // In case of abnormality, the value is stored in err.
		log.Fatal(err)
	}
	fmt.Printf("Answer : %d \n", result2)

	// Logic not processed because an error occurs immediately before
	result3, err := calcDiv(128, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Answer : %d \n", result3)
}

// =================================
//           Output Sample
// =================================
// ~ $ go build main.go 
// ~ $ ./main 
// Answer : 32 
// panic: runtime error: integer divide by zero

// goroutine 1 [running]:
// main.calcDiv(...)
// 	/Users/***/main.go:10
// main.main()
// 	/Users/***/main.go:26 +0x6b