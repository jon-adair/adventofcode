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

// pretty straight-forward grind with a recursive depth-first search
// for part 2 I did a little ugly and validate the path each call - there are more efficient ways 
// I did waste a chunk of time because when I recursed on paths() I left out the part2 param so after the first call it was skipping part 2 logic

// spent a bit before I started thinking about the reaches data structure.
// Almost went with a 2D array instead 
// would like to be able to filter() a map :(

// still annoying that has() autocompletes on Array 

let reaches = new Map()

for (a of arr) {
	let s = a.split('-')
	if (!reaches.has(s[0])) {
		reaches.set(s[0],[s[1]])
	} else {
		reaches.get(s[0]).push(s[1])
	}
	if (!reaches.has(s[1])) {
		reaches.set(s[1],[s[0]])
	} else {
		reaches.get(s[1]).push(s[0])
	}
}

console.log(reaches)

let sum = paths('start',[],false)
console.log(sum)
console.log() 

sum = paths('start',[],true)
console.log(sum) // 119760

let v = []


function paths(current,visited,part2) {
	// console.log("checking path",current,visited)
	// handle part 2 by throwing out invalid paths (more than 1 small cave visited more than once)
	if (part2) {
		let twice = false
		for (sc of reaches.keys()) {
			if (sc != sc.toLowerCase()) { continue; }
			let c = visited.filter(e=>sc==e).length
			// console.log("count for",visited,sc,c)
			if (c>2) { return 0; }
			if (twice && c==2) { return 0; }
			if (!twice && c==2) { twice = true;}
		}
		//if (current=='start' || current=='end') { return 0; }
	}

	if (current == 'end') {
		console.log("  at end", visited)
		return 1;
	}
	if (current == current.toLowerCase() && visited.includes(current)) { // has, includes, contains?
		if (part2) {
			if (current=='start' || current=='end') { return 0; }
		} else {
			// console.log("already visited",current,visited)
			return 0;
		}
	}
	visited.push(current)
	let s = 0
	for (r of reaches.get(current)) {
		// console.log("  trying",r)
		s += paths(r,visited,part2)
	}
	visited.pop()
	return s
}