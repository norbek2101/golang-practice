package main

import (
    "fmt"
)

func fibo(n int){
	x := 0
	y := 1
	if n <= 0{
		fmt.Printf("Please enter the positive integer!\n")
	}else{
		for count:= 0; count <= n; count++{
			fmt.Println(x)
			t := x
			x = y
			y = t + x
		}
	}
	
}


func main() {
    fibo(10)
}