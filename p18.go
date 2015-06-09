package main

import (
   "fmt"
   "net"
   "runtime"
)

var ln net.Listener

/*
func handler(conn) {
   s := "hello"
   ss := []byte(s)
   n, _ := conn.Write(ss)
   fmt.Printf("%v\n",ss)
   conn.Close()
}*/

func main() {
   runtime.GOMAXPROCS(2)
   ln, err := net.Listen("tcp", ":8080")
   defer ln.Close()
   if err != nil {
      panic("listen error")
   }
   for {
      ln.Accept()
     //conn, err := ln.Accept()
     if err != nil {
       panic("connect error")
     }
     fmt.Println("aa")
     //go handler(conn)
   }
}
