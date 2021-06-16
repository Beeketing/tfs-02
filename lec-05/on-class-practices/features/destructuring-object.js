// Destructuring object
const obj = { a: 1, b: 2, c: 4};
const { a } = obj;
console.log('basic', a);

// Default
const { defaultObj = 10 } = obj;
console.log('default', defaultObj);

// Alias
const { a: aliasA } = obj;
console.log('aliasA', aliasA);

// Dynamic name property
const key = 'a';
const { [key]: name } = obj;
console.log('Dynamic', name);

// Rest object after destructuring
const { a: restA, ...rest } = obj;
console.log('restA', restA);
console.log('rest', rest);
