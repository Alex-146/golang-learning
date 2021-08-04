package main

import "fmt"

func main() {
	arrays()
	slices()
	maps()
}

func arrays() {
  var foo [3]int
  foo[1] = 146
  fmt.Println(foo, len(foo))

  bar := [4]int { 1, 2, 3, 4 }
  fmt.Println(bar, len(bar))

  grid := [2][2]int{}
  fmt.Println(grid)
}

func slices() {
  var foo [3]int = [3]int{ 1, 2, 3 }
  var bar []int = foo[1:2]

  fmt.Println(bar)
  fmt.Println(bar[:cap(bar)])

  var a []int = []int { 1, 2, 3 }
  a = append(a, len(a) + 1)
  fmt.Println(a)
  
  var b = make([]int, 5)

  for i, e := range(b) {
    fmt.Println(fmt.Sprintf("%v: %v", i, e))
  }

  fmt.Println(b)
}

func maps() {
  var a map[string]int = map[string]int {
    "foo": 123,
    "bar": 456,
  }
  fmt.Println(a)
  
  a["temp"] = 789
  fmt.Println(a)
  
  delete(a, "foo")
  fmt.Println(a)

  value, ok := a["temp"]
  if ok {
    fmt.Println(value)
  } else {
    fmt.Println("doesnt exist")
  }

  b := make(map[string]int)
  fmt.Println(b)
}