package main

import (
  "net/http"
  "fmt"
  //"io"
  //"os"
  "io/ioutil"
)

func main() {
   r, err := http.Get("http://www.baidu.com")
   if err != nil {
     return
   }
  defer r.Body.Close()
  //io.Copy(os.Stdout,r.Body)
  data, _ := ioutil.ReadAll(r.Body)
  fmt.Println(data.Cookie)
}
