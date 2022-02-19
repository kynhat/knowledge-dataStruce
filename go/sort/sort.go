package main

// package sort

import (
	"fmt"
)

func Swap(px, py *int) {
	// tempx := *px
	// tempy := *py
	// *px = tempy
	// *py = tempx

	*px, *py = *py, *px
}

func Max(a int32, b int32) int32 {

	if a == b {

		return b

	}

	if a > b {

		return a

	}

	return b
}

//----Merge Sort, Quick Sort and Bucket Sort are the most commonly used in interviews.----

// Giới thiệu giải thuật sắp xếp nổi bọt
//Bubble Sort Runtime:
// 0( n 2 )
// average and worst case. Memory:
// 0( 1) .
func BubbleSort(arrayInput []int) []int {

	length := len(arrayInput)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {

			if arrayInput[i] > arrayInput[j] {

				Swap(&arrayInput[i], &arrayInput[j])

			}
		}
	}

	return arrayInput
}

//Giới thiệu giải thuật sắp xếp chọn (Selection Sort)
// Selection Sort I Runtime: 0( n 2 ) average and worst case. Memory: 0( 1) .
// Bước 1: Thiết lập MIN về vị trí 0
// Bước 2: Tìm kiếm phần tử nhỏ nhất trong danh sách
// Bước 3: Tráo đổi với giá trị tại vị trí MIN
// Bước 4: Tăng MIN để trỏ tới phần tử tiếp theo
// Bước 5: Lặp lại cho tới khi toàn bộ danh sách đã được sắp xếp
func SelectionSort(arrayInput []int) []int {

	length := len(arrayInput)
	var indexMin int

	for i := 0; i < length-1; i++ {
		indexMin = i

		for j := i + 1; j < length; j++ {

			if arrayInput[indexMin] > arrayInput[j] {

				indexMin = j

			}
		}

		if i != indexMin {

			Swap(&arrayInput[indexMin], &arrayInput[i])

		}
	}

	return arrayInput
}

//Giới thiệu giải thuật sắp xếp chèn (Insertion Sort).
//Insertion Sort I Runtime: 0( n 2 )
// Bước 1: Kiểm tra nếu phần tử đầu tiên đã được sắp xếp. trả về 1
// Bước 2: Lấy phần tử kế tiếp
// Bước 3: So sánh với tất cả phần tử trong danh sách con đã qua sắp xếp
// Bước 4: Dịch chuyển tất cả phần tử trong danh sách con mà lớn hơn giá trị để được sắp xếp
// Bước 5: Chèn giá trị đó
// Bước 6: Lặp lại cho tới khi danh sách được sắp xếp
func InsertionSort(arrInput []int) []int {

	length := len(arrInput)
	// var index, value int

	for i := 1; i < length; i++ {

		/* chọn một giá trị để chèn */
		index := i
		value := arrInput[i]

		/*xác định vị trí cho phần tử được chèn */
		for index > 0 && arrInput[index-1] > value {

			arrInput[index] = arrInput[index-1]
			index--

		}

		/* chèn giá trị tại vị trí trên */
		arrInput[index] = value
	}

	return arrInput
}

//Shell Sort là một giải thuật sắp xếp mang lại hiệu quả cao dựa trên giải thuật sắp xếp chèn (Insertion Sort).
// interval sẽ nhận giá trị lần lượt là n/2, n/4, n/8 cho đến khi interval = 1.
//Sell Sort I Runtime: 0( n2 )
func ShellSort(arrInput []int) []int {

	length := len(arrInput)
	var interval, i, j, temp int

	for interval = length / 2; interval > 0; interval /= 2 {
		for i = interval; i < length; i++ {

			temp = arrInput[i]

			for j = i; j >= interval &&
				arrInput[j-interval] > temp; j -= interval {

				arrInput[j] = arrInput[j-interval]

			}

			arrInput[j] = temp
			fmt.Println(arrInput)
		}
	}

	return arrInput
}

func partition(arr []int, left, right int) ([]int, int) {

	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {

		if arr[j] < pivot {

			arr[i], arr[j] = arr[j], arr[i]
			i++

		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	//Trả ra 1 mảng đã được chia dưa theo đk là lấy số cuối làm chốt
	//-----------------Vi dụ------------------
	// nums := []int{5, 6, 7, 2, 1, 3}
	//Chốt là 3
	// Mảng được trả ra là [2 1 3 5 6 7] 2
	// số 2 là số lần lập
	return arr, i
}

// Giới thiệu sắp xếp nhanh (Quick Sort).
func QuickSort(arr []int, left, right int) []int {

	if left < right {

		var p int
		arr, p = partition(arr, left, right)
		arr = QuickSort(arr, left, p-1)
		arr = QuickSort(arr, p+1, right)

	}

	return arr
}

// Merge Sort Runtime: 0 ( n log ( n)) average and wo1rst case. Memory: Depends.
//Sắp xếp trộn (Merge Sort) là một giải thuật sắp xếp dựa trên giải thuật Chia để trị (Divide and Conquer).
func MergeSort(arrayInput []int) []int {

	// 1. Tìm chỉ số nằm giữa mảng để chia mảng thành 2 nửa:
	length := len(arrayInput)

	if length < 2 {

		return arrayInput

	}

	// 2. Gọi đệ quy hàm mergeSort cho nửa đầu tiên:
	arrayFirst := MergeSort(arrayInput[:length/2])

	// 3. Gọi đệ quy hàm mergeSort cho nửa thứ hai:
	arraySecond := MergeSort(arrayInput[length/2:])

	// 4. Gộp 2 nửa mảng đã sắp xếp ở (2) và (3):

	return merge(arrayFirst, arraySecond)
}

func merge(arrayInputOne []int, arrayInputTwo []int) []int {

	lengthOfArrayOne := len(arrayInputOne)
	lengthOfArrayTwo := len(arrayInputTwo)
	final := []int{}
	i := 0
	j := 0

	for i < lengthOfArrayOne && j < lengthOfArrayTwo {

		if arrayInputOne[i] < arrayInputTwo[j] {

			final = append(final, arrayInputOne[i])
			i++

		} else {

			final = append(final, arrayInputTwo[j])
			j++

		}
	}

	for ; i < len(arrayInputOne); i++ {

		final = append(final, arrayInputOne[i])

	}

	for ; j < len(arrayInputTwo); j++ {

		final = append(final, arrayInputTwo[j])

	}

	return final
}

func main() {

	// nums := []int{5, 6, 7, 2, 1, 0}
	nums := []int{5, 6, 7, 2, 1, 3}
	// fmt.Println(MergeSort(nums))
	fmt.Println(ShellSort(nums))
	// fmt.Println(partition(nums, 0, len(nums)-1))
	// fmt.Println(QuickSort(nums, 0, len(nums)-1))
}
