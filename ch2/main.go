package main

import "fmt"

// Constant
const x int64 = 10

func main(){
	// Booleans 
	var flag bool
	var isAwersome = true
	
	fmt.Println(flag,isAwersome)	

	// Int
	var x int = 10
	x *= 2
	fmt.Println(x)
	
	// Complex Types
	com := complex(1,2)
	fmt.Println(real(com))
	fmt.Println(imag(com))

	// Explicit type check
	var name string = ""
	check := name == ""
	fmt.Println(check)

}

