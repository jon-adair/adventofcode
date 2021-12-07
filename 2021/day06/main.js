let fs = require("fs");
const { argv } = require("process");

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
arr = arr[0].split(',').map(Number);

console.log("read %d lines", arr.length);
console.log("first: %s, last: %s\n", arr[0], arr[arr.length-1]);

let fish = new Array(9).fill(0);
fish
for (let f of arr) {
	fish[f]++
}

for (let d=0;d<256;d++) {
	// console.log(d);
	let hatch=new Array(9).fill(0);
	for (let i=0;i<fish.length;i++) {
		if (i==0) {
			hatch[8] += fish[i]
			hatch[6] += fish[i]
		} else {
			hatch[i-1] += fish[i]
		}
	}
	//console.log(hatch)
	fish = hatch;
	// console.log(d,fish.reduce((s,e)=>s+e,0))
}
console.log(fish);
console.log(fish.reduce((s,e)=>s+e,0))