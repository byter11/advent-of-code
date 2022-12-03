package main

import (
	"fmt"
	"io"
)

var wtg map[byte]int

type round struct {
  a byte
  b byte
}

func (r *round) points() int {
  if (wtg[r.b] + 1) % 3 == wtg[r.a] {
    return ord(r.b) + 7  // win
  } else if (wtg[r.a] + 1) % 3 == wtg[r.b] {
    return ord(r.b) + 1      // lose
  } else {
    return ord(r.b) + 4  // draw
  }
}

func (r *round) stratPoints() int {
  // hmmmmm
  if r.b == 'X' {
    if r.a == 'A' {
      return ord('C') + 1
    } else if r.a == 'B' {
      return ord('A') + 1
    } else { return ord('B') + 1 }
  } else if r.b == 'Y' {
    return ord(r.a) + 4
  } else {
    return (ord(r.a) + 1) % 3 + 7
  }
}

func init() {
  wtg = map[byte]int{
    'A': 0,
    'B': 2,
    'C': 1,
    'X': 0,
    'Y': 2,
    'Z': 1,
  }
  // A > C > B > A
}

func ord(c byte) int {
  i := int(c)
  if i < 68 {
    return i-65
  } else {
    return i-88
  }
}

func main() {
  var (
    a byte
    b byte
    rounds []round
  ) 

  for {
    _, err := fmt.Scanf("%c %c\n", &a, &b)
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Println(err)
      break
    }
    
    rounds = append(rounds, round{
      a: a, 
      b: b,
    })
  }

  sum := 0
  stratSum := 0
  for _, r := range rounds {
    sum += r.points()    
    stratSum += r.stratPoints()
  }

  fmt.Printf("Part One: %d\n", sum)
  fmt.Printf("Part Two: %d\n", stratSum)
}
