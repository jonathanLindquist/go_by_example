package main

import "fmt"

func sum(nums ...int) {
  fmt.Println(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
  sum(1,2)
  sum(1,2,3)

  nums := []int{1,2,3,4}
  sum(nums...)

  nums = make([]int, 3)
  nums[0] = 5
  nums[1] = 6
  nums[2] = 7
  sum(nums...)

  //WILL THROW AN ERROR, ARRAYS CANNOT BE USED
  //TYPE [2]int is not same as TYPE []int

  // var a [2]int
  // a[0] = 0
  // a[1] = 1
  // sum(a...)
}
