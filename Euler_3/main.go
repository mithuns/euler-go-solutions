//https://projecteuler.net/problem=3
//find the largest prime factor of the given N
package main

import (
	"fmt"
	"time"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var numCases int
	fmt.Scanf("%d", &numCases)

	for i := 1; i <= numCases; i++ {
		var N int
		fmt.Scanf("%d", &N)
		LargestPrimeFactor(N)
	}
}
func LargestPrimeFactor(N int) {
	defer elapsed("findLargestPrimeFactor")()
	var copyNumber = N
	var i = 2
	loopCounter := 0
	for i <= copyNumber {
		if copyNumber%i == 0 {
			copyNumber = copyNumber / i
		}
		i = i + 1
		loopCounter++
	}
	fmt.Println("loopCounter:", loopCounter)
	fmt.Println(i - 1)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start).Milliseconds())
	}
}
