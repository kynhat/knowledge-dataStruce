class Node {
  constructor(data, left, right) {
    this.data = data
    this.left = left
    this.right = right
  }

  show() {
    return this.data
  }

}

class BTS {
  constructor() {
    this.root = null
    // this.insert = this.insert
    // this.inOrder = this.inOrder
  }

insert(data) {
  node = new Node(data)

  if(this.root == null) {
    this.root = Node
  } else {

    var current = this.root
    var parent

    while(true) {
      parent = current

      if(data < current.data) {
        current = current.left

        if(current == null) {
          parent.left = node
          break
        }
      }

      if(data > current.data) {
        current = current.right

        if(current == null) {
          parent.right = node
          break
        }
      }
    }
  }
}

//inOrderTraversal
//postOrderTraversal
//preOrderTraversal

}
