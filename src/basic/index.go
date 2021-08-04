package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "math"
)

func main() {
}

func part_one() {
  var user string = "Dima"
  fmt.Println(len(user))
  
  var age uint8 = 20
  if age == 18 {
    fmt.Println("congrats")
  } else if age < 18 {
    fmt.Println("not an adult")
  } else {
    fmt.Println(age)
  }
  
  const pi = 3.14159265359
  fmt.Printf("%T %.2f", pi, math.Pi)
}

func part_two() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print(">")
  scanner.Scan()
  input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
  var output = fmt.Sprintf("%T-%v", input, input)
  fmt.Println(output)
}

func part_three() {
  // ! for math operations numbers should have same type
  var a = 5
  var b = 2
  var c = float64(a) + float64(b)
  fmt.Println(c)
}

func loops() {
  i := 0
  for i < 3 {
    fmt.Println(fmt.Sprintf("i: %v", i))
    i++
  }

  for j := 0; j < 3; j++ {
    fmt.Println(fmt.Sprintf("j: %v", j))
  }
}

func switchCase() {
  value := 146
  switch value {
  case 146:
    fmt.Println("Cucumber")
  case 0:
    fmt.Println("z3r0")
  case 1, 2:
    fmt.Println("1 or 2")
  default:
    fmt.Println("d3f4ult")
  }
}