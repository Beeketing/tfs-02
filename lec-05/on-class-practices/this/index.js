const test = {
  prop: 42,
  func: function() {
    return this.prop;
  },
};

// console.log('Test', test.func());

// this in a Function
/*function fn1() {
  return this;
}

console.log('fn1', fn1());*/

// Error
const assign = test.func;
console.log('assign', assign());

// Fix
const assignBind = test.func.bind(test);
console.log('assign', assignBind());
