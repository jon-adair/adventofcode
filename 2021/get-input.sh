#!/bin/bash
# get the day number from the directory name (TODO: get the year too)
d=$(basename $(pwd))
d=${d:3}
# strip off leading zero
if [ ${d:0:1} == "0" ]; then
	d=${d:1}
fi
echo Grabbing input file for day $d:
curl -o input.txt "https://adventofcode.com/2021/day/$d/input" -H "Cookie: session=$(cat ../cookie)"
echo ""
echo Input:
head -3 input.txt
echo "  ..."
tail -3 input.txt
