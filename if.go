package main


import (
	"fmt"
)

func main() {
	if num := 10; num % 2 == 0 { //checks if number is even
		fmt.Println(num,"is even")
	}  else {
		fmt.Println(num,"is odd")
	}

}