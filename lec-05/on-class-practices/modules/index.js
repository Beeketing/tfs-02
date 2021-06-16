// Basic
import { sum } from './lib.js';
console.log(sum(5, 6));

// Import
import remove from './utils.js';
console.log(remove('Hello OCG, 123'))

import { sum as sumTwoNumbers } from './lib.js';
console.log(sumTwoNumbers(5, 7));

import { sumNumbers } from './lib.js';
console.log(sumNumbers(5, 6));
