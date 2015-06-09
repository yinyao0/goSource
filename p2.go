package main

import "fmt"

type Animal struct {
     name string
     age int
}

type Bird struct {
    Animal
    color string
}

type Cat struct {
    Animal
    weight int
}

func (a Animal) Eat() {
   fmt.Println("animal eat")
}

func (a Animal) Sleep() {
  fmt.Println("animal sleep")
}

func (c Cat) Eat() {
  fmt.Println("cat eat")
} 

func (b Bird) Sing() {
  fmt.Println("bird sing")
}

type all interface {
  Sleep()
  Eat()
}

func main() {
  var am Animal
  var cat Cat
  var bird Bird
  var a all
  am.Eat()
  cat.Eat()
  bird.Eat()
  a = bird
  a.Sing()
}
