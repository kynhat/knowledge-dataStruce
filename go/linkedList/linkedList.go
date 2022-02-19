package main

import "fmt"

type node struct {
	data int
	next *node
}

// tạo thêm 1 Node dể làm bài tập khác
type ListNode struct {
	data int
	next *ListNode
}

type linkedList struct {
	head *node
	// tail   *node
	length int
}

//Them node vao dau linkedList
func (l *linkedList) prepend(n *node) {

	//tao 1 vùng nhớ mới  gắn node đầu vào
	second := l.head

	//gắn node mới vào vị trí vùng nhớ head
	l.head = n

	// trỏ với địa chỉ node củ
	l.head.next = second

	//tăng biến lên 1
	l.length++
}

//Xoa phần tử trong node
func (l *linkedList) deleteWithValue(value int) {

	if l.length == 0 {

		return

	}

	if l.head.data == value {

		l.head = l.head.next
		l.length--
		return

	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {

		//nếu như node kế của nut kế tiếp mà la null
		if previousToDelete.next.next == nil {

			return

		}

		//Thì gán nút kế = cái nút kế kế
		previousToDelete = previousToDelete.next
	}

	// Nếu trương tìm thấy kết quả cần xóa thì gắn bằng nút kế của kế tiếp luôn
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

// Xuất k ết quả ra màng hình
func (l linkedList) printListData() {

	toPrint := l.head

	for l.length != 0 {

		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--

	}

	fmt.Printf("\n")
}

// ---------------------------------------Bài tập ---------------------------------------

// bài 1: Add two number
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]

// Tổng 2 số nguyên
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head, sum := new(ListNode), 0

	for cur := head; l1 != nil || l2 != nil || sum != 0; {

		if l1 != nil {

			sum += l1.data
			l1 = l1.next

		}

		if l2 != nil {

			sum += l2.data
			l2 = l2.next

		}

		cur.next = &ListNode{data: sum % 10}

		// y nghia sum = sum / 10
		sum /= 10
		cur = cur.next
		fmt.Println(cur)

	}

	return head.next
}

// func main() {
// 	//Lý thuyết

// 	myList := linkedList{}
// 	node1 := &node{data: 48}
// 	node2 := &node{data: 49}
// 	node3 := &node{data: 50}
// 	node4 := &node{data: 51}
// 	node5 := &node{data: 52}

// 	myList.prepend(node1)
// 	myList.prepend(node2)
// 	myList.prepend(node3)
// 	myList.prepend(node4)
// 	myList.prepend(node5)
// 	// myList.deleteWithValue(48)
// 	myList.printListData()

// }
