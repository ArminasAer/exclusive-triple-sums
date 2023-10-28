package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 5, 8}
	tripplyDipply(arr)
}

func tripplyDipply(arr []int) {

	// calculation cycles are length of the array minus the total number of factors (3)
	// plus one to account for arrays starting at 0
	iterations := len(arr) - 3 + 1
	if iterations < 1 {
		fmt.Println("You need at least three integers in the array!")
		return
	}

	// set of unique integers
	sums := make(map[int]struct{})

	// 'i' is the index in the range of 'arr'; not the values
	// the values are dropped and array position is used instead
	for i := range arr {

		for n := 0; n < iterations; n++ {
			// a is the second factor in the equation so it's shifted up +1 positions
			a := n + 1 + i

			// b is the third factor in the equation so it's shifted up +2 positions
			b := n + 2 + i

			// if a or b increase over the size of the array then cycle back to 0
			// plus the current index position in the range
			// the index position is coming from 'i' in 'n + 1 + i' and 'n + 2 + i'
			if a >= len(arr) {
				a = a - len(arr)
			}
			if b >= len(arr) {
				b = b - len(arr)
			}

			// set sum to key value of map
			sums[arr[i]*arr[a]*arr[b]] = struct{}{}
		}

	}

	fmt.Println("Unique sums:", len(sums))

	fmt.Print("**************\n")
	for k := range sums {
		fmt.Printf("%v\n", k)
	}
	fmt.Print("**************\n")

	return
}
