package main

import (
   "fmt"
   "runtime"
)

func say() {
   for {
   fmt.Println("hello")
   }
}

func f(a... int) {
   k := len(a)
   fmt.Println(k)
   fmt.Println(a)
   fmt.Printf("%v,%v,%v\n",a[1],a[2],a[3])
}

func main() {
   //runtime.GOMAXPROCS(2)
   //go say()
   fmt.Println(runtime.NumCPU())
   //for {
   //}
   f(1,2,3)
}
