package main

import (
   "fmt"
   "strings"
)

func main() {
  s := "SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings. If sep is empty, SplitAfterN splits after each UTF-8 sequence. The count determines the number of substrings to return"
  s1 := strings.Split(s," ")
  fmt.Println(strings.Join(s1[0:10]," "))
  fmt.Println(len(s1))
}
