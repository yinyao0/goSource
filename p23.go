package main

import "time"
import "fmt"

func main() {
  ticks := time.Tick(1*time.Second)
  for i := range ticks{
     a := <-ticks
     fmt.Println(a)
     fmt.Println(i)
     println("aaaa")
  }
}
