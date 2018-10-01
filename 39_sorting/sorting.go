package main

import "fmt"
import "sort"

func main() {
  strs := []string{"c", "b", "a"}
  sort.Strings(strs)
  fmt.Println("Strings:", strs)

  ints := []int{7,2,4}
  sort.Ints(ints)
  fmt.Println("Ints:  ", ints)

  s := sort.IntsAreSorted(ints)
  fmt.Println("Ints are sorted: ", s)

  a := sort.StringsAreSorted(strs)
  fmt.Println("Strings are sorted: ", a)
}
