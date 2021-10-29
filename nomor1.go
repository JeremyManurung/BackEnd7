package main

import "fmt"


func fibo(n int) int {
	var tampung = []int{0, 1, 1}
	var tes int
	if len(tampung) > n {
		return tampung[n]
	} else {
		tes = fibo(n-1) + fibo(n-2)
	}

	tampung = append(tampung, tes)
	return tes
}

//

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