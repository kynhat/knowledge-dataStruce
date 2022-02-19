// package main

// import "fmt"

// // Vertex represents a graph vertex
// type Vertex struct {
// 	visited  bool
// 	key      int
// 	adjacent []*Vertex
// }

// // Graph represents an adjacency list graph
// type Graph struct {
// 	vertices []*Vertex
// }

// //add Vertex adds a Vertex to the Graph
// func (g *Graph) addVertex(k int) {

// 	if contains(g.vertices, k) {

// 		err := fmt.Errorf("Vertex %v not added because it is an exsting key", k)
// 		fmt.Println(err.Error())

// 	} else {

// 		g.vertices = append(g.vertices, &Vertex{key: k, visited: false})

// 	}
// }

// func (g *Graph) addEdge(from, to int) {

// 	// Get vertex
// 	fromVertex := g.getVertex(from)
// 	toVertex := g.getVertex(to)

// 	//check error
// 	if fromVertex == nil || toVertex == nil {

// 		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
// 		fmt.Println(err.Error())

// 	} else if contains(fromVertex.adjacent, to) {

// 		err := fmt.Errorf("Existing edge (%v --> %v)", from, to)
// 		fmt.Println(err.Error())

// 	} else {

// 		// Add edge
// 		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

// 	}
// }

// // GetVertex returns a pointer to the Vertex with a key interge
// func (g *Graph) getVertex(k int) *Vertex {

// 	for i, v := range g.vertices {

// 		if v.key == k {

// 			return g.vertices[i]

// 		}
// 	}

// 	return nil
// }

// // Contains it is func which check vertex it had existed
// func contains(s []*Vertex, k int) bool {

// 	for _, v := range s {

// 		if k == v.key {

// 			return true

// 		}
// 	}

// 	return false
// }

// // Print will print the adjacent list for each vertex of the graph
// func (g *Graph) Print() {

// 	for _, v := range g.vertices {

// 		fmt.Printf("\n Vertex %v :", v.key)

// 		for _, v := range v.adjacent {

// 			fmt.Printf(" Vertex adjacent %v", v.key)

// 		}
// 	}

// 	fmt.Println()
// }

// func (g *Graph) dfs(vertex *Vertex) {

// 	if !vertex.visited {

// 		vertex.visited = true
// 		fmt.Println("tesst", vertex)

// 		if len(vertex.adjacent) != 0 {

// 			for _, v := range vertex.adjacent {
// 				g.dfs(v)
// 			}

// 		}

// 	} else {

// 		return

// 	}
// }

// // n = so lượng node
// // g = list (graph)
// // visited = [false, false, false, ... ]
// func main() {

// 	test := &Graph{}

// 	for i := 0; i < 5; i++ {

// 		test.addVertex(i)

// 	}

// 	test.addEdge(0, 1)
// 	test.addEdge(0, 2)
// 	test.addEdge(0, 3)

// 	test.addEdge(2, 1)
// 	test.addEdge(2, 2)
// 	test.addEdge(2, 3)

// 	test.addEdge(3, 1)
// 	test.addEdge(3, 2)
// 	test.addEdge(3, 3)

// 	test.dfs(test.getVertex(0))
// 	test.Print()
// }
