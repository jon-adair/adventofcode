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

/*
for(var i=0;i<arr.length;i++) {
	var line = ""
	for (var j=0;j<arr[i].length;j++) {
		line += arr[i][j];
	}
	console.log(line)
}
*/

console.log("read %d lines", arr.length);
console.log("first: %s, last: %s\n", arr[0], arr[arr.length-1]);


// From someone else:
// const input = fs.readFileSync(path.join(__dirname, 'input.txt')).toString().trim().split('\n').map(x => x.split(' '))
// then x would be an array ['forward', 2]
// another tip: didn't need to check for whole string just first letter

// I had a few slowdowns to work on:
// - left day 1 files open in editor causing some confusion (did this all the time last year)
// - didn't have a terminal open to right place to run code
// - spent time arranging / finding windows - just need to have them all positioned ahead of time
// - had to look up startsWith() - didn't autocomplete so I wasn't sure of myself - why not? doesn't know type of x?
//   yeah - could:
//   	for (xx of arr) {
//			let x = String(xx);
// - tried to .parseInt() which isn't valid syntax
// - a couple coughing fits cost me almost a minute and probably 1000 spots or more
/*
 * so for future days:
 * rename main.js to day02.js
 * add easy toggle between input files or just run on argv[1]
 * auto handle eol issue
 * strip to a framework well ahead of time
 * open a terminal and position windows
 * coughdrop in, drink ready
 * would be good to have some common idioms handy
 */


let h = 0;
let d = 0;

A:
for (xx of arr) {

	// this is clumsy but would have made vscode's autocomplete work
	let x = String(xx);
	if (x.startsWith('forward')) {
		//h += Number(x.split(' ')[1]);
		h += x.split(' ').map(Number)[1]
	}

	if (x.startsWith('up')) {
		d -= Number(x.split(' ')[1]);
	}
	if (x.startsWith('down')) {
		d += Number(x.split(' ')[1]);
	}
}
console.log(h,d, h*d);

B:
h=0; d=0; let aim=0;

for (x of arr) {
	if (x.startsWith('forward')) {
		h += Number(x.split(' ')[1]);
		d += Number(x.split(' ')[1]) * aim;
	}
	if (x.startsWith('up')) {
		aim -= Number(x.split(' ')[1]);
	}
	if (x.startsWith('down')) {
		aim += Number(x.split(' ')[1]);
	}
}
console.log(h,d, h*d);

