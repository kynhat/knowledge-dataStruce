package main

import (
	"fmt"
	"io"
)

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	data  int
}

type binaryTree struct {
	root *binaryNode
}

func (n *binaryNode) insert(data int) {

	if n == nil {

		return

	}

	if data <= n.data {

		if n.left == nil {

			n.left = &binaryNode{data: data, left: nil, right: nil}

		} else {

			n.left.insert(data)

		}
	}

	if data > n.data {

		if n.right == nil {

			n.right = &binaryNode{data: data, left: nil, right: nil}

		} else {

			n.right.insert(data)

		}
	}
}

// Dùng vong lập
func (t *binaryTree) insertByLoop(data int) *binaryTree {

	if t.root == nil {

		t.root = &binaryNode{data: data, left: nil, right: nil}

	} else {

		temp := t.root

		for true {

			if data > temp.data {

				if temp.right == nil {

					t.root = &binaryNode{data: data, left: nil, right: nil}

				} else {

					temp = temp.right

				}
			} else {

				if temp.left == nil {

					t.root = &binaryNode{data: data, left: nil, right: nil}

				} else {

					temp = temp.left

				}
			}
		}
	}

	return t
}

func (t *binaryTree) InsertTree(data int) *binaryTree {

	if t.root == nil {

		t.root = &binaryNode{data: data, left: nil, right: nil}

	} else {

		t.root.insert(data)

	}

	return t
}

//----------------- Có 6 cách duyệt cây -----------------
func NLR(root *binaryNode) {

	if root != nil {

		if root != nil {

			fmt.Println(root.data)

		}

		NLR(root.left)
		NLR(root.right)
	}
}

// func NRL(root *binaryNode) {

// 	if root != nil {

// 		if root != nil {
// 			fmt.Println(root.data)
// 		}

// 		NRL(root.right)
// 		NRL(root.left)
// 	}
// }

func LNR(root *binaryNode) {

	if root != nil {
		LNR(root.left)

		if root != nil {

			fmt.Println(root.data)

		}

		LNR(root.right)
	}
}

func LRN(root *binaryNode) {

	if root != nil {

		LRN(root.left)
		LRN(root.right)

		if root != nil {

			fmt.Println(root.data)

		}
	}
}

// func RNL(root *binaryNode) {

// 	if root != nil {

// RNL(root.right)

// 		if root != nil {

// 			fmt.Println(root.data)

// 		}

// 		RNL(root.right)
// 	}
// }

// func RLN(root *binaryNode) {

// 	if root != nil {

// 		RLN(root.right)
// 		RLN(root.left)

// 		if root != nil {
// 			fmt.Println(root.data)
// 		}
// 	}
// }

func print(w io.Writer, node *binaryNode, ns int, ch rune) {
	if node == nil {

		return

	}

	for i := 0; i < ns; i++ {

		fmt.Fprint(w, " ")

	}

	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main33() {
	tree := &binaryTree{}
	tree.
		InsertTree(10).
		InsertTree(5).
		InsertTree(15).
		InsertTree(3).
		InsertTree(7).
		InsertTree(18)

	// fmt.Println(tree.root)
	// fmt.Println(tree.root.left)
	// fmt.Println(tree.root.right)

	//Duyệt tiền thứ tự trong cây nhị phân

	NLR(tree.root)
	fmt.Println("---------------")

	//Duyệt trung thứ tự trong cây nhị phân
	//Nếu một cây nhị phân được duyệt trung thứ tự, kết quả tạo ra sẽ là các giá trị khóa được sắp xếp theo thứ tự tăng dần.
	LNR(tree.root)
	fmt.Println("---------------")

	// Duyệt hậu thứ tự trong cây nhị phân
	LRN(tree.root)
	// print(os.Stdout, tree.root, 0, 'M')
}
