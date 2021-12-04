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
//arr = arr.map(Number);

console.log("read %d lines", arr.length);
console.log("first: %s, last: %s\n", arr[0], arr[arr.length-1]);


// Well Day 3 was a shit-show. Partly due to being so tired (drops at midnight local time)
// Might have helped if I'd remembered JS's parseInt can handle base 2
// Partway into part 1 I have this nagging that I should just count 0s or 1s not both but kept going. That hurt in part 2.
// And looking at part 1 right now I see that of course gamma and epsilon sum to 2^n-1. Should have realized that.
// really should have used map + filter more



let g=0;
let e=0;
let b1 = []; let b0=[];
for (let j=0;j<arr[0].length;j++) {
	b0.push(0); b1.push(0);
}
for(let i=0;i<arr.length;i++) {
	for (let j=0;j<arr[i].length;j++) {
		if (arr[i][j] == '0') {
			b0[j]++;
		} else { b1[j]++ }
	}
	
}
console.log(b0);
console.log(b1);
for (let j=0;j<arr[0].length;j++) {
	if (b0[j]>b1[j]) {
		g = g*2 + 1;
		e = e*2;
	} else {
		g = g*2;
		e = e*2 + 1;
	}
}
console.log(g,e,g*e);

let bb1=[]; let bb0=[];
for (let j=0;j<arr.length;j++) {
	bb0.push(arr[j]); 
	bb1.push(arr[j]);
}

B1:
for (let k=0;k<1+0*arr.length;k++){

	for (let j=0;j<arr[0].length;j++) {
		for(let i=0;i<arr[0].length;i++) {
			b0[i]=0; b1[i] =0;
		}
		for(let i=0;i<bb0.length;i++) {
			for (let jj=0;jj<bb0[i].length;jj++) {
				if (bb0[i][jj] == '0') {
					b0[jj]++;
				} else { b1[jj]++ }
			}
			
		}
		console.log(k,bb0,b0,b1);
	
		if (b0[j]<=b1[j]) {
			console.log("removing bit %d == 0", j)
			bb0 = bb0.filter(function(item) {
				return String(item[j]) !== '0'
			})
		} else {
			console.log("removing bit %d == 1", j)
			bb0 = bb0.filter(function(item) {
				return String(item[j]) !== '1'
			})

		}
		if (bb0.length == 1) { break B1 }
		console.log(k,bb0);
	}
}

console.log(bb0);
console.log()
B2:
for (let k=0;k<1+0*arr.length;k++){

	for (let j=0;j<arr[0].length;j++) {
		for(let i=0;i<arr[0].length;i++) {
			b0[i]=0; b1[i] =0;
		}
		for(let i=0;i<bb1.length;i++) {
			for (let jj=0;jj<bb1[i].length;jj++) {
				if (bb1[i][jj] == '0') {
					b0[jj]++;
				} else { b1[jj]++ }
			}
			
		}
		console.log(k,bb1,b0,b1);
	
		if (b0[j]>b1[j]) {
			console.log("removing bit %d == 0", j)
			bb1 = bb1.filter(function(item) {
				return String(item[j]) !== '0'
			})
		} else {
			console.log("removing bit %d == 1", j)
			bb1 = bb1.filter(function(item) {
				return String(item[j]) !== '1'
			})

		}
		if (bb1.length == 1) { break B2 }
		console.log(k,bb1);
	}
}

console.log(bb1)

console.log()

console.log(bb0)
console.log(bb1)

let a=0; let b=0;
for (let j=0;j<bb0[0].length;j++) {
	if (bb0[0][j]=='0') {
		a = a*2;
	} else {
		a = a*2 + 1;
	}
}
for (let j=0;j<bb1[0].length;j++) {
	if (bb1[0][j]=='0') {
		b = b*2;
	} else {
		b = b*2 + 1;
	}
}
console.log(a,b,a*b);

