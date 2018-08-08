package main

import "fmt"

func intSeq(i int) func() int {
  local_i := i
  return func() int {
    local_i++
    return local_i
  }
}

func main() {
  nextInt := intSeq(0)

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  newInts := intSeq(5)
  fmt.Println(newInts())
}
