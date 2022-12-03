package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
  calorieArray := readCalories()
  sort.Slice(calorieArray, func(i, j int) bool {
    return calorieArray[i] > calorieArray[j]    
  })
  
  top3 := 0
  for i := 0; i<3; i++ {
    top3 += calorieArray[i]
  }
  fmt.Printf("Part One: %d\n", calorieArray[0])
  fmt.Printf("Part Two: %d\n", top3)
}

func readCalories() []int {
  sums := []int{}
  cur := 0
	for {
    var i int 
    _, err := fmt.Scanf("%d", &i)
    if err == io.EOF {
      break
    }
    if err != nil {
      sums = append(sums, cur)
      cur = 0
    } else {
      cur += i
    }
	}

  return sums
}
