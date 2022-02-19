
const merge = (array1, array2) => {
  let final = [];

  while(array1.length && array2.length) {
    final.push(array1[0] > array2[0] ? array2.shift() : array1.shift())
  }

  //if we still have values, let's add them at the end of `c`
  while(array1.length) {
    final.push(array1.shift())
  }

  while(array2.length) {
    final.push(array2.shift())
  }

  return final;
}

const _mergeSort = (array) => {
  let lenOfArray = array.length;
  let middle = Math.floor(array.length / 2);

  if (lenOfArray < 2) {
    return array;
  }

  const arr_l = array.slice(0, middle)
  const arr_r = array.slice(middle, array.length)

  const sorted_l = _mergeSort(arr_l);
  const sorted_r = _mergeSort(arr_r);

  return merge(sorted_l, sorted_r);
}

let arr = [2, 3, 14, 2, 13, 4, 8, 9];
console.log(_mergeSort(arr));

