package main

import "time"

func main() {
   var p chan int
   p = make(chan int,1)
   go addch(p)
   a := <- p
   println(a)
}

func addch(p chan int) {
   time.Sleep(1*time.Second)
   p <- 1
}

