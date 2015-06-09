
package main

import (
   "fmt"
   "github.com/garyburd/redigo/redis"
)

func main() {
   c, err := redis.Dial("tcp",":6379")
   if err != nil {
      panic(err)
   }
   defer c.Close()

  c.Send("multi")
  c.Send("set","k1","aa")
  c.Send("set","k2","bb")
  c.Send("set","k3","cc")
  c.Send("mset","k4","dd","k5","ee","k6","ff")
  r, err := c.Do("exec")
  fmt.Println(r)
}
