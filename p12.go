package main

import (
    "fmt"
    //"time"
    "runtime"
)

var a string
var c = make(chan int)

func f() {
   a = "hello"
   c <- 1
}

func g() {
  for i:=0; i<10; i++ {
     fmt.Printf("%v\t",i)
     //time.Sleep(time.Second)
  }
  c <- 1
}

func h() {
  for i:=10;i<20;i++ {
     fmt.Printf("%v\t",i)
     //time.Sleep(time.Second)
  }
  c<-1
}

func main() {
  runtime.GOMAXPROCS(2)
  go g()
  go h()
  <-c
  fmt.Println("")
  <-c
  fmt.Println("")
}

