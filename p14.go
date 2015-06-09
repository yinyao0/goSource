package main

import (
   "fmt"
   "encoding/json"
)


type JJ struct {
   Name   string
   Order  string
}

func main() {
   str := []byte (`[
   {"Name":"Haha", "Order":"Monno"}
   ]`)
   var jj []JJ
   _ = json.Unmarshal(str,&jj)
   fmt.Printf("%+v\n",jj)
}
