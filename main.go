package main

import (
	"fmt"
)

func main() {
	arr := []int{70, 4, 5, 8, 9}
	tripplyDipply(arr)
}

func tripplyDipply(arr []int) {

	iterations := len(arr) - 3 + 1
	if iterations < 1 {
		fmt.Println("You need at least three integers in the array!")
		return
	}

	// sums := []int{}
	sums := make(map[int]struct{})

	for i := range arr {

		for n := 0; n < iterations; n++ {
			a := n + 1 + i
			b := n + 2 + i

			if a >= len(arr) {
				a = a - len(arr)
			}

			if b >= len(arr) {
				b = b - len(arr)
			}

			sums[arr[i]*arr[a]*arr[b]] = struct{}{}
		}

	}

	fmt.Println("Total sums:", len(sums))

	fmt.Print("**************\n")
	for k := range sums {
		fmt.Printf("%v\n", k)
	}
	fmt.Print("**************\n")

	return
}
