class Queue {
  constructor() {
    this.dataStore = [];
    this.enqueue = this.enqueue;
    this.dequeue = this.dequeue;
    this.front = this.front;
    this.back = this.back;
    this.toString = this.toString;
    this.empty = this.empty;
  }

  // enqueue is add an element to the end of a queue
  enqueue(element) {
    this.dataStore.push(element);
  }

  //dequeue is removes an element  from the front of a queue
  dequeue() {
    return this.dataStore.shift();
  }

  // front amd back elements of a queue using these func
  front() {
    return this.dataStore[0];
  }

  back() {
    return this.dataStore[this.dataStore.length - 1];
  }

  //we need a function that lets us know if a queue is empty
  empty() {
    if (this.dataStore.length == 0) {
      return true;
    } else {
      return false;
    }
  }
}

var q = new Queue();
q.enqueue("Meredith");
q.enqueue("Cynthia");
q.enqueue("Jennifer");

console.log(q.dataStore);

q.dequeue();
console.log(q.dataStore);
console.log("Front of queue: " + q.front());
console.log("Back of queue: " + q.back());