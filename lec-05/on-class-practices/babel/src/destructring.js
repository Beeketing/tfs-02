const arr = [1, 2, 3];
const [a, b] = arr;

console.log(a, b);

const [a1, , b1] = arr;

console.log(a1, b1);

const a2 = arr[0];
const b2 = arr[2];


const arr2 = [1];
const [a3, c3 = 20] = arr2;
console.log('a3', a3);
console.log('c3', c3);

const obj = { name: 'OCG', age: 123 };
console.log(obj.name);
const { name, age } = obj;
console.log(name, age);

const obj1 = obj;
obj1.floor = 7;
console.log(obj1);
console.log(obj);
