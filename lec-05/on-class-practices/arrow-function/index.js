// Basic
const fn = name => 'Hello, ' + name;
console.log(fn('OCG'));

// Two args
const fn1 = (name, age) => 'Hello: ' + name + ', age: ' + age;
console.log(fn1('OCG', 8));

// Destructuring args
const fn2 = ({ name, age }) => 'Hello: ' + name + ', age: ' + age;
console.log(fn2({ name: 'OCG', age: 8 }));

// Destructuring args with defaults
const fn3 = ({ name, age = 5 }) => 'Hello: ' + name + ', age: ' + age;
console.log(fn3({ name: 'OCG' }));

// Using in callback
const arrays = [1, 2, 3, 4, 5];
arrays.forEach((number) => {
  const str = 'Number:' + number;
  console.log(str);
  return str;
});

const obj = {
  count : 10,
  doSomethingLater: function (){
    setTimeout(function(){
      this.count++;
      console.log(this.count);
    }, 300);
  }
}

obj.doSomethingLater();

const obj2 = {
  count : 15,
  doSomethingLater : function(){
    setTimeout( () => {
      this.count++;
      console.log(this.count);
    }, 300);
  }
}

obj2.doSomethingLater();

