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

let f = []
for (j of arr) {
	let r = []
	for (i of j) {
		r.push(Number(i))
	}
	f.push(r)
}
console.log(f)

let flashes = 0
for (let step = 0; step < 1000; step++) {
	let flashed = new Array(10)
	for (let j=0; j<arr.length;j++) {
		flashed[j] = new Array(10).fill(false)
	}
	
	console.log(flashed)
	for (let i=0; i<arr[0].length;i++) {
		for (let j=0; j<arr.length;j++) {
			f[i][j]++
		}
	}
	while(true) {
		let anyFlashed = false
		for (let i=0; i<arr[0].length;i++) {
			for (let j=0; j<arr.length;j++) {
				if (f[i][j] > 9 && !flashed[i][j]) {
					anyFlashed = true;
					console.log("flashing",i,j)
					flashed[i][j] = true
					flashes++

					for (let di=-1;di<=1;di++) {
						for (let dj=-1;dj<=1;dj++) {
							let x = i+di; let y = j+dj;
							if (x>=0 && x<10 && y>=0 && y<10) {
								f[x][y]++
							}
						}
					}
				}
			}
		}
		if (!anyFlashed) { break; }
		console.log(f)
		console.log(flashed)
	}
	console.log(flashed)
	for (let i=0; i<arr[0].length;i++) {
		for (let j=0; j<arr.length;j++) {
			if (flashed[i][j]) {
				f[i][j] = 0;
			}
		}
	}

	console.log("After step",step+1);
	console.log(f)

	let allFlashed = true;
	for (let i=0; i<arr[0].length;i++) {
		for (let j=0; j<arr.length;j++) {
			if (!flashed[i][j]) {
				allFlashed = false;
				break
			}
		}
	}
	if (allFlashed) {
		console.log("All flashed at step", step+1)
		break;
	}

}

console.log(flashes)
