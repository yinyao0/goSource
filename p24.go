package main

import (
    "time"
    "math/rand"
    "fmt"
)

func main() {
   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   for i:=1;i<10;i++ {
      fmt.Println(r.Intn(10))
   }
}
