class HashTable {
  constructor(size = 7) {
    this.dataMap = new Array(size);
  }

  _hash(key) {
    let hash = 0;

    for (let i = 0; i < key.length; i++) {
      //  hash += key.charCodeAt(i);
      hash = (hash + key.charCodeAt(i) * 23) % this.dataMap.length;
    }

    return hash;
  }

  set(key, value) {
    let index = this._hash(key);

    if (!this.dataMap[index]) {
      // tao mang rong de luu key va value
      this.dataMap[index] = [];
    }

    this.dataMap[index].push([key, value]);

    return this;
  }

  get(key) {
    let index = this._hash(key);

    if (this.dataMap[index]) {
      for (let i = 0; i < this.dataMap[index].length; i++) {
        // key
        if (this.dataMap[index][i][0] === key) {
          // value
          return this.dataMap[index][i][1];
        }
      }
    }

    return undefined;
  }

  key() {
    let allKeys = [];
    for (let i = 0; i < this.dataMap.length; i++) {
      // Nếu rỗng thì nhảy sang phần tử kế tiếp
      if (this.dataMap[i]) {
        for (let j = 0; j < this.dataMap[i].length; j++) {
          allKeys.push(this.dataMap[j][j][0]);
        }
      }
    }

    return allKeys;
  }
}

let tesst = new HashTable();
tesst.set("lop1", "an");
tesst.set("lop1", "an11");
tesst.set("lop2", "ggan");
tesst.set("lop2", "ggan");

console.log(tesst);

console.log(tesst.get("lop1"));
console.log(tesst.get("lop1"));
