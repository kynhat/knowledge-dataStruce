//Giải thuật mẫu cho tìm kiếm tuyến tính

func linearSearch(array []int, value int) int {

	for i = 0; i < len(array); i++ {

		if array[i] == x {

			return i

		}
	}

	return -1
}

func linearSearchMin(array []int, value int) int {

	for i = 1; i < len(array); i++ {

		int numberMin = 0

		if array[numberMin] > array[i] {

			numberMin = i

		}
	}

	return numberMin
}