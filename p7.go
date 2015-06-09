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

   c.Do("SET","foo",1)
   exists, _ := redis.Bool(c.Do("exists","foo"))
   n, _ := redis.Int(c.Do("get","foo"))
   fmt.Println(n)
   fmt.Println(exists)
}
