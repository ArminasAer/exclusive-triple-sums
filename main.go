package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tripplyDipply(arr)
}

func tripplyDipply(arr []int) {

	iterations := len(arr) - 3 + 1
	if iterations < 1 {
		fmt.Println("You need at least three integers in the array!")
		return
	}

	sums := []int{}

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

			sums = append(sums, arr[i]*arr[a]*arr[b])
		}

	}

	fmt.Println("Total sums:", len(sums))

	fmt.Println(sums)
	return
}
