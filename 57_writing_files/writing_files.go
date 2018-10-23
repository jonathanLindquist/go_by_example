package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  d1 := []byte("hello\ngo\n")
  err := ioutil.WriteFile("/tmp/dat", d1, 0644)
  check(err)

  f, err := os.Create("/tmp/dat2")
  check(err)

  defer f.Close()


  d2 := []byte{115, 111, 109, 101, 10}
  n2, err := f.Write(d2)
  check(err)
  fmt.Printf("wrote %d bytes\n", n2)

  n3, err := f.WriteString("writes\n")
  fmt.Printf("wrote %d bytes\n", n3)

  //flush writes to stable storage
  f.Sync()

  w := bufio.NewWriter(f)
  n4, err := w.WriteString("bufferd\n")
  fmt.Printf("wrote %d bytes\n", n4)

  //make sure all buffered operations are applied to underlying writer
  w.Flush()
}
