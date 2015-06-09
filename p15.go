//匿名函数
package main

import (
   "fmt"
)

func main() {
    x,y := func (i,j int)(m,n int) {
        return j,i
    }(1,9)
   fmt.Println(x,y)
}
