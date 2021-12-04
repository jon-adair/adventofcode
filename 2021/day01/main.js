let fs = require("fs");

// let arr = fs.readFileSync('input.txt').toString('UTF8').split('\n').map(Number);
let arr = fs.readFileSync('in.txt').toString('UTF8').split('\n').map(Number);

console.log("read %d lines", arr.length);
console.log("first: %d, last: %d\n", arr[0], arr[arr.length-1]);

// pretty straightforward

A:
var last = arr[0];
let count = 0;
for (i of arr) {
	if (i > last) { last = i; count++ }
	last = i;
}
console.log(count);

B:
count = 0;
last = arr[0] + arr[1] + arr[2];
for (let i = 2; i < arr.length; i++) {
	sum = arr[i-2] + arr[i-1] + arr[i];
	if (sum > last) { count++ }
	last = sum;
}

console.log(count)