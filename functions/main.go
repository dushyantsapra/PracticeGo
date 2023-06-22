package main

import "fmt"

type calcFunc func(int, int) int

func main() {
	fmt.Println("add ", add(1, 2))
	fmt.Println("sub ", sub(2, 1))
	v1, v2 := increment(1, 2)
	fmt.Println("increment ", v1, " ", v2)
	a, m := addMul(1, 2)
	fmt.Println("addMul ", a, " ", m)
	deferExample1()
	deferExample2()
	fmt.Println("funcAsType with sub with func type ", calc(1, 2, sub))
	fmt.Println("funcAsType with add ", funcAsType(2, 2, add))
	fmt.Println("Anonymous function ", addition(2, 2))
	fmt.Println("Anonymous function ", subtraction)
}

type CalcFunc func(int, int) int

// Return value
func add(a int, b int) int {
	return a + b
}

// Return value
func sub(a int, b int) int {
	return a - b
}

// Multiple return values
func increment(a int, b int) (int, int) {
	a += 1
	b += 1
	return a, b
}

// Named return values
func addMul(a int, b int) (add int, mul int) {
	add = a + b
	mul = a * b
	return
}

func deferFunction(value string) {
	fmt.Println("String is", value)
}

// defer keyword
// Add defered function call are stored in the stack with argument value,
// So multiple defered function call in function would happen in LIFO mode,
// so last defer call would be executed first.
func deferExample1() {
	name := "Dushyant"

	fmt.Println("Name ", name)

	//This is will Print "Dushyant",
	//not sapra as it was pushed into the stack with all available argument values
	defer deferFunction(name)

	name = "Sapra"

	fmt.Println("New Name ", name)
}

// defer keyword
func deferExample2() {
	deferFunction("Hello 1")
	defer deferFunction("Bye 1")

	deferFunction("Hello 2")
	defer deferFunction("Bye 2")

	deferFunction("Hello 3")
	defer deferFunction("Bye 3")
}

// Function as type
func funcAsType(a int, b int, f func(int, int) int) int {
	r := f(a, b)
	return r
}

func calc(a int, b int, f calcFunc) int {
	r := f(a, b)
	return r
}

// anonymous function
var addition = func(a int, b int) int {
	return a + b
}

var subtraction = func(a int, b int) int {
	return a - b
}(5, 3)
