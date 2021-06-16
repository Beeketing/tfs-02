const fn = name => 'Hello, ' + name;
console.log(fn('OCG'));

const sum = (a, b) => a + b;
console.log(sum(5, 6));

function sum1(a, b) {
  return a + b;
}

console.log(sum1(6, 7));

const person = (name, age = 10) => {
  console.log('Name', name);
  console.log('Age', age);
}

const room = ({ floor, }) => {
  console.log('floor', floor);
}

person('OCG', 15);
room({ floor: 7 });
