package main

import "fmt"

func vals() (int, int, string) {
  return 3, 7, "go"
}

func main() {
  a, b, _ := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c, s1 := vals()
  fmt.Println(c)
  fmt.Println(s1)

  d, _, s2 := vals()
  fmt.Println(d)
  fmt.Println(s2)
}
