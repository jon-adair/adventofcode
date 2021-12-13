let fs = require("fs");
const { networkInterfaces } = require("os");
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

let dots = new Set();
let ai = 0
while(true) {
	if (arr[ai] == '') {
		ai++
		break;
	}
	let x = Number(arr[ai].split(',')[0])
	let y = Number(arr[ai].split(',')[1])
	dots.add(x*10000+y)

	ai++;
}
console.log(dots)
console.log(dots.size);

while(true) {
	console.log("folding", arr[ai]);
	if (arr[ai].startsWith('fold along x')) {
		let newDots = new(Set)
		let f = Number(arr[ai].split('=')[1]);
		for (d of dots) {
			x = Math.floor(d / 10000)
			y = d % 10000
			//console.log("checking", x,y);
			if (x>f) {
				x = f - (x-f)
			}
			//console.log("new dot",x,y)
			newDots.add(x*10000+y)
		}
		dots = newDots;

	} else {
		let newDots = new(Set)
		let f = Number(arr[ai].split('=')[1]);
		for (d of dots) {
			x = Math.floor(d / 10000)
			y = d % 10000
			//console.log("checking", x,y);
			if (y>f) {
				y = f - (y-f)
			}
			//console.log("new dot",x,y)
			newDots.add(x*10000+y)
		}
		dots = newDots;

	}
	console.log("new size:",dots.size);
	ai++
	if (ai==arr.length) { break }
}

//console.log(dots)


// I know by looking at the output that I have a range of about 100x5 but could check max values instead
let letters = new Array(6)
for (let j=0;j<6;j++) {
	letters[j] = new Array(120).fill(' ')
}

for (d of dots) {
	x = Math.floor(d / 10000)
	y = d % 10000
	letters[y][x] = "#"
}
for (let j=0;j<6;j++) {
	console.log(letters[j].join(''));
}


