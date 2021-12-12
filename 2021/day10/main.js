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

// pretty straight-forward stack implementation
// didn't golf at all but I saw others store the weights and matching all in one data structure shrug
// got a little bit for a second on sort() because I threw (a,b)=>a>b at it


let m = new Map([['(',')'],['[',']'],['{','}'],['<','>']])
let sum = 0; let sum3 = [];
// arr = ["{([(<{}[<>[]}>{[]{[(<()>"]
// {([(<[}>{[]{[(<()>
for (a of arr) {
	let s = []
	let mismatch = false
	console.log(a)
	for (c of a) {
		console.log(c)
		if (m.has(c)) {
			s.push(c)
			console.log("s push:",s)
		} else {
			cc = s.pop()
			console.log("s pop :",s)
			// console.log("checking",c,cc)
			if (c != m.get(cc)) {
				switch (c) {
					case ')': sum += 3; break;
					case ']': sum += 57; break;
					case '}': sum += 1197; break;
					case '>': sum += 25137; break;
				}
				console.log("mismatch",c,cc,sum)
				mismatch = true
				break
			} else {
				console.log("match",c,cc)
			}
		}
	}
	if (!mismatch) {
		console.log("remaining:",s.join(''))
		let sum2 = 0; 
		for (c of s.reverse()) {
			switch (m.get(c)) {
				case ')': sum2 = sum2*5 + 1; break;
				case ']': sum2 = sum2*5 + 2; break;
				case '}': sum2 = sum2*5 + 3; break;
				case '>': sum2 = sum2*5 + 4; break;
			}
		}
		console.log("sum2:",sum2)
		sum3.push(sum2);
	}
}
console.log(sum)
// console.log(sum2)
let sum4 = sum3.sort((a,b)=>a-b)
console.log(sum4[(sum4.length-1)/2])