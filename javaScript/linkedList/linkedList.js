class Node {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

class LinkedList {
  constructor(value) {
    const newNode = new Node(value);
    this.head = newNode;
    this.tail = this.head;
    this.length = 1;
  }

  push(value) {
    const newNode = new Node(value);

    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      this.tail.next = newNode;
      this.tail = newNode;
    }

    this.length++;
    return this;
  }

  pop() {
    if (this.length === 0) {
      this.head = null;
      this.tail = null;
    }

    if (!this.head) {
      return undefined;
    }

    let prep = this.head;
    let temp = this.head;

    while (temp.next) {
      prep = temp;
      temp = temp.next;
      console.log("1", prep);
    }

    console.log("2");
    this.tail = prep;
    this.tail.next = null;
    this.length--;

    return this;
  }

  // insert node in head
  unshift(value) {
    const newNode = new Node(value);

    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      newNode.next = this.head;
      this.head = newNode;
    }

    this.length++;
    return this;
  }

  // delete node in head
  shift() {
    if (!this.head) {
      return undefined;
    }

    if (this.length === 0) {
      this.tail = null;
    }

    let temp = this.head;
    this.head = this.head.next;
    temp.next = null;
    this.length--;

    return this;
  }

  get(index) {
    if (index < 0 || index >= this.length) {
      return undefined;
    }

    let temp = this.head;
    for (let i = 0; i < index; i++) {
      temp = temp.next;
    }

    return temp;
  }

  set(index) {
    let temp = this.get(index);
    if (temp) {
      temp.value = value;
      return true;
    }

    return false;
  }

  insert(index, value) {
    if (index < 0 || index > this.length) return false;
    if (index === 0) return this.unshift(value);
    if (index === this.length) return this.push(value);

    const newNode = new Node(value);
    const temp = this.get(index - 1);

    temp.next = newNode;
    newNode.next = temp.next;
    this.length++;
    return true;
  }

  delete(index) {
    if (index < 0 || index > this.length) return undefined;
    if (index === 0) return this.shift();
    if (index === this.length - 1) return this.pop();

    const before = this.get(index - 1);
    const temp = before.next;

    before.next = temp.next;
    temp.next = null;
    this.length--;
    return temp;
  }

  reverse() {
    let temp = this.head;
    this.head = this.tail;
    this.tail = temp;

    let next = temp.next;
    let prev = null;
    for (let i = 0; i < this.length; i++) {
      next = temp.next;
      temp.next = prev;
      prev = temp;
      temp = next;
    }

    return this;
  }
}

myLinkedList = new LinkedList(5);
myLinkedList.push(6);
myLinkedList.push(9);
myLinkedList.push(7);
myLinkedList.pop();
// console.log(myLinkedList.push(1))
