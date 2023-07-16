const { networkInterfaces } = require("os");
const { argv } = require("process");
let fs = require("fs");

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

let score = 0
const v = {"X":1,"Y":2,"Z":3}
const w = {"AX":3,"AY":6,"AZ":0,"BX":0,"BY":3,"BZ":6,"CX":6,"CY":0,"CZ":3}
const b = {"AX":0+3,"AY":3+1,"AZ":6+2,"BX":0+1,"BY":3+2,"BZ":6+3,"CX":0+2,"CY":3+3,"CZ":6+1}
for (l of arr) {
	let [o,p] = l.split(" ")
	let s = v[p] + w[o+p]
	console.log(o,p,s)
	score += s
}
console.log(score)
score = 0
for (l of arr) {
	let [o,p] = l.split(" ")
	let s = b[o+p]
	console.log(o,p,s)
	score += s
}
console.log(score)
