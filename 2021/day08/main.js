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
// arr = arr[0].split(',').map(Number);

console.log("read %d lines", arr.length);
console.log("first: %s, last: %s\n", arr[0], arr[arr.length-1]);

// Really suffered from being too sleepy when I did this one
// I fumbled a few approaches then wound up solving for each wire connection
// There are simpler solutions that worked on the whole segment grouping


// 0 1 2 3 4 5 6 7 8 9
// 6 2 5 5 4 5 6 3 7 6
if (false) {
let sum = 0
for (a of arr) {
	let outputs = a.split(" | ")[1].split(" ")
	console.log(outputs)
	for (o of outputs) {
		if (o.length != 5 && o.length != 6) {
			sum++
		}
	}
}
console.log(sum) // 390
}


let sevenSegmentsShort = new Map([
	["abcefg", 0],
	["cf", 1],
	["acdeg", 2],
	["acdfg", 3],
	["bcdf", 4],
	["abdfg", 5],
	["abdefg", 6],
	["acf", 7],
	["abcdefg", 8],
	["abcdfg", 9],
])

sum = 0
for (a of arr) {
	let mapping = new Map()
	let imapping = new Map()
	let set5 = new Set()
	let set6 = new Set()
	let segments = new Map()
	let isegments = new Map()
	let possibles = new Map()
	for (c of "abcdefg") {
		possibles.set(c,new Set([..."abcdefg"]))
	}
	console.log(possibles)
	let inputs = a.replace(" | ", " ").split(" ")
	let outputs = a.split(" | ")[1].split(" ")

	/*
	for (p of permutator([..."abcdefg"])) {
		for (i of inputs) {
			if ()
		}
	}
	*/
	
	// console.log(inputs)
	for (o of inputs) {
		o = o.split('').sort().join('')
		if (mapping.has(o)) { 
			// console.log(`already had ${o}: ${segments.get(o)}`); 
			continue; 
		} else {
			// So I can easily determine 1, 4, 7, 8 just from counts
			// do I need to know the actual 7 segments?
			// or can I just check for unions?
			// I think I need to figure out the segments
			
			// 0 1 2 3 4 5 6 7 8 9
			// 6 2 5 5 4 5 6 3 7 6
			if (o.length == 2) { mapping.set(o,1); imapping.set(1,o); } 
			else if (o.length == 4) { mapping.set(o,4); imapping.set(4,o); } 
			else if (o.length == 3) { mapping.set(o,7); imapping.set(7,o); } 
			else if (o.length == 7) { mapping.set(o,8); imapping.set(8,o); } 
			else if (o.length == 5) { set5.add(o); }
			else if (o.length == 6) { set6.add(o); }
		}
	}
	// 5: 2 3 5
	// 6: 0 6 9
	// so the diff between 1 and 7 is a
	// if I look at all the 5lens, I can tell what's a d g and bcef
	// 8 is useless 
	// so 2 3 5 I'll have 2 with c lit, 2 with f lit, 1 with b lit
	// doh;
	// 0 tells me d, 6 tells me c, 9 tells me e
	// double-doh - I don't know which is 0 6 or 9 yet
	// --- 
	// 

	// use the knowledge that 7 - 1 gives us the 'a'
	segmentA = findDiff(imapping.get(7),imapping.get(1))
	segments.set(segmentA,'a')
	// possibles.set(segmentA, new Set(['a']))
	for ([p,ps] of possibles.entries()) {
		if (p === segmentA) {
			ps.clear()
			ps.add('a')
		} else {
			ps.delete('a')
		}
	}

	// mark the ones we know: 1 4 7 8

	/*
		for ([p,ps] of possibles.entries()) {
			if (p === imapping.get(1)) {
				ps.clear()
				ps.add('a')
			} else {
				ps.delete('a')
			}
		}
		*/

		/*
		think about how I would solve it by hand
		1: ab
		4: abef
		7: abd
		8: abcdefg

		5's (2 3 5): bcdef acdfg abcdf
		6's (0 6 9): abcdef bcdefg abcdeg

		So know d -> a from 7-1
		and (ab) -> (cf) from 7 and 1
		and (ef) -> (bd) from 4 and 1
		8 is useless
		// count the segments in 5's: 3a 1b 2c 3d 1e 2f 3g
		// count the segments in 6's: 3a 3b 2c 2d 2e 3f 3g
		so from the 5's e and g appear once (be)
		from the 6's e appears 3, g appears 2
		so e -> b and g -> e
		so now I know: d->a, e->b, g->e 
		from the 4-1 I now know f->d
		from the 6's I know there's just one 2 left and it'll map to c
		that's a->c
		so now I know: d->a, e->b, g->e, f->d, a->c
		from 5's there's just one 3 left (g) that's c->g
		from 6's now just one 3 left (f) that's b->f
		so now I know: d->a, e->b, g->e, f->d, a->c, c->g, b->f
		that's it right?
		yep

		so how do I code that?
		don't need possibles at all, just apply that process
		need a "give me the letters that appear N times in this map/set"
		so get 5's.count(1) intersection 6's.count(3) that's e
		then 6's count(2) 5's count(1) that's g



		*/

		console.log("5,1:",appears(set5,1))
		console.log("5,2:",appears(set5,2))
		console.log("5,3:",appears(set5,3))
		console.log("6,1:",appears(set6,1))
		console.log("6,2:",appears(set6,2))
		console.log("6,3:",appears(set6,3))

		let A = appears(set5,1)
		let B = appears(set6,3)
		console.log("b is", intersection(A,B))
		segments.set(intersection(A,B).values().next().value,'b')

		B = appears(set6,2)
		console.log("e is", intersection(A,B))
		segments.set(intersection(A,B).values().next().value,'e')

		console.log("4 is",imapping.get(4))
		console.log("1 is",imapping.get(1))
		B = subtract(imapping.get(4), imapping.get(1))
		console.log("map for b is", getKeyByValue(segments,"b"))
		B = subtract(B,getKeyByValue(segments,"b"))
		console.log("d is ", B)
		segments.set(B.values().next().value,'d')

		A = appears(set6,2)
		console.log("known:",getKeys(segments))
		A = subtract(A,getKeys(segments))
		console.log("c is", A)
		segments.set(A.values().next().value,'c')

		A = appears(set5,3)
		console.log("known:",getKeys(segments))
		A = subtract(A,getKeys(segments))
		console.log("g is", A)
		segments.set(A.values().next().value,'g')

		A = appears(set6,3)
		console.log("known:",getKeys(segments))
		A = subtract(A,getKeys(segments))
		console.log("f is", A)
		segments.set(A.values().next().value,'f')

	// floundering around here. Think about my needs. 
	// I ultimately need to map the input strings to numbers
	// but to get there I need to map the input strings and characters to segments
	// diff between 1 and 7 gives me a
	// count the segments in 5's: 3a 1b 2c 3d 1e 2f 3g
	// count the segments in 6's: 3a 3b 2c 2d 2e 3f 3g
	// yeah I think I'll need to start with possibles and eliminate them
	// So all seven segments start with all possibilities
	// then take the ones I know and eliminate them
	// like the simplest example, 1 is 'ab' so I can eliminate a and b for outputs abdeg
	// 7 + 1 combination tells me a
	// the 5length that contains 1 is 3
	// what about the opposite direction? What if I just try to iterate over the 
	//   mapping and see what the results are? I've got 1 to start. is it 6! remaining?

	console.log("for line: ", a)
	console.log("possibles: ", possibles)
	console.log("segments: ", segments)
	console.log("mapping: ", mapping)
	console.log("imapping: ", imapping)
	console.log("fives: ", set5)
	console.log("sixes: ", set6)

	// OK so I get here and I have segments which maps input letter to output segment
	// now i need to loop through outputs and map each to the right number
	console.log()
	let subsum = 0
	for (o of outputs) {
		console.log("Mapping",o)
		let mapped = rewire(o,segments)
		console.log(mapped)
		console.log(sevenSegmentsShort.get(mapped))
		subsum = sevenSegmentsShort.get(mapped) + subsum * 10
	}
	
	sum += subsum
	console.log();
}
console.log(sum) // 1011785

function rewire(i,s) {
	return [...i].map(c=>s.get(c)).sort().join('')
}

/*

found a more clever solution that uses a few things:
- the 4 unique lengths (1 4 7 8)
- then of the 6-lengths (0 6 9):
  - only 9 contains everything in 4
  - (now) only 0 contains everything in 1
  - (now) only 6 is left
- same with 5-lengths (2 3 5)
  - only 3 contains 1
  - complex: only 5 shares 3 segments with 4
  - only 2 is left
- 

.split(/\r?\n/).filter(i => i).map(e => e.split('|').map(i => i.split(' ').filter(i => i).map(e => [...e].sort())));
	// 5: 2 3 5
	// 6: 0 6 9

   0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

*/


function findDiff(str1, str2) { 
	console.log("diff",str1, str2)x
	if (str1.length<str2.length) { console.log("longer string first"); }
	let m = new Set();
	for (c of str1) {
		m.add(c)
	}
	for (c of str2) {
		m.delete(c)
	}
	console.log(m)
	let diff=[...m].join('');
	console.log("diff:", str1, str2, diff)
	return diff;
  }

  const permutator = (inputArr) => {
	let result = [];
  
	const permute = (arr, m = []) => {
	  if (arr.length === 0) {
		result.push(m)
	  } else {
		for (let i = 0; i < arr.length; i++) {
		  let curr = arr.slice();
		  let next = curr.splice(i, 1);
		  permute(curr.slice(), m.concat(next))
	   }
	 }
   }
  
   permute(inputArr)
  
   return result;
  }

//   console.log(permutator([..."abcdefg"]))

function intersection(a,b) {
	return new Set([...a].filter(x => b.has(x)));
}

function subtract(A,B) {
	let a = new Set([...A])
	let b = new Set([...B])

	return new Set([...a].filter(x => !b.has(x)));
}

function getKeyByValue(set, value) {
	for ([k,v] of set.entries()) {
		if (v == value) { return k }
	}
}

function getKeys(map) {
	return new Set([...map.keys()])
}

function appears(s,n) {
	ss = new Set()
	m = new Map()
	for (xx of s) {
		for (x of xx) {
			// console.log("counting",x)
			if (!m.has(x)) {
				m.set(x,0)
			}
			m.set(x,m.get(x)+1)
		}
	}
	// console.log(m)
	for ([x,_] of m.entries()) {
		if (m.get(x) == n) {
			ss.add(x)
		}
	}
	// console.log(`appears(${n},${s}) -> ${ss.toString()}`)
	// console.log(s)
	// console.log(ss)
	return ss
}