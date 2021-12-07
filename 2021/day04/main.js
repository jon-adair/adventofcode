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



let nums = arr[0].split(",")
console.log(nums)

let boards = []
for (let l=2;l<arr.length;l++) {
	let b = []
	for (let r=0;r<5;r++) {
		let rr = []
		console.log(arr[l+r])
		console.log(arr[l+r].replace(/\s+/g, " ").replace(/^\s/g, ""))
		console.log()
		b.push(arr[l+r].replace(/\s+/g, " ").replace(/^\s/g, "").split(" ").map(Number))
	}
	l+=5;
	boards.push(b)
}

let marks = []

for(let b of boards) {
	console.log(b)
	console.log()
	let m = []
	for (let i=0; i<5;i++) {
		m.push([false,false,false,false,false]);
	}
	marks.push(m)
}
console.log(marks)

function mark(n) {
	for(let b=0;b<boards.length;b++) {
		for(let i=0;i<5;i++) {
			for(let j=0;j<5;j++) {
				if (boards[b][i][j] == n) {
					marks[b][i][j] = true;
				}
			}
		}
	}
}

function checkboard(b) {
	for(let i=0;i<5;i++) {
		if (b[i][0]&&b[i][1]&&b[i][2]&&b[i][3]&&b[i][4]) {
			return true;
		}
		if (b[0][i]&&b[1][i]&&b[2][i]&&b[3][i]&&b[4][i]) {
			return true;
		}
	}
}

function getsum(b) {
	let sum = 0 
	for(let i=0;i<5;i++) {
		for(let j=0;j<5;j++) {
			if (!marks[b][i][j]) {
				sum += boards[b][i][j]
			}
		}
	}
	return sum
}
console.log()

A:
for (x of nums) {
	console.log("marking ", x)
	mark(x)
	for(let b=0;b<boards.length;b++) {
		if (checkboard(marks[b])) {
			console.log("winner")
			console.log(boards[b])
			console.log(marks[b])
			console.log(getsum(b),getsum(b)*x)
			break A
		}
	}
}
console.log()

let winners = []
for(let b=0;b<boards.length;b++) {
	winners.push(false)
}

B:
for (x of nums) {
	console.log("marking ", x)
	mark(x)
	for(let b=0;b<boards.length;b++) {
		if (!winners[b]&&checkboard(marks[b])) {
			console.log("winner ", b)
			console.log(boards[b])
			console.log(marks[b])
			console.log(getsum(b),getsum(b)*x)
			winners[b] = true;
			// break B
		}
	}
}

