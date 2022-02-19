package main

import "fmt"

//time O(logN)
//Vì thuật toán Binary Search yêu cầu mảng đã được sắp xếp. Thế nên, đầu vào của chúng ta là một mảng đã được sắp xếp.

func binarySearch(arr []int, number int) int {

	length := len(arr)
	low := 0
	high := length - 1

	for high >= low {
		mid := (low + high) / 2

		// Nếu phần tử có ở chính giữa
		if arr[mid] == number {

			return mid

		}

		if arr[mid] < number {

			low = mid + 1

		} else {

			high = mid - 1

		}
	}

	return -1
}

func binarySearchBetter(arr []int, l, r int, number int) int {

	if l <= r {
		mid := (l + r) / 2

		// Nếu phần tử có ở chính giữa
		if arr[mid] == number {

			return mid

		}

		// Nếu phần tử nhỏ hơn giữa, thì nó chỉ có thể
		// hiện diện trong mảng con bên trái
		if arr[mid] > number {

			return binarySearchBetter(arr, l, mid-1, number)

		} else {

			// Ngược lại, phần tử chỉ có thể có mặt
			// trong mảng con bên phải
			return binarySearchBetter(arr, mid+1, r, number)

		}
	}

	// Nếu phầnt tử không có trong mảng
	return -1
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6}
	result1 := binarySearch(items, 9)
	result2 := binarySearchBetter(items, 0, len(items)-1, 5)

	if result1 == -1 || result2 == -1 {

		fmt.Println("Phần tử không tồn tại.")

	} else {

		fmt.Println("Phần tử được tìm thấy tại vị trí: ", result1, result2)

	}
}
