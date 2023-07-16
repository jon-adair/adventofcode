const { networkInterfaces } = require("os");
const { argv } = require("process");
let fs = require("fs");

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

let bits = ""
for (c of arr[0]) {
	let d = parseInt(c,16)
	bits += d.toString(2).padStart(4,"0")
}
console.log(bits)

let s1 = 0
pkt(bits)
console.log(s1)

function pkt(bits) {
	while(true) {
		if (bits.length == 0) { break; }
		// version
		let v = parseInt(bits.slice(0,3),2)
		let t = parseInt(bits.slice(3,6),2)
		console.log("version:",v, "type:",t)
		bits = bits.slice(6)
		console.log(bits)
		s1 += v
		if (t==4) { 
			console.log("literal")
			let l = 0
			while(true) {
				console.log(bits.slice(0,5))
				l = l * 16 + parseInt(bits.slice(1,5),2)
				if (bits[0] == "0") {
					bits = bits.slice(5)
					break
				}
				bits = bits.slice(5)
			}
			console.log("literal value:",l)
			continue
		}
		// else operator
		console.log("operator")
		if (bits[0] == '0') {
			// length in bits
			let l = parseInt(bits.slice(1,16),2)
			console.log("subpacket bits:",l)
			bits = bits.slice(16)
			console.log(bits)
			pkt(bits.slice(0,l))
			bits = bits.slice(l)
		} else {
			// length in subpackets
			let l = parseInt(bits.slice(1,12),2)
			console.log("subpackets:",l)
			bits = bits.slice(12)
			console.log(bits)
			pkt(bits)

		}

		// need to be able to recurse and loop on the whole parsing
		// I guess return the # of bits consumed? nah just consume them
		// need the running sum too

		// while(true) {
		// 	let d = 0
		// 	if 
		// }
		break
	}
}


