#!/bin/bash
# get the year and day number from the directory name
# this assumes a directory structure of year/day
# also expects a cookie file at the year level directory

d=$(basename $(pwd))
d=${d:3}
# strip off leading zero
if [ ${d:0:1} == "0" ]; then
	d=${d:1}
fi

cd .. > /dev/null
y=$(basename $(pwd))
cd - > /dev/null

echo Grabbing input file for year $y day $d:
curl -o input.txt "https://adventofcode.com/2023/day/$d/input" -H "Cookie: session=$(cat ../cookie)"
echo ""
echo Input:
head -3 input.txt
echo "  ..."
tail -3 input.txt
