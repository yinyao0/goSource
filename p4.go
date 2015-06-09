package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.OpenFile("1",os.O_RDWR|os.O_CREATE,0666)
  //file, err := os.Open("1")
  defer file.Close()
  if err != nil {
     fmt.Println(err)
     os.Exit(1)
  }
  /*str := "this it test\n"
  b := []byte(str)
  n, err := file.Write(b)
  if err != nil {
     fmt.Println(err)
     os.Exit(1)
  }
  fmt.Println(n)*/
  buf := make([]byte,256)
  n, _ := file.Read(buf)
  os.Stdout.Write(buf[:n])
  fmt.Println(n)
}

