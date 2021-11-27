package main

import (
	"fmt"
)

func fizzbuzz(n1,n2 int) {
	for i := 0; i <= 100; i++ {
		if i % n1 == 0 && i % n2 == 0 {
			fmt.Println("FizzBuzz")
		}else if i % n1 == 0{
			fmt.Println("Fizz")
		}else if i % n2 == 0{
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
		}
	}
}

func main(){
	fizzbuzz(3,5)
}

 