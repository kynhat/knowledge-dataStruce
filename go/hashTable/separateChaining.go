package main

import (
	"fmt"
)

//bucketNode struct
type bucketNode struct {
	key  string
	next *bucketNode
}

//bucket is a linkerList in each slot of the
type bucket struct {
	head *bucketNode
}

//ArraySize is the size of the hash table array
const ArraySize = 7

//Hash Table will hold an array
type HashTable struct {
	// Ex var a [3]int //int array with length 3
	//Khai báo kiểu dư liệu là bucket để khi sự dụng biến array thì có thể khởi được
	array [ArraySize]*bucket
}

//Hàm insert node của linkerList
//Insret  will take in a key, create a node with key and insert tyje node in the bucket
func (b *bucket) insert(k string) {

	// b.head hiện tại bằng nil
	// if b.head != nil {
	// 	newNode.next = b.head
	// 	b.head = newNode
	// 	b.length++
	// } else {
	// 	b.head = newNode
	// 	b.length++
	// }

	if !b.search(k) {

		//Khơi tạo node
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode

	} else {

		fmt.Println(k, "already exists")

	}
}

//Search will take in a key and return true and if the bucket has that key
func (b *bucket) search(k string) bool {

	currentNode := b.head

	for currentNode != nil {

		if currentNode.key == k {

			return true

		}

		currentNode = currentNode.next
	}

	return false
}

//Delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {

		b.head = b.head.next
		return

	}

	previousNode := b.head
	for previousNode.next != nil {

		if previousNode.next.key == k {

			previousNode.next = previousNode.next.next

		}

		previousNode = previousNode.next
	}
}

//Hàm Hash
func hash(key string) int {

	sum := 0
	// Giá trị đầu tiên là chỉ số (vị trí) của phần tử
	//Giá trị thứ hai là bản sao của phần tử đó (cùng giá trị)

	for _, v := range key {

		sum += int(v)

	}

	return sum % ArraySize
}

//Insert will take in a key add it to the hash table array
func (h *HashTable) Insert(key string) {

	//Băm key
	index := hash(key)
	h.array[index].insert(key)
}

//Search will take in a key and return true if that key is store in the hash table
func (h *HashTable) Search(key string) bool {

	index := hash(key)
	return h.array[index].search(key)
}

//Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {

	index := hash(key)
	h.array[index].delete(key)
}

func Init() *HashTable {

	//Kết quả khi khơi tạo 1 array
	//&{[<nil> <nil> <nil> <nil> <nil> <nil> <nil>]}
	result := &HashTable{}

	for i := range result.array {

		//Khởi tạo node vào mảng
		//Mỗi phần tử sẽ khơi tạo 1 node
		result.array[i] = &bucket{}

	}

	//&{[0xc00000e028 0xc00000e030 0xc00000e038 0xc00000e040 0xc00000e048 0xc00000e050 0xc00000e058]}
	return result
}

func main() {
	hashTable := Init()

	// nums := []string{"1", "2", "3", "1", "1", "2", "2", "2", "3"}

	// &{1 <nil>} 0xc00009e220
	// &{2 <nil>} 0xc00009e230
	// &{3 <nil>} 0xc00009e240
	// &{1 0xc0000ae018} 0xc00009e220
	// &{1 0xc0000ae0a8} 0xc00009e220
	// &{2 0xc0000ae048} 0xc00009e230
	// &{2 0xc0000ae108} 0xc00009e230
	// &{2 0xc0000ae138} 0xc00009e230
	// &{3 0xc0000ae078} 0xc00009e240

	list := []string{
		"KY",
		"Vun",
		"bac",
		"bac",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	// 	hashTable.Delete("KY")

	fmt.Println("KY", hashTable.Search("KY"))
	// fmt.Println(list)
}
