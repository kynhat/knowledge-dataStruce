package main

import "fmt"

//a^b = a * a^b-1
func dequy(a, b int) int {

	if b == 0 {

		return 1

	}

	fmt.Println(b)
	return a * dequy(a, b-1)
}

func Ucln(a, b int) int {

	if b == 0 {

		return a

	}

	if a%b == 0 {

		return b

	}

	return Ucln(b, a%b)
}

// Giải thích: 3 số đó là 12, 13, 80.
// 12 biến đổi 3 lần: 12->6->3->10.
// 13 biến đổi 3 lần: 13->40->20->10.
// 80 biến đổi 3 lần: 80->40->20->10.
func convertNumber(n, k int) int {

	countNumber := 0

	if countNumber == 0 {

		countNumber++

	} else {

		convertNumber(2*n, k-1)

		if (n-1)%3 == 0 && ((n-1)/3)%2 == 1 {

			convertNumber((n-1)/3, k-1)

		}
	}

	return n
}

//Với n = 4 thì kết quả mong muốn là: "1234 1243 1324 1342 1423 1432 2134 2143 2314 2341 2413 2431 3124 3142 3214 3241 3412 3421 4123 4132 4213 4231 4312 4321
var b [11]bool
var x [11]int

func hoanViSoNguyen(k, n int) {

	for i := 1; i <= n; i++ {

		if b[i] {

			x[k] = i

			if k == n {

				fmt.Println(x, n)

			} else {

				b[i] = false
				hoanViSoNguyen(k+1, n)
				b[i] = true

			}
		}
	}
}

func main() {
	// fmt.Println(dequy(2, 5))
	// fmt.Println(Ucln(10, 15))
	fmt.Println((39 / 3))

}
