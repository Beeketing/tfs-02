"use strict";

var fn = name => 'Hello, ' + name;

console.log(fn('OCG'));

var sum = (a, b) => a + b;

console.log(sum(5, 6));

function sum1(a, b) {
  return a + b;
}

console.log(sum1(6, 7));

var person = function person(name) {
  var age = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 10;
  console.log('Name', name);
  console.log('Age', age);
};

var room = _ref => {
  var {
    floor
  } = _ref;
  console.log('floor', floor);
};

person('OCG', 15);
room({
  floor: 7
});