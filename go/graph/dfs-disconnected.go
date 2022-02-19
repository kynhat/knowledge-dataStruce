// package main

// import "fmt"

// /*
//    Go program for
//    DFS of disconnected graph
// */

// type AjlistNode struct {
// 	// Vertices node key
// 	id   int
// 	next *AjlistNode
// }

// func getAjlistNode(id int) *AjlistNode {
// 	// return new AjlistNode
// 	return &AjlistNode{
// 		id,
// 		nil,
// 	}
// }

// type Vertices struct {
// 	data int
// 	next *AjlistNode
// 	last *AjlistNode
// }

// func getVertices(data int) *Vertices {
// 	// return new Vertices
// 	return &Vertices{
// 		data,
// 		nil,
// 		nil,
// 	}
// }

// type Graph struct {
// 	// Number of Vertices
// 	size int
// 	node []*Vertices
// }

// func getGraph(size int) *Graph {
// 	// return new Graph
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
// 	if start >= 0 && start < this.size && last >= 0 && last < this.size {
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
// 			fmt.Print("\nAdjacency list of vertex ", index, " :")
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

// // Pirnt Path in dfs of given vertex
// func (this Graph) dfs(point int, path []bool) {
// 	if path[point] {
// 		// When alredy visit
// 		return
// 	} else {
// 		// Display result node
// 		fmt.Print("  ", point)
// 		// Collect the current node into the path
// 		path[point] = true
// 		var temp *AjlistNode = this.node[point].next
// 		for temp != nil {
// 			// Visit to edge node
// 			this.dfs(temp.id, path)
// 			// Visit to next edge
// 			temp = temp.next
// 		}
// 	}
// }
// func (this Graph) dfsImp() {
// 	if this.size <= 0 {
// 		// No nodes
// 		return
// 	}
// 	var path = make([]bool, this.size)
// 	var n int = 0
// 	fmt.Print("\nResult : ")
// 	for n < this.size {
// 		if path[n] != true {
// 			// When node is not visit
// 			this.dfs(n, path)
// 		}
// 		// next node
// 		n++
// 	}
// }
// func main() {
// 	// 7 implies the number of nodes in graph
// 	var g *Graph = getGraph(7)
// 	// Connect node with an edge
// 	g.addEdge(1, 0)
// 	g.addEdge(1, 4)
// 	g.addEdge(1, 6)
// 	g.addEdge(2, 0)
// 	g.addEdge(2, 5)
// 	g.addEdge(3, 5)
// 	g.addEdge(6, 1)
// 	// Print graph element
// 	g.printGraph()
// 	// Test
// 	g.dfsImp()
// }
