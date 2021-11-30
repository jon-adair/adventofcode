let fs = require("fs");

let arr = fs.readFileSync('input.txt').toString('UTF8').split('\n').map(Number);

console.log("read %d lines", arr.length);
console.log("first: %d, last: %d\n", arr[0], arr[arr.length-1]);

A:
for (i of arr) {
	for (j of arr) {
		if (i+j == 2020) {
			console.log("A:",i*j);
			break A;
		}
	}
}


B:
for (i of arr) {
	for (j of arr) {
		for (k of arr) {
			if (i+j+k == 2020) {
				console.log("B:",i*j*k);
				break B;
			}
		}
	}
}