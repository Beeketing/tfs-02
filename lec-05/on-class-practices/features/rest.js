function sum(...theArgs) {
  return theArgs.reduce((previous, current) => previous + current);
}

console.log(sum(1, 2, 3));
console.log(sum(1, 2, 3, 4));

function myFun(a, b, ...manyMoreArgs) {
  console.log('a', a);
  console.log('b', b);
  console.log('manyMoreArgs', manyMoreArgs);
}

myFun('one', 'two', 'three', 'four', 'five', 'six');
