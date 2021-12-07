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

// day 4 was my best day on the leaderboard, nearly top 1000 on each part
// I really fumbled with splitting the boards though, needing to google the replace syntax
//   and then I didn't strip the leading space on lines, costing me time on part 2
// It would have been much faster for me to vi the input.txt and strip the extra spaces from the data
// Got a little twisted in my logic but not too bad
// Was totally expecting to check diagonals until I got right answer without them and re-read the text
// Got kind of lucky that I didn't have to reset anything for part 2 since it would just re-mark the same numbers
// also didn't write any logic to identify the *last* board to win but just read the output
//   and kept marking boards after they'd all won

let m = new Map();
let sum = 0;
for (let l of arr) {
	let s = l.split(" -> ");
	let x1 = Number(s[0].split(",")[0])
	let y1 = Number(s[0].split(",")[1])
	let x2 = Number(s[1].split(",")[0])
	let y2 = Number(s[1].split(",")[1])
	if (x1!=x2 && y1!=y2) { continue; } // skip diagonals
	console.log("   Range %d,%d -> %d,%d",x1,y1,x2,y2);
	if (x1>x2) { let t=x1; x1=x2; x2=t; }
	if (y1>y2) { let t=y1; y1=y2; y2=t; }
	console.log("Plotting %d,%d -> %d,%d",x1,y1,x2,y2);
	if (x1 != x2)
		while(x1<=x2){
			// console.log(x1)
			if (m.has(`${x1},${y1}`)) {
				let v = m.get(`${x1},${y1}`);
				if (v == 1) sum++;
				// console.log("Had ",x1,y1,v,sum)
				m.set(`${x1},${y1}`,v+1);
			} else { 
				// console.log(" ",x1);
				m.set(`${x1},${y1}`,1); 
				// console.log(m); 
			}
			x1++;
		}
	else 
		while(y1<=y2){
			// console.log(y1)
			if (m.has(`${x1},${y1}`)) {
				let v = m.get(`${x1},${y1}`);
				if (v == 1) sum++;
				// console.log("Had ",x1,y1,v,sum)
				m.set(`${x1},${y1}`,v+1);
			} else { 
				// console.log(" ",y1);
				m.set(`${x1},${y1}`,1); 
				// console.log(m); 
			}
			y1++;
		}
}
console.log(m);
console.log(sum);

m = new Map();
sum = 0;
for (let l of arr) {
	let s = l.split(" -> ");
	let x1 = Number(s[0].split(",")[0])
	let y1 = Number(s[0].split(",")[1])
	let x2 = Number(s[1].split(",")[0])
	let y2 = Number(s[1].split(",")[1])
	// if (x1!=x2 && y1!=y2) { continue; } // skip diagonals
	//if (x1==x2 || y1==y2) continue;
	console.log("   Range %d,%d -> %d,%d",x1,y1,x2,y2);
	// if (x1>x2) { let t=x1; x1=x2; x2=t; t=y1; y1=y2; y2=t; }
	// else if (y1>y2) { let t=y1; y1=y2; y2=t; t=x1; x1=x2; x2=t; }
	console.log("Plotting %d,%d -> %d,%d",x1,y1,x2,y2);
	let slope = {x:0,y:0} 
	if (x1>x2) slope.x = -1;
	if (x1<x2) slope.x = 1;
	if (y1>y2) slope.y = -1;
	if (y1<y2) slope.y = 1;
	console.log(slope);
	while(x1!=x2 || y1!=y2){
		// console.log("  ",x1,y1)
		if (m.has(`${x1},${y1}`)) {
			let v = m.get(`${x1},${y1}`);
			if (v == 1) sum++;
			// console.log("Had ",x1,y1,v,sum)
			m.set(`${x1},${y1}`,v+1);
		} else { 
			// console.log(" ",x1);
			m.set(`${x1},${y1}`,1); 
			// console.log(m); 
		}
		x1+=slope.x;y1+=slope.y;
	}
	// console.log("  ",x1,y1)
	if (m.has(`${x1},${y1}`)) {
		let v = m.get(`${x1},${y1}`);
		if (v == 1) sum++;
		// console.log("Had ",x1,y1,v,sum)
		m.set(`${x1},${y1}`,v+1);
	} else { 
		// console.log(" ",x1);
		m.set(`${x1},${y1}`,1); 
		// console.log(m); 
	}
	
}
for(let j=0;j<10;j++) {
	for(let i=0;i<10;i++) {
		if (!m.has(`${i},${j}`)) {
			process.stdout.write(".")
		} else 
		process.stdout.write(String(m.get(`${i},${j}`)))
	}
	process.stdout.write("\n")
}
// console.log(m);
console.log(sum);
