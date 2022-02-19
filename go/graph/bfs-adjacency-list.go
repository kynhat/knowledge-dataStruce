// package main

// import "fmt"

// /*
//    Go Program for
//    Breadth first traversal in directed graph
// */
// type AjlistNode struct {
// 	// Vertices node key
// 	id   int
// 	next *AjlistNode
// }

// func getAjlistNode(id int) *AjlistNode {

// 	return &AjlistNode{id, nil}
// }

// // Graph Vertices
// type Vertices struct {
// 	data int
// 	next *AjlistNode
// 	last *AjlistNode
// }

// func getVertices(data int) *Vertices {
// 	return &Vertices{data, nil, nil}
// }

// // Create class of queue
// type QNode struct {
// 	data int
// 	next *QNode
// }

// func getQNode(value int) *QNode {
// 	return &QNode{value, nil}
// }

// type MyQueue struct {
// 	head  *QNode
// 	tail  *QNode
// 	count int
// }

// func getMyQueue() *MyQueue {

// 	return &MyQueue{nil, nil, 0}
// }
// func (this MyQueue) size() int {
// 	return this.count
// }
// func (this MyQueue) isEmpty() bool {
// 	return this.count == 0
// }

// // Add new node of queue
// func (this *MyQueue) enqueue(value int) {
// 	// Create dynamic node
// 	var node *QNode = getQNode(value)
// 	if this.head == nil {
// 		// Add first element into queue
// 		this.head = node
// 	} else {
// 		// Add node at the end using tail
// 		this.tail.next = node
// 	}
// 	this.count++
// 	this.tail = node
// }

// // Delete a element into queue
// func (this *MyQueue) dequeue() int {
// 	if this.head == nil {
// 		fmt.Println("Empty Queue")
// 		return -1
// 	}
// 	// Pointer variable which are storing
// 	// the address of deleted node
// 	var temp *QNode = this.head
// 	// Visit next node
// 	this.head = this.head.next
// 	this.count--
// 	if this.head == nil {
// 		// When deleting a last node
// 		this.tail = nil
// 	}
// 	return temp.data
// }

// // Get front node
// func (this MyQueue) peek() int {
// 	if this.head == nil {
// 		fmt.Println("Empty Queue")
// 		return -1
// 	}
// 	return this.head.data
// }

// type Graph struct {
// 	// Number of Vertices
// 	size int
// 	node []*Vertices
// }

// func getGraph(size int) *Graph {
// 	var me *Graph = &Graph{size, make([]*Vertices, size)}
// 	me.setData()
// 	return me
// }

// // Set initial node value
// func (this *Graph) setData() {
// 	if this.size <= 0 {
// 		fmt.Println("\nEmpty Graph")
// 	} else {
// 		for index := 0; index < this.size; index++ {
// 			// Set initial node value
// 			this.node[index] = getVertices(index)
// 		}
// 	}
// }

// //  Handling the request of adding new edge
// func (this *Graph) addEdge(start, last int) {
// 	if start >= 0 && start < this.size &&
// 		last >= 0 && last < this.size {
// 		// Safe connection
// 		var edge *AjlistNode = getAjlistNode(last)
// 		if this.node[start].next == nil {
// 			this.node[start].next = edge
// 		} else {
// 			// Add edge at the end
// 			this.node[start].last.next = edge
// 		}
// 		// Get last edge
// 		this.node[start].last = edge
// 	} else {
// 		// When invalid nodes
// 		fmt.Println("\nHere Something Wrong")
// 	}
// }
// func (this Graph) printGraph() {
// 	if this.size > 0 {
// 		// Print graph ajlist Node value
// 		for index := 0; index < this.size; index++ {
// 			fmt.Print("\nAdjacency list of vertex ",
// 				index, " : ")
// 			var temp *AjlistNode = this.node[index].next
// 			for temp != nil {
// 				// Display graph node
// 				fmt.Print("  ", this.node[temp.id].data)
// 				// Visit to next edge
// 				temp = temp.next
// 			}
// 		}
// 	}
// }

// // Breadth first traversal for a directed graph node
// func (this Graph) bfs(point int) {
// 	if point > this.size || this.size <= 0 {
// 		return
// 	}
// 	var q *MyQueue = getMyQueue()
// 	var temp *AjlistNode = nil
// 	// Set false to each element
// 	var visit = make([]bool, this.size)
// 	//Add first element into queue
// 	q.enqueue(point)
// 	fmt.Print("\n BFS :")
// 	for q.isEmpty() == false {
// 		temp = this.node[q.peek()].next
// 		for temp != nil {
// 			if !visit[temp.id] {
// 				visit[temp.id] = true
// 				q.enqueue(temp.id)
// 			}
// 			// visit to next edge
// 			temp = temp.next
// 		}
// 		visit[q.peek()] = true
// 		fmt.Print("  ", q.peek())
// 		q.dequeue()
// 	}
// }
// func main() {
// 	// 6 implies the number of nodes in graph
// 	var g *Graph = getGraph(6)
// 	// Connect node with an edge
// 	g.addEdge(0, 1)
// 	g.addEdge(0, 5)
// 	g.addEdge(1, 1)
// 	g.addEdge(2, 1)
// 	g.addEdge(3, 0)
// 	g.addEdge(3, 3)
// 	g.addEdge(4, 2)
// 	g.addEdge(4, 3)
// 	g.addEdge(5, 1)
// 	// print graph element
// 	g.printGraph()
// 	g.bfs(4)
// }
