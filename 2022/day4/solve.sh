#!/bin/bash

cat in.test \
  | tr ',-' '\n\n' \
  | xargs -n4 bash -c '(    ([ "$2" -ge "$0" ] && [ "$3" -le "$1" ]) \
                         || ([ "$0" -ge "$2" ] && [ "$1" -le "$3" ])) && echo 1 \
                         || echo 0' \
  | awk '{s+=$1} END {print "Part One: "s}'


cat in.prod \
  | tr ',-' '\n\n' \
  | xargs -n4 bash -c '(    ([ "$2" -le "$1" ] && [ "$2" -ge "$0" ]) \
                         || ([ "$3" -ge "$0" ] && [ "$3" -le "$1" ])) \
                         || ([ "$0" -ge "$3" ] && [ "$0" -le "$2" ]) \
                         || ([ "$1" -ge "$2" ] && [ "$1" -le "$3" ]) && echo 1 \
                         || echo 0' \
  | awk '{s+=$1} END {print "Part Two: "s}'
