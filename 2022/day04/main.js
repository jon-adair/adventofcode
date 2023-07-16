const { networkInterfaces } = require("os");
const { argv } = require("process");
let fs = require("fs");
const { parse } = require("path");

// handle alternate test case input files
let infile = 'input.txt';
if (argv[argv.length-1].endsWith('.txt')) {
	infile = argv[argv.length-1];
}
console.log('processing %s', infile);

let arr = fs.readFileSync(infile).toString('UTF8').split('\n');
// drop any extra EOLs at the end of the file
while (arr[arr.length-1] == '') {
	console.log('dropping trailing blank line / EOL');
	arr.pop();
}

// convert to "int" if needed
// arr = arr[0].split(',').map(Number);

console.log("read %d lines", arr.length);
console.log("first: %s, last: %s\n", arr[0], arr[arr.length-1]);

// --- crude boilerplate ends ---

let sum = 0
let sum2 = 0
for (l of arr) {
	console.log(l)
	let [s1,e1,s2,e2] = l.split(/\D/).map(Number) // gotta remember this one for quickly splitting ints
	console.log(" ",s1,e1,s2,e2)
	if ((s1>=s2 && e1<=e2)||(s1<=s2&&e1>=e2)) {
		sum++
	}
	// gotta think. does s1 or e1 fall within s2-e2 and vice-versa?
	if ((s1>=s2&&s1<=e2) || (e1>=s2&&e1<=e2) || (s2>=s1&&s2<=e1) || (e2>=s1&&e2<=e1) ) {
		sum2++
	}
}
console.log(sum) // 582
console.log(sum2) // 893
