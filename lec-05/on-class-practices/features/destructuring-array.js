const arr = [1, 2, 3, 4, 5, 6];

// First item
const [num1] = arr;
console.log('num', num1);

// Ignore item
const [numA, , numC] = arr;
console.log('numA', numA);
console.log('numC', numC);

// Assignment
const [ numa, numb, ...numall ] = arr;
console.log('numa', numa);
console.log('numb', numb);
console.log('numall', numall);

// Default value
const [defaultA = 5, defaultB = 7] = [1];
console.log('defaultA', defaultA);
console.log('defaultB', defaultB);
