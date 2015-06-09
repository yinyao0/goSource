package main

import (
   "fmt"
   "strings"
   "crypto/rand"
   "encoding/hex"
   //"io"
)

func main() {
   s := "hello"
   ss := strings.ToUpper(s)
   fmt.Println(s)
   fmt.Println(ss)

   b := make([]byte, 4)
   /*if _, err := io.ReadFull(rand.Reader,b); err!=nil {
      panic(err)
   }*/
   n, err := rand.Read(b)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%v\n",n)
   fmt.Printf("%v\n",b)
   str := hex.EncodeToString(b)
   fmt.Println(str)
}
