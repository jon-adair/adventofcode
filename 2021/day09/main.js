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
// arr = arr[0].split(',').map(Number);

console.log("read %d lines", arr.length);
console.log("first: %s, last: %s\n", arr[0], arr[arr.length-1]);

// another basic grind it out implementation

// only trickery I did was to pad the field with a ring of 9's so I can avoid that annoying index check

let f = []
f.push(new Array(arr[0].length+2).fill(9))
for (a of arr) {
	let r = [9]
	
	for (n of a) {
		r.push(Number(n))
	}
	r.push(9)
	f.push(r)
}
f.push(new Array(arr[0].length+2).fill(9))
console.log(f)
let s = 0; let s2 = 1; let s3 = []
for (let i=1;i<f.length-1;i++) {
	for (let j=1;j<f[0].length-1;j++) {
		// console.log(i,j)
		if (f[i][j] < f[i+1][j] && f[i][j] < f[i-1][j] && f[i][j] < f[i][j+1] && f[i][j] < f[i][j-1]) {
			s += f[i][j] + 1
			console.log("low at",i,j,f[i][j])
			s2 = walk(i,j,f)
			console.log("basin size:",s2)
			s3.push(s2)
		}
	}
}

let s4 = s3.sort((a,b)=>b-a)
console.log(s4[0]*s4[1]*s4[2])

function walk(I,J,f) {
	console.log(" walk",I,J,f[I][J])
	// console.log(f)
	if (f[I][J] == 9) { 
		return 0
	}
	f[I][J] = 9;
	return 1+walk(I-1,J,f)+walk(I+1,J,f)+walk(I,J-1,f)+walk(I,J+1,f);
}