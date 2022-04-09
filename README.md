# Go Error Handling

## Error Handling
Error handling in Go is achieved by returning multiple values in a function.  
The function returns the result of processing and the error value, and by checking the error value, the case of abnormal processing can be distinguished.  
In case of normal processing, the error value is nil (Null).  
If an error occurs, the error information is stored in the error value.  

## Code
```Go
// Functions for division
func calcDiv(molec int, denomi int) (result int, err error) {
	result = molec / denomi
	return result, err
}

func main() {

	// Normal computed division
	result1, err := calcDiv(128, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Answer : %d \n", result1)

	// Divide by 0 which results in an error
	result2, err := calcDiv(128, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Answer : %d \n", result2)
}
```

## Output Sample
~ $ go build main.go  
~ $ ./main  
Answer : 32  
panic: runtime error: integer divide by zero  
  
goroutine 1 [running]:  
main.calcDiv(...)  
&nbsp;&nbsp;/Users/main.go:10  
main.main()  
&nbsp;&nbsp;/Users/main.go:26 +0x6b  
