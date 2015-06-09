package main

import (
    "fmt"
    "os"
    "bufio"
)

func Help(args string) int {
   fmt.Println(args)
   return 0
}

func Play(args string) int {
   fmt.Println(args)
   return 0
}

func Quit(args string) int {
   fmt.Println(args)
   return 1
}

func main() {
   handlers := map[string]func(args string) int {
     "help" : Help,
     "h" : Help,
     "play" : Play,
     "p" : Play,
     "quit" : Quit,
     "q" : Quit,
   }
   r := bufio.NewReader(os.Stdin)
   for {
   b, _, _ := r.ReadLine()
   line := string(b)
   handler, ok := handlers[line]
   if ok {
       req := handler(line)
       if req == 1 {
          os.Exit(1)
       }
   }
  }
}
