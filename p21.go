package main

import "fmt"

func main() {
   var p chan int;
   p = make(chan int,2);

   go func(){
     p <- 1;
     p <- 2;
     close(p);
   }()

   for k:=range p {
     fmt.Println(k);
   }
}
