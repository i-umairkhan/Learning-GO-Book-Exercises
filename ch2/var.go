package main

import "fmt"

const x int64 = 10

func main() {
	var x int = 10
	fmt.Println(x)
	x , y := 20 , 10
	x = 30
	fmt.Println(x)
	fmt.Println(y)
}

