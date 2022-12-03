#!/bin/sh

cat in.prod | sed -e 's/X/1/g' -e 's/Y/2/g' -e 's/Z/3/g' | sed -e 's/A 2/8/g' -e 's/B 3/9/g' -e 's/C 1/7/g' | sed -e 's/A 1/4/g' -e 's/B 2/5/g' -e 's/C 3/6/g'  | sed -e 's/[A |B |C ]//g' | awk '{s+=$1} END {print "Part One: "s}'
