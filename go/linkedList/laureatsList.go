package main

import (
	"fmt"
)

//Laureate: người đoạt giải
type Laureates struct {
	next  *Laureates
	name  string
	field string
	year  int
}

type LaureatsList struct {
	head *Laureates
	name string
}

//Khởi tạo danh sách ngươi đoat giải
func createLaureatesList(n string) *LaureatsList {

	return &LaureatsList{

		name: n,
	}
}

//Thêm người đoạt giải vào cuối mang
func (l_list *LaureatsList) addLaureat(n string, f string, y int) error {

	fmt.Printf("add %s %s %d \n", n, f, y)
	//Khởi tạo 1 nốt mới có dữ liệu mới
	l := &Laureates{
		name:  n,
		field: f,
		year:  y,
	}

	if l_list.head == nil {

		l_list.head = l

	} else {

		//copy nude dầu tiên vào 1 vung nhớ khác
		curentNode := l_list.head

		//Duyệt node từ node kế tiếp đế cuối
		for curentNode.next != nil {

			curentNode = curentNode.next

		}

		curentNode.next = l
	}

	return nil
}

func (l_list *LaureatsList) deleteLaureatsList(n string, f string, y int) error {

	fmt.Printf("delete %s %s %d \n", n, f, y)
	curentNode := l_list.head

	if curentNode == nil {

		fmt.Println("Empty list")
		return nil

	}

	// Head
	if curentNode.name == n &&
		curentNode.field == f &&
		curentNode.year == y {

		if curentNode == l_list.head {

			l_list.head = curentNode.next

		}

		return nil
	}

	//orthers
	fmt.Printf("*currentNode %+v\n", *curentNode)

	for curentNode.next != nil {

		fmt.Printf("*currentNode.next %+v\n", *&curentNode.next)
		next := curentNode.next

		if next.name == n &&
			next.field == f &&
			next.year == y {

			fmt.Printf("matching laureat found in next")
			curentNode.next = next.next
			break

		}

		curentNode = curentNode.next
	}

	return nil
}

func (l_list *LaureatsList) showAllLaureatsList() error {

	fmt.Printf("\n Current list:\n")
	fmt.Println("============================================")
	curentNode := l_list.head

	if curentNode == nil {

		fmt.Printf("empty List")
		return nil

	}

	fmt.Printf("%+v\n", *curentNode)

	for curentNode.next != nil {

		curentNode = curentNode.next
		fmt.Printf("%+v\n", *curentNode)

	}

	fmt.Println("============================================")
	return nil
}

func rectangle(l int, b int) (area int, parameter int) {

	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

// func main() {
// 	myLaureatsList := createLaureatesList("myLaureatsList")
// 	myLaureatsList.addLaureat("dasds", "dasdas", 3)
// 	myLaureatsList.addLaureat("eqqq", "eee", 4)
// 	myLaureatsList.addLaureat("fffff", "fffff", 5)
// 	myLaureatsList.showAllLaureatsList()

// 	myLaureatsList.deleteLaureatsList("fffff", "fffff", 5)
// 	myLaureatsList.showAllLaureatsList()

// 	strVar := "100"

// 	intVar, err := strconv.ParseInt(strVar, 0, 8)
// 	fmt.Println(intVar, err, reflect.TypeOf(intVar))

// 	//Golang Functions Returning Multiple Values
// 	// var a, p int
// 	// a, p = rectangle(20, 30)
// 	// fmt.Println("Area:", a)
// 	// fmt.Println("Parameter:", p)

// 	a := [8]int{5, 4, 3, 3, 1, 7, 8, 4} // short hand declaration to create array

// 	for i, _ := range a {
// 		fmt.Println("dddddd:", a[i]%8)
// 	}
// }
