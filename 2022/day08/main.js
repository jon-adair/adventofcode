const { networkInterfaces } = require("os");
const { argv } = require("process");
let fs = require("fs");
const { parse } = require("path");

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

// 	let [_,n,f,t] = l.split(/\D+/).map(Number) // gotta remember this one for quickly splitting ints

// --- crude boilerplate ends ---

let f = []
for (l of arr) {
	let r = []
	for (c of l) {
		r.push(parseInt(c))
	}
	f.push(r)
}

console.log(f)
console.log(f.length,f[0].length)
const n = f[0].length

let s = 0, smax = 0
for (r in f) {
	for (c in f[r]) {
		let ss = 1
		console.log("*** checking:", f[r][c], "at", r,c)
		let blocked = false, sss = 0
		for (let i = 0;i<c;i++) {
			console.log("  left check:",r,i)
			sss++
			if(f[r][c] <= f[r][i]) {
				console.log("  left blocked by",f[r][i],"at",r,i)
				blocked = true
				break
			}
		}
		console.log("   * ",sss)
		ss *= sss
		if (!blocked) {
			s++
			console.log("  clear to left")
			// continue
		}

		blocked = false
		sss = 0
		console.log("right range:",+c+1,n)
		for (let i = +c+1;i<n;i++) {
			console.log("  right check:",r,i)
			sss++
			if(f[r][c] <= f[r][i]) {
				console.log("  right blocked by",f[r][i],"at",r,i)
				blocked = true
				break
			}
		}
		console.log("   * ",sss)
		ss *= sss
		if (!blocked) {
			s++
			console.log("  clear to right")
			// continue
		}

		
		blocked = false
		sss = 0
		for (let i = 0;i<r;i++) {
			console.log("  top check:",i,c)
			sss++
			if(f[r][c] <= f[i][c]) {
				console.log("  top blocked by",f[i][c],"at",i,c)
				blocked = true
				break
			}
		}
		console.log("   * ",sss)
		ss *= sss
		if (!blocked) {
			s++
			console.log("  clear to top")
			// continue
		}

		blocked = false
		sss = 0
		for (let i = +r+1;i<n;i++) {
			console.log("  bottom check:",i,c)
			sss++
			if(f[r][c] <= f[i][c]) {
				console.log("  bottom blocked by",f[i][c],"at",i,c)
				blocked = true
				break
			}
		}
		console.log("   * ",sss)
		ss *= sss
		if (!blocked) {
			s++
			console.log("  clear to bottom")
			// continue
		}
		console.log("score:",sss)
	}
}

console.log("sum:",s)