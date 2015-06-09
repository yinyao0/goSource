package main

import "fmt"

const (
  WHITE = "a" 
  BALCK = "b"
  BLUE  = "c"
  RED   = "d"
  YELLOW= "e" 
)

func add(a *int) int {
   *a=*a+1
   return *a
}

func main() {
   fmt.Println("hello world")
   var x int
   x = 2
   y := add(&x)
   fmt.Println("x =",x)
   fmt.Println("y =",y)
   fmt.Println(WHITE)
   fmt.Println(RED)

   if x!=0 && y!=0 {
     fmt.Println("It is ok!")
   }
}

