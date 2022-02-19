class Node {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

class Queue {
  constructor(value) {
    const newNode = new Node(value);
    this.frist = newNode;
    this.last = newNode;
    this.length = 1;
  }

  // thêm đối tượng o vào cuối hàng đợi.
  enqueue(value) {
    const newNode = new Node(value);

    if (this.length === 0) {
      this.frist = newNode;
      this.last = newNode;
    } else {
      this.last.next = newNode;
      this.last = newNode;
    }

    this.length++;
    return this;
  }

  // lấy đối tượng ở đầu queue ra khỏi hàng đợi và trả về giá trị của nó.
  dequeue(value) {
    if (this.length === 0) return undefined;
    let temp = this.frist
    
    if(this.length === 1) {
      this.frist = null
      this.last = null
    }else {
      this.frist = this.frist.next
      temp.next = null
    }

    this.length--
    return temp
  }
}
