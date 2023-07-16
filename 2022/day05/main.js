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

let s = [
	"NSDCVQT",
	"MFV",
	"FQWDPNHM",
	"DQRTF",
	"RFMNQHVB",
	"CFGNPWQ",
	"WFRLCT",
	"TZNS",
	"MSDJRQHN"
]

// let s = [
// 	"ZN", "MCD", "P"
// ]

console.log(s)
for (l of arr) {
	let [_,n,f,t] = l.split(/\D+/).map(Number) // gotta remember this one for quickly splitting ints
	//let [_,n,__,f,___,t] = l.split(/ /).map(Number) // gotta remember this one for quickly splitting ints
	console.log(l,n,f,t)
	for (let i=0;i<n;i++) {
		s[t-1] += s[f-1].slice(-1)
		s[f-1] = s[f-1].substring(0,s[f-1].length-1)
	}
	console.log(s)

}

let top = ""
for (c of s) {
	top += c.slice(-1)
}
console.log(top)

let ss = [
	"NSDCVQT",
	"MFV",
	"FQWDPNHM",
	"DQRTF",
	"RFMNQHVB",
	"CFGNPWQ",
	"WFRLCT",
	"TZNS",
	"MSDJRQHN"
]
// let ss = [
// 	"ZN", "MCD", "P"
// ]


console.log("--------------")


console.log(ss)
for (l of arr) {
	// let [_,n,f,t] = l.split(/\D*/).map(Number) // gotta remember this one for quickly splitting ints
	let [_,n,__,f,___,t] = l.split(/ /).map(Number) // gotta remember this one for quickly splitting ints
	console.log(l,n,f,t)
	// for (let i=0;i<n;i++) {
		ss[t-1] += ss[f-1].slice(-n)
		ss[f-1] = ss[f-1].substring(0,ss[f-1].length-n)
	// }
	console.log(ss)

}

let top2 = ""
for (c of ss) {
	top2 += c.slice(-1)
}
console.log(top) // FRDSQRRCD
console.log(top2) //HRFTQVWNN