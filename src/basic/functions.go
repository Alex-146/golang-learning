package main

import "fmt"

func main() {
	fmt.Println(ternary(true, 1, 2))
  fmt.Println(ternary(false, 1, 2))

	t := func(x bool, a, b int) int {
		if x {
			return a
		} else {
			return b
		}
	}

	fmt.Println(t(true, 3, 4))
  fmt.Println(t(false, 3, 4))

  a, b := magic()
  fmt.Println(a, b)
  fmt.Println(magic_two())
}

func ternary(x bool, a, b int) int {
  if x {
    return a
  } else {
    return b
  }
}

func magic() (int, int) {
  return 146, 42
}

func magic_two() (a int, b int) {
  defer fmt.Println("Magic happens!")
  a = 146
  b = 42
  return
}

// func that takes func and return another func
func outer(myFunc func(int) int) func(string) int {
	return func(str string) int {
		return 146
	}
}