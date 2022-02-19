package main

import "fmt"

// Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high]

type Node struct {
	left  *Node
	right *Node
	data  int
}

type BinaryTree struct {
	root *Node
}

func (n *Node) insert(data int) {

	if n == nil {

		return

	}

	if data <= n.data {

		if n.left == nil {

			n.left = &Node{data: data, left: nil, right: nil}

		} else {

			n.left.insert(data)

		}
	}

	if data > n.data {

		if n.right == nil {

			n.right = &Node{data: data, left: nil, right: nil}

		} else {

			n.right.insert(data)

		}
	}
}

func (t *BinaryTree) InsertTrees(data int) *BinaryTree {

	if t.root == nil {

		t.root = &Node{data: data, left: nil, right: nil}

	} else {

		t.root.insert(data)

	}

	return t
}

var arr []int

func rangeSumBST(root *Node, low int, high int) int {

	sum := 0
	dfs(root, low, high)

	for _, v := range arr {
		sum = sum + v
	}

	fmt.Println(sum)

	return sum
}

func dfs(node *Node, L, R int) {

	var ans = 0

	if node != nil {

		if node.data >= L && node.data <= R {

			ans += node.data
			arr = append(arr, ans)
			fmt.Println("ans", ans)

		}

		if node.data > L {

			dfs(node.left, L, R)

		}

		if node.data < R {

			dfs(node.right, L, R)

		}
	}
}

func main() {

	tree := &BinaryTree{}
	tree.
		InsertTrees(10).
		InsertTrees(5).
		InsertTrees(15).
		InsertTrees(3).
		InsertTrees(7).
		InsertTrees(13).
		InsertTrees(18).
		InsertTrees(1).
		InsertTrees(6)

	rangeSumBST(tree.root, 6, 10)
}
