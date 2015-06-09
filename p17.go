package main

import (
    "fmt"
)

const toLower = 'a' - 'A'

func Dostring(s string) string {
   upper := true
   for i := 0; i < len(s); i++ {
     c := s[i]
     if upper && 'a' <= c && c <= 'z' {
        return dostring([]byte(s))
     }
     if !upper && 'A' <= c && c <= 'Z' {
        return dostring([]byte(s))
     }
     upper = c == '-'
   }
   return s
}


func dostring(a []byte) string {
upper := true
for i, c := range a {
// Canonicalize: first letter upper case
// and upper case after each dash.
// (Host, User-Agent, If-Modified-Since).
// MIME headers are ASCII only, so no Unicode issues.
if c == ' ' {
c = '-'
} else if upper && 'a' <= c && c <= 'z' {
c -= toLower
} else if !upper && 'A' <= c && c <= 'Z' {
c += toLower
}
a[i] = c
upper = c == '-' // for next time
}
// The compiler recognizes m[string(byteSlice)] as a special
// case, so a copy of a's bytes into a new string does not
// happen in this map lookup:
return string(a)
}

func main() {
   s := "accept-encoding"
   //ss := Dostring(s)
   ss := dostring([]byte(s))
   fmt.Println(s)
   fmt.Println(ss)
   fmt.Printf("%s\n",ss)
}
