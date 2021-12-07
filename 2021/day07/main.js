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


// kinda of straightforward 
// A little code golf:
// initially used reduce() to find min and max of arr - doh
// I feel like there's a better way to compute part B than that while loop
//   - yeah, I had the nagging feeling that it's a simple eqn and it is - triangle numbers
//   - need to be better at identifying and using those
// Feel like ranging from min to max is overkill
//   - seeing a few people sorting and searching around the median for a min
//   - someone else took median of diff for part 1 and artihmetic mean for part 2
//     https://github.com/michelefenu/advent-of-code/blob/main/Advent%20of%20Code%202021/modules/crab-utils.js
// Could totally replace loops with reduce() and use min
// I used just 999999999 as a big number - can use Infinity instead


let min = Math.min(...arr);
let max = Math.max(...arr);

let low = Infinity
for (let a=min; a<max; a++) {
	let s = 0;
	s = arr.reduce((s,e)=>s+Math.abs(e-a),0)
	if (s<low) {
		low = s
	}
	console.log(a,s)
}

console.log(low) // 340987

console.log()

low = Infinity
for (let a=min; a<max; a++) {
	let s = 0;

	for(let e of arr) {
		/*
		c=1;
		let d = a;
		while(true) {
			s+= c;
			c++;
			d = d>e ? d - 1 : d + 1;
			if (d==e) break;
		}
		*/
		let x = Math.abs(a-e)
		s += x * (x+1) / 2
		// console.log(`moving from ${a} to ${e} costs ${s}`)
	}
	if (s<low) {
		low = s
	}
	console.log(`sum for ${a} is ${s}\n`)
}

console.log(low) // 96987874