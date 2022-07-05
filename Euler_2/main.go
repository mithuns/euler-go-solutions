//https://projecteuler.net/problem=2
//find the sum of the even-valued terms of Fibonacci sequence for given N
package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var numCases int
	fmt.Scanf("%d", &numCases)

	for i := 1; i <= numCases; i++ {
		var N int
		fmt.Scanf("%d", &N)
		EvenSumFibo(N)
	}
}

func EvenSumFibo(N int) {
	if N == 0 || N == 1 {
		fmt.Println(0)
		return
	}
	if N == 2 {
		fmt.Println(2)
		return
	}
	var a, b, c = 1, 2, 0
	var sum = b
	for i := 3; i <= N; i++ {
		c = a + b
		if c >= N {
			break
		}
		if c%2 == 0 {
			sum = sum + c
		}
		a = b
		b = c
	}
	fmt.Println(sum)
}
