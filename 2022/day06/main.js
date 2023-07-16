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

// 	let [_,n,f,t] = l.split(/\D+/).map(Number) // gotta remember this one for quickly splitting ints

// --- crude boilerplate ends ---

// arr[0] = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
let l = arr[0]
// l = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

console.log(l)


let f = wl => {
	for (let i = 0; i<=l.length-wl; i++) {
		let w = l.slice(i,i+wl)
		let s = new Set(w)
		// console.log(w, s.size, s)
		if (s.size == wl) {
			console.log("found at ",i+wl)
			return i+wl
		}
	}
}
console.log(f(4)) // 1833
console.log(f(14)) // 3425

