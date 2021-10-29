package main

import "fmt"

func fibo(n int) int {
	tampung := []int{0, 1, 1}
	tes := 3
	for tes <= n {
		tampung = append(tampung, tampung[tes-1]+tampung[tes-2])
		tes++
	}

	return tampung[n]
}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 34

	fmt.Println(fibo(10)) // 55

}