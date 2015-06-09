package main

import (
    "fmt"
)

func Count(ch chan int) {
   fmt.Println("haha")
   ch <- 1
   fmt.Println("counting...")
}

func main() {
   chs := make([]chan int, 10)
   for i:=0; i<10; i++ {
       chs[i] = make(chan int)
       go Count(chs[i])
   }
   for _, ch := range(chs) {
       <-ch
   }
}
