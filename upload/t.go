package main

import (
  "fmt"
  "crypto/md5"
  "time"
)

func main() {
   h := md5.New()
   token := fmt.Sprintf("%x",h.Sum(nil))
   fmt.Println(token)
   t := time.Now()
   fmt.Println(t)
}
