# real    0m2.631s
# user    0m2.333s
# sys     0m0.330s

time ( 
cat in.prod \
  | awk  '{f=substr($1,0,length/2); l=substr($1,length/2+1,length); print f"\n"l}' \
  | xargs -n2 bash -c 'comm -12 <(echo $0 | fold -w1 | sort) <(echo $1 | fold -w1 | sort) | head -n1 | od -t d1 ' \
  | awk '{if($2>=96) print $2-96; else if($2>=65) print $2-65+27}' \
  | awk '{s+=$1} END {print "Part One: "s}';

cat in.prod \
  | xargs -n3 bash -c 'comm -12 <(echo $0 | fold -w1 | sort) <(comm -12 <(echo $1 | fold -w1 | sort) <(echo $2 | fold -w1 | sort)) | head -n1 | od -t d1' \
  | awk '{if($2>=96) print $2-96; else if($2>=65) print $2-65+27}' \
  | awk '{s+=$1} END {print "Part Two: "s}'
)
