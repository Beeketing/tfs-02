// Basic
function sum(x, y, z) {
  return x + y + z;
}

const numbers = [1, 2, 3];
console.log(sum(...numbers));

// Concat
let arr1 = [0, 1, 2];
let arr2 = [3, 4, 5];
arr1 = [...arr1, ...arr2];
console.log('concat', arr1);

// Spread in object literals
let obj1 = { foo: 'bar', x: 42 };
let obj2 = { foo: 'baz', y: 13 };
let clonedObj = { ...obj1 };
let mergedObj = { ...obj1, ...obj2 };
console.log('clonedObj', clonedObj);
console.log('mergedObj', mergedObj);
