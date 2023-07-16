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

let p = arr[0]

let r = new Map()
for (rr of arr.slice(2)) {
	let rrr = rr.split(" -> ") 
	r.set(rrr[0],rrr[1])
}
console.log(r)
let pp = new Map()
for (c of r.keys()) {
	pp.set(c,0)
}
for (let i=1;i<p.length;i++) {
	pp.set(p[i-1]+p[i],pp.get(p[i-1]+p[i])+1)
}
console.log("Initial",pp)

for (let s=0;s<40;s++) {
	let ppp = new Map()
	for (c of r.keys()) {
		ppp.set(c,0)
	}
		// console.log("mapping",pp.keys())
	for (pair of pp.keys()) {
		if (pp.get(pair) == 0) { continue; }
		// for each existing pair (in pp), need to:
		// map that to two new pairs and add the existing pair's count to each
		let left = pair.split('')[0] + r.get(pair)
		let right = r.get(pair) + pair[1].split('')
		console.log("mapping",pair,left,right,pp.get(left),pp.get(right))
		ppp.set(left,ppp.get(left)+pp.get(pair))
		ppp.set(right,ppp.get(right)+pp.get(pair))
	}
	pp = ppp
	console.log("Step",s)
	console.log(pp)

console.log(p[p.length-1])

let min = Infinity; let max = 0;
for (let a of "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
	let x = 0
	for (pair of pp.keys()) {
		if (pair.split('')[0] == a) {
			x += pp.get(pair)
		}
		if (pair.split('')[1] == a) {
			// x += pp.get(pair)
		}

	}
	if (a == p[p.length-1]) {
		x++
	}
// let x = [...pp].reduce((s,e)=>e==a?s+1:s,0)
	if (x==0) { continue; }
	console.log(a,x)
	max = Math.max(x,max)
	min = Math.min(x,min)	
}
console.log(max-min)
}

return

for (let s=0;s<4;s++) {
	let newP = ""
	for (let i=1;i<p.length;i++) {
		
		newP += p[i-1] + r.get(p[i-1]+p[i])

	}
	newP += p[p.length-1];
	p = newP
	// console.log("step",s,p)
	console.log("step", s, p.length)


let c = {}
let min = Infinity; let max = 0;
for (let a of "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
	let x = [...p].reduce((s,e)=>e==a?s+1:s,0)
	if (x==0) { continue; }
	console.log(a,x)
	max = Math.max(x,max)
	min = Math.min(x,min)	
}
console.log(max-min)
}