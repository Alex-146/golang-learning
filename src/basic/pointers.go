package main

import "fmt"

func main() {
	// one()
	// two()
	three()
}

func one() {
	a := 42
	fmt.Println("address is:", &a)

	fmt.Println("before:", a)
	b := &a
	*b = 146
	fmt.Println("after:", a)

	fmt.Println(fmt.Sprintf("%T", b)) // pointer - *int
}

func two() {
	a := 42
	fmt.Println("before:", a)
	change(&a)
	fmt.Println("after:", a)
}

func change(value *int) {
	*value = 146
}

func three() {
	value := 1
	var pointer *int = &value
	*pointer += 1
	fmt.Println(value, *pointer)
}