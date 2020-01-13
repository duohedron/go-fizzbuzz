package fizzbuzz

import (
	"fmt"
)

// Fizz variable
var Fizz = "Fizz"

// Buzz variable
var Buzz = "Buzz"

// FizzBuzz function
func FizzBuzz(n int) <-chan string {
	out := make(chan string, n)
	go func() {
		for i := 1; i <= n; i++ {
			result := ""
			if i%3 == 0 {
				result += Fizz
			}
			if i%5 == 0 {
				result += Buzz
			}
			if result == "" {
				result = fmt.Sprintf("%v", i)
			}
			out <- result
		}
		close(out)
	}()
	return out
}
