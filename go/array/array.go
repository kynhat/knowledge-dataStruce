package array

import "fmt"

// var inputs []int = []int{1, 2, 3, 4, 4, 4}

func hourglassSum(arr [][]int32) int32 {

	sum := int32(0)

	for i := 0; i < 4; i++ {
		for j := 1; j < 5; j++ {

			temp := arr[i][j-1] + arr[i][j] + arr[i][j+1] + arr[i+1][j] + arr[i+2][j-1] + arr[i+2][j] + arr[i+2][j+1]

			if i == 0 && j == 1 {

				sum = temp

			} else {

				sum = max(sum, temp)

			}
		}
	}

	return sum
}

func max(a int32, b int32) int32 {

	if a == b {

		return b

	}

	if a > b {

		return a

	}

	return b
}

func twoSum(nums [4]int, target int) []int {

	var array []int

	for index, _ := range nums {
		for i := 1; i < len(nums); i++ {

			if i == index {

				i++

			}

			if nums[i]+nums[index] == target {

				return append(array, index, i)

			}
		}
	}

	return array
}

func swap(px, py *int) {
	// tempx := *px
	// tempy := *py
	// *px = tempy
	// *py = tempx

	*px, *py = *py, *px
}

//sắp xếp nổi bọt
func sortArrayLowToHigh(arrayInput []int) []int {

	length := len(arrayInput)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {

			if arrayInput[i] > arrayInput[j] {

				swap(&arrayInput[i], &arrayInput[j])

			}
		}
	}

	return arrayInput
}

//chưa hoàn tất
func findMedianSortedArrays(nums1 [2]int, nums2 [1]int) float64 {

	var permutedList []int

	for _, value1 := range nums1 {

		permutedList = append(permutedList, value1)

	}

	for _, value2 := range nums2 {

		permutedList = append(permutedList, value2)

	}

	permutedList = sortArrayLowToHigh(permutedList)
	fmt.Println(permutedList)
	return 1
}

func buildArray(nums []int) []int {

	var ansNumber []int
	var arrayFinish []int

	for _, val := range nums {

		ansNumber = append(ansNumber, val)

	}

	for _, val := range ansNumber {

		arrayFinish = append(arrayFinish, nums[val])

	}

	return arrayFinish
}

func getConcatenation(nums []int) []int {

	var ansNumber []int

	for idx, _ := range nums {

		ansNumber = append(ansNumber, nums[idx])

	}

	ansNumber = append(ansNumber, nums...)
	return ansNumber
}

func finalValueAfterOperations(operations []string) int {

	X := 0

	for _, val := range operations {

		if val == "--X" {

			X = X - 1

		}

		if val == "X--" {

			X = X - 1

		}

		if val == "++X" {

			X = X + 1

		}

		if val == "X++" {

			X = X + 1

		}
	}

	return X
}

func runningSum(nums []int) []int {

	var arrayNumber []int
	temp := 0

	for _, val := range nums {

		temp = temp + val
		arrayNumber = append(arrayNumber, temp)

	}

	return arrayNumber
}

func shuffle(nums []int, n int) []int {

	var arrayNumberA []int
	var arrayNumberB []int
	var arrayNumberAB []int

	length := len(nums)

	for i := 0; i <= n; i++ {

		arrayNumberA = append(arrayNumberA, nums[i])

	}

	for i := n; i < length; i++ {

		arrayNumberB = append(arrayNumberB, nums[i])

	}

	for i := 0; i < length-n; i++ {

		arrayNumberAB = append(arrayNumberAB, arrayNumberA[i], arrayNumberB[i])

	}

	return arrayNumberAB
}

func maximumWealth(listCustomer [][]int) int {

	sumMoney := 0
	var listWealth []int
	var customerHasWealth int32
	var sortedList []int

	for _, customer := range listCustomer {
		for _, money := range customer {

			sumMoney = sumMoney + money

		}

		listWealth = append(listWealth, sumMoney)
		sumMoney = 0
	}

	length := len(listWealth)

	if length < 2 {

		return listWealth[0]

	}

	//Sort mang tu thâp đến cao
	sortedList = sortArrayLowToHigh(listWealth)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {

			if sortedList[i] == sortedList[j] {

				customerHasWealth = int32(sortedList[i])

			}

			if sortedList[i] < sortedList[j] {

				customerHasWealth = max(int32(sortedList[i]), int32(sortedList[j]))

			}

		}
	}

	return int(customerHasWealth)
}

//Approach(Greedy)
func kidsWithCandies(candies []int, extraCandies int) []bool {

	var arrExreaCandies []bool
	length := len(candies)
	maxCandies := 0

	for i := 0; i < length; i++ {

		if candies[i] > maxCandies {

			maxCandies = candies[i]

		}
	}

	for i := 0; i < length; i++ {
		// arrExreaCandies[i] = (candies[i]+extraCandies >= maxCandies)

		if candies[i]+extraCandies >= maxCandies {

			fmt.Println(candies[i])
			arrExreaCandies = append(arrExreaCandies, true)

		} else {

			fmt.Println("---", candies[i])
			arrExreaCandies = append(arrExreaCandies, false)

		}
	}

	return arrExreaCandies
}

func numIdenticalPairs(nums []int) int {

	length := len(nums)
	var countHaveNumberGood = 0

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {

			if nums[i] == nums[j] {

				countHaveNumberGood++

			}
		}
	}

	return countHaveNumberGood
}

func main() {
	// arr := [6][6]int32{
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 1, 0, 0, 0, 0},
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 0, 1, 2, 4, 0},
	// }

	// array1 := [2]int{5, 1}
	// array2 := [1]int{2}

	// fmt.Println(findMedianSortedArrays(array1, array2))

	// array1 := []int{5, 0, 1, 2, 3, 4}
	// fmt.Println(buildArray(array1))

	// array1 := []int{1, 2, 1}
	// fmt.Println(getConcatenation(array1))

	// operations := []string{"--X", "X++", "X++"}
	// fmt.Println(finalValueAfterOperations(operations))

	// array1 := []int{1, 2, 3, 4}
	// fmt.Println(runningSum(array1))

	// array1 := []int{2, 5, 1, 3, 4, 7}
	// fmt.Println(shuffle(array1, 3))

	// accounts := [][]int{
	// 	{2, 8, 7},
	// 	{7, 1, 3},
	// 	{1, 9, 5},
	// }

	// accounts := [][]int{
	// 	{6, 59, 64, 19, 30, 76, 71, 86, 90, 25, 56, 17, 19, 72, 61, 56, 24, 40, 35, 39, 67, 28, 52, 11, 82, 72, 8, 82, 81, 47},
	// }

	// accounts := [][]int{
	// 	{76, 6, 55, 11, 30, 65, 33, 74, 14, 16, 57, 79, 17, 87, 36, 61, 6},
	// 	{10, 18, 5, 55, 94, 28, 8, 36, 73, 62, 23, 62, 20, 70, 91, 7, 10},
	// 	{73, 52, 73, 50, 5, 13, 45, 52, 46, 66, 68, 31, 89, 38, 23, 70, 90},
	// 	{53, 46, 30, 37, 72, 8, 71, 51, 40, 29, 73, 96, 63, 87, 63, 73, 61},
	// 	{19, 42, 55, 82, 58, 24, 9, 39, 49, 4, 78, 87, 76, 22, 14, 5, 90},
	// 	{94, 61, 54, 11, 53, 13, 62, 36, 87, 51, 59, 71, 14, 42, 10, 36, 75},
	// 	{96, 96, 16, 75, 27, 16, 56, 48, 90, 38, 40, 7, 65, 20, 98, 54, 84},
	// 	{13, 2, 3, 62, 89, 28, 68, 3, 36, 85, 59, 71, 61, 32, 44, 65, 33},
	// 	{43, 29, 78, 49, 59, 18, 30, 39, 43, 43, 36, 33, 20, 20, 2, 23, 39},
	// 	{72, 52, 52, 7, 89, 78, 18, 1, 7, 20, 24, 88, 15, 36, 34, 25, 32},
	// 	{78, 91, 59, 10, 97, 39, 93, 76, 11, 22, 81, 59, 50, 100, 77, 61, 6},
	// 	{67, 24, 53, 76, 81, 45, 38, 90, 52, 5, 91, 63, 66, 37, 59, 100, 20},
	// 	{39, 56, 10, 80, 100, 89, 49, 75, 49, 4, 17, 7, 69, 48, 25, 53, 18},
	// 	{91, 78, 84, 98, 31, 47, 97, 79, 1, 33, 68, 52, 30, 1, 2, 95, 47},
	// 	{95, 62, 46, 59, 11, 57, 26, 2, 64, 27, 25, 45, 97, 43, 31, 40, 89},
	// 	{59, 35, 65, 46, 50, 44, 60, 23, 99, 80, 47, 57, 44, 89, 69, 99, 48},
	// 	{60, 72, 57, 2, 66, 25, 64, 37, 31, 78, 23, 75, 18, 19, 92, 91, 86},
	// 	{58, 42, 94, 94, 24, 17, 17, 34, 97, 90, 25, 100, 46, 44, 74, 29, 77},
	// 	{49, 62, 43, 14, 96, 38, 27, 93, 52, 92, 81, 62, 38, 89, 71, 30, 77},
	// 	{1, 97, 54, 83, 95, 63, 22, 37, 36, 96, 4, 95, 12, 70, 93, 84, 40},
	// 	{75, 32, 21, 19, 5, 51, 70, 69, 10, 34, 29, 23, 53, 85, 96, 71, 30},
	// 	{88, 79, 42, 51, 40, 24, 7, 38, 56, 26, 82, 100, 19, 16, 92, 75, 100},
	// 	{56, 14, 70, 31, 60, 45, 29, 71, 43, 8, 14, 1, 51, 63, 28, 99, 11},
	// 	{56, 63, 54, 56, 50, 34, 99, 99, 11, 97, 15, 20, 62, 76, 45, 98, 60},
	// 	{45, 37, 22, 71, 57, 40, 95, 52, 91, 78, 57, 99, 79, 45, 100, 79, 10},
	// 	{74, 96, 40, 94, 75, 12, 74, 73, 18, 51, 73, 62, 69, 50, 63, 47, 91},
	// 	{38, 60, 1, 74, 95, 28, 93, 100, 57, 66, 47, 55, 2, 89, 90, 80, 88},
	// 	{9, 13, 73, 47, 78, 79, 46, 82, 7, 16, 36, 15, 87, 66, 3, 59, 55},
	// 	{20, 66, 23, 18, 41, 13, 75, 70, 7, 3, 97, 74, 52, 60, 79, 79, 94},
	// 	{50, 4, 65, 88, 30, 92, 72, 39, 22, 34, 7, 61, 45, 58, 32, 14, 73},
	// 	{48, 61, 29, 70, 67, 21, 19, 72, 41, 14, 73, 26, 63, 73, 72, 13, 16},
	// 	{79, 14, 31, 84, 20, 67, 41, 74, 87, 53, 52, 6, 14, 77, 33, 39, 18},
	// 	{50, 70, 48, 65, 63, 20, 76, 80, 49, 76, 43, 41, 51, 33, 68, 78, 61},
	// 	{97, 62, 37, 13, 31, 84, 89, 92, 33, 35, 95, 9, 77, 22, 29, 57, 71},
	// 	{89, 16, 24, 84, 36, 59, 33, 100, 55, 48, 47, 45, 59, 11, 20, 67, 69},
	// 	{65, 58, 96, 19, 9, 84, 92, 55, 69, 61, 80, 59, 78, 76, 19, 29, 61},
	// 	{84, 66, 72, 59, 55, 71, 77, 23, 55, 63, 87, 57, 36, 72, 12, 58, 91},
	// }
	// fmt.Println(maximumWealth(kidsWithCandies))

	// candies := []int{2, 3, 5, 1, 3}
	// fmt.Println(kidsWithCandies(candies, 3))

	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))

}
