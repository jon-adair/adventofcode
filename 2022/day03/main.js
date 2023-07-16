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

let sum = 0
for (l of arr) {
	s1 = l.substring(0,l.length/2)
	s2 = l.substring(l.length/2)
	console.log(s1, "   ", s2)
	for (c of s1) {
		if (s2.includes(c)) {
			cc = c.charCodeAt(0)
			if (cc>96) { cc -= 96 } else { cc -= (64 - 26) }
			console.log("  ", c, cc)
			sum += cc
			console.log(" ",sum)
			break
		}
	}
}

console.log("")

let sumb = 0
for (let i=0;i<arr.length/3;i++) {
	for (c of arr[i*3]) {
		if (arr[i*3+1].includes(c) && arr[i*3+2].includes(c)){
			cc = c.charCodeAt(0)
			if (cc>96) { cc -= 96 } else { cc -= (64 - 26) }
			console.log("  ", c, cc)
			sumb += cc
			console.log(" ",sumb)
			break
		}
	}
}
