package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n < 0 {
		return -1 //, fmt.Errorf("Can not find fibonacci for negative integers")
	}
	if n <= 1 {
		return n
	}
	var ans int = fibonacci(n-1) + fibonacci(n-2)
	return ans
}

func main() {
	fmt.Println("Hello, Wod!")
	for i := -2; i <= 10; i++ {
		fmt.Printf("%d, %d\n", i, fibonacci(i))
	}
}
