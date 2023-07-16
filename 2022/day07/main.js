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

let t = { size: 0, path: "/", parent: null, subs: [] }
let p = [ "" ]
let tt = t

for (l of arr) {
	console.log("* (pre)",l, " :", p, tt)

	switch (l) {
		case "$ cd /": 
			p = [ "" ]
			tt = t
			console.log("cd",tt)
			break
		case "$ cd ..":
			p.pop()
			tt = tt.parent
			break
		case "$ ls":
			break
		default:
			if (l.startsWith("$ cd ")) {
				p.push(l.split(" ")[2])
				let path = p.join("/")
				let n = tt.subs.filter(e => e.path == l.split(" ")[2])[0]
				console.log("n:",n, tt.subs)
				n.path = path
				n.parent = tt
				n.size = 0
				n.subs = []
				tt = n
				console.log("cd",tt)
				

			} else if (l.startsWith("dir ")) {

				let n = {}
				n.path = l.split(" ")[1]
				n.parent = tt
				n.size = 0

				tt.subs.push(n)
				console.log("dir", tt)

				
			} else { // {size} filename
				console.log("file:", tt, parseInt(l.split(" ")[0]))
				let ttt = tt
				while (ttt != null) {
					ttt.size += parseInt(l.split(" ")[0])
					ttt = ttt.parent
				}

			}

			break
	}
	console.log("* (post)",l, " :", p, tt)


}

console.log()

console.log(t)
console.log()

var x = []
var cts = []
console.log(count(t))
console.log(70000000 - t.size)
console.log(cts.sort((a,b)=>b-a).filter(e=>e > 70000000 - t.size))
console.log(cts.sort((a,b)=>b-a))
console.log(x)
// 25640133
// 44359867
function count(t) {
	// console.log(t)
	cts.push(t.size)
	x.push({d:t.path, s:t.size})
	let s = 0
	console.log(t.size.toString().padStart(10),t.path)
	if (t.size < 100000) {
		s += t.size
	}
	for (c of t.subs) {
		s += count(c)
		console.log(c.size.toString().padStart(10),c.path)
	}
	return s
}