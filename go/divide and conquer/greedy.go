package main

import "fmt"

// we can say that our algorithm has a
// O(n) running time.
//Coin Change Problem with Greedy Algorithm
func coinChangeGreedy(n int) {

	coins := []int{20, 10, 5, 1}
	i := 0

	for n > 0 {

		if coins[i] > n {

			i = i + 1

		} else {

			fmt.Println("ket qua", coins[i])
			n = n - coins[i]

		}
	}
}

func main() {
	coinChangeGreedy(50)
}
