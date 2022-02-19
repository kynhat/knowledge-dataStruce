class Stack {
  constructor() {
    this.dataStore = [];
    this.top = 0;
    this.push = this.push;
    this.pop = this.pop;
    this.peek = this.peek;
    this.clear = this.clear;
    this.length = this.length;
  }

  push(element) {
    this.dataStore[this.top++] = element;
  }

  pop() {
    return this.dataStore[this.top--];
  }

  length() {
    return this.top;
  }

  peek() {
    return this.dataStore[this.top - 1]
  }

  clear() {
    this.top = 0;
  }
}

var s = new Stack();
s.push("David");
s.push("Raymond");
s.push("Bryan");

console.log("length: " + s.length());
console.log("peek:", s.peek());

var popped = s.pop();
console.log("The popped element is: " + popped);
console.log("peek:", s.peek());

s.push("Cynthia");
console.log("peek:", s.peek());
s.clear();

console.log("length: " + s.length());
console.log("peek:", s.peek());
s.push("Clayton");
