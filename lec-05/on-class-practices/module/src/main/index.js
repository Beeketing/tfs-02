const { sum, fullName: hoten } = require('../utils/lib');
const getUnixTime  = require('../utils/date');
const lodash = require('lodash');
const ocg = require('ocg');

console.log(sum(1, 2));
console.log(hoten('Phuc', 'Pham'));
console.log(getUnixTime());
console.log(lodash);
console.log(ocg);
