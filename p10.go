package main

import (
   "os"
   "fmt"
)

func main() {
   err := os.Mkdir("a",os.ModePerm)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("mkdir sucess")
   err = os.Remove("/home/go/"+"a")
}
