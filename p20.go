package main
import "fmt"
import "runtime"
import "time"

func main() {
  var i int;
  i=0
  go func() {
    for i:=0;i<10;i++ {
     fmt.Println(i);
    }
  }()
  runtime.Gosched();
  time.Sleep(time.Second);
  println(i);
}
