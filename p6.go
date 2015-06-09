package main

import (
   "fmt"
   "bufio"
   "os"
   "strings"
   "strconv"
)

func init() {
   fmt.Println("init")
}

func main() {
  fmt.Println("main")
  r := bufio.NewReader(os.Stdin)
  line, _, _ := r.ReadLine()
  //fmt.Printf("%v\n",int(line[0])-48)
  l := string(line)
  ll := strings.Split(l," ")
  fmt.Println(strconv.Atoi(ll[0]))
}
