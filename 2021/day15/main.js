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
f[0][0] = 0
console.log(f)

// Ok so I know there are path-finding algorithms I can look up and implement
// but I don't want to. What can I do from memory or just reasoning?
// 

/*
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581

I think I can track the cheapest path to each point I've visited and walk from that.
Take the cheapest "frontier" and expand that each time.

So:
0

0 1
1

0 1 7 
1 4
3 

then I keep growing from the 3, or do I just do the whole diagonal?

0 1 7
1 4
3 4
6 

0 1 7 
1 4 
3 4 7 
6 10

Field is square and I don't think it's that complex to do 100-long diagonal

so each pass, run along the diagonal and find the cheapest route to get to that point

indexing that diagonal
first diagonal (r,c) 1,0 and 0,1
second: 2,0 1,1 0,2
3,0 2,1 1,2 0,3

....
no, that won't work if the cheapest path turns north or west.

So A* it is. Sigh
Let me at least code it from scratch.

So I need a list of: visited status, current?, shortest distance from start,
heuristic distance to end, total distance, previous node?
point to current node
what about the heuristic for end? ortho distance?


 1  2  8 11
 2  4 12 13


*/

let dim = f.length // we know input is square
console.log("dim", dim)
let nodes = []
for (let r=0;r<dim;r++) {
	for (let c=0;c<dim;c++) {
		nodes.push({r:r,c:c,dFromStart:r==0&&c==0?f[r][c]:Infinity,
			dToEnd:dim-r-1+dim-c-1,dTotal:Infinity,prevNode:"",visited:false})
	}
}
// console.log(nodes)
let current=nodes.filter(e=>e.r==0&&e.c==0)[0]
console.log(current)

while(current.r!=dim-1 || current.c!=dim-1) {
	let r = current.r; let c = current.c;
	if (r>0) {
		let p=nodes.filter(e=>e.r==r-1&&e.c==c)[0]
		if (!p.visited) {
			p.dFromStart = current.dFromStart + f[r-1][c]
			p.dTotal = p.dFromStart + p.dToEnd
		}
	}
	if (r<dim-1) {
		let p=nodes.filter(e=>e.r==r+1&&e.c==c)[0]
		if (!p.visited) {
			p.dFromStart = current.dFromStart + f[r+1][c]
			p.dTotal = p.dFromStart + p.dToEnd
		}
	}

	if (c>0) {
		let p=nodes.filter(e=>e.r==r&&e.c==c-1)[0]
		if (!p.visited) {
			p.dFromStart = current.dFromStart + f[r][c-1]
			p.dTotal = p.dFromStart + p.dToEnd
		}
	}
	if (c<dim-1) {
		let p=nodes.filter(e=>e.r==r&&e.c==c+1)[0]
		if (!p.visited) {
			p.dFromStart = current.dFromStart + f[r][c+1]
			p.dTotal = p.dFromStart + p.dToEnd
		}
	}
	current.visited = true;
	let cameFrom = current;
	// current = nodes.filter(e=>!e.visited).reduce((min,e)=>!e.dFromStart<min ? e.dFromStart : min, Infinity)
	let min = Infinity
	for(n of nodes) {
		if (n.visited) { continue; }
		if (n.dTotal < min) {
			current = n
			min = current.dTotal
			// console.log("new min",min,n.r,n.c)
		}
	}
	current.cameFrom = cameFrom;
	console.log("new node:")
	console.log(current)


	// console.log(nodes)
	//break;
}

while(current.cameFrom) {
	console.log(current.r,current.c,current.dFromStart,f[current.r][current.c])
	current = current.cameFrom
}
