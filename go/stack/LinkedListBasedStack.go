package main

type Node struct {
	data int
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push(): Thêm 1 phần tử vào Stack. Hàm này tương tự với hàm thêm 1 phần tử vào đầu danh sách liên kết đơn, dĩ nhiên các bạn cũng có thể xây dựng giống với hàm thêm phần tử vào cuối danh sách liên kết đơn. Tuy nhiên, khi xây dựng như vậy thì hàm Pop(), tương ứng với hàm Push() ấy cũng phải là hàm xóa phần tử ở cuối danh sách, và việc này yêu cầu chi phí là O(n), trong khi hàm xóa 1 phần tử ở đầu danh sách(ứng với hàm xây dựng bên dưới) chỉ là O(1)

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {

	// s.nodes = append(s.nodes, n)
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {

	if s.count == 0 {
		return nil
	}

	s.count--
	return s.nodes[s.count]
}

func (s *Stack) rangeSumBST(root *Node, low, high int) int {

	ans := 0
	s.Push(root)

	for s.IsEmpty() {

		node := s.Pop()

		if root != nil {

			if node.data >= low && node.data <= high {

				ans += node.data

			}

			// if node.data > low {

			// 	s.Push(node.left)

			// }

			// if node.data < high {

			// 	s.Push(node.right)

			// }
		}
	}

	return ans
}

func (s Stack) IsEmpty() bool {

	return s.count == 0
}

func main() {
	s := NewStack()
	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})

	s.rangeSumBST(&Node{1}, 1, 2)
	// fmt.Println(s.nodes)
	// fmt.Println(s.Pop(), s.Pop(), s.Pop())

	// bai tap

}
