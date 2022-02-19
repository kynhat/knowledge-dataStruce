class Node {
  constructor(value) {
    this.value = value
    this.left = null
    this.right = null
  }
}

class BST {
  constructor() {
    this.root = null
  }

  insert(value){
    const newNode = new Node(value)
    if(this.root === null){
      this.root = newNode
      return this
    }

    let temp = this.root
    while(true) {

      if(newNode.value === temp.value) return undefined

      if(newNode.value < temp.value) {
        if(temp.left === null) {
          temp.left = newNode
          return this
        }

        temp = temp.left
      }

      if(newNode.value > temp.value) {
        if(temp.right === null) {
          temp.right = newNode
          return this
        }

        temp = temp.right
      }
    }
  }

  contant(value) {

    if(this.root === null) return false
    let temp = this.root

    while(temp) {

      if(value < temp.value) {
        temp = temp.left
      }

      if(value > temp.value) {
        temp = temp.right
      } else {
        return true
      }
    }

    return false
  }

  minValueNode(currentNode) {
    while(currentNode.left != null) {
      currentNode = currentNode.left
    }

    return currentNode
  }
}

let mytree = new BST()
mytree.insert(1)
mytree.insert(2)
mytree.insert(2)
mytree.insert(4)

console.log(mytree.contant(1))
console.log(mytree.contant(33))