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

let sums = []
let sum = 0

for (a of arr) {
  console.log(a)
  if (a == "") {
  	console.log("push")
  	sums.push(sum)
	sum = 0
  } else {
	sum += parseInt(a)
  }
  console.log("  ",sum)

}
sums.push(sum)

console.log(sums)
console.log(Math.max(...sums))

// part 2
sums=sums.sort((a,b)=>b-a)
console.log(sums[0]+sums[1]+sums[2])
