package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	var iValue int
	iValue = 5
	fmt.Println("Hello World!", iValue)
	fmt.Println(fmt.Sprintf("Value is %v", iValue))

	var jValue int = 10
	fmt.Println(jValue)

	kValue := 15
	fmt.Println(kValue)

	lValue := 20
	sValue := "Twenty"
	fmt.Println(fmt.Sprintf(sValue+" : %v", lValue))

	s1Value := "Hello"
	s2Value := s1Value + " World From String"

	fmt.Println(s2Value)

	if lValue >= 20 {
		fmt.Println("Greater")
	} else {
		fmt.Println("Less ")
	}

	var arr [5]int
	arr[1] = 2

	fmt.Println(arr)

	arr1 := [5]int{0, 3, 1, 5, 4}
	fmt.Println(arr1)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)

	s1 := append(s, 12)
	fmt.Println(s1)

	m := make(map[int]string)
	m[0] = "Zero"
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"
	fmt.Println(m)

	delete(m, 2)
	fmt.Println(m)

	for iLoop := 0; iLoop < 5; iLoop++ {
		fmt.Println(fmt.Sprintf("Value is %v", iLoop))
	}

	fmt.Println()

	jLoop := 1

	for jLoop <= 5 {
		fmt.Println(fmt.Sprintf("jLoop is %v", jLoop))
		jLoop++
	}

	slRange := []string{"This", "is", "Loop", "Over", "Slice", "Using", "Range"}
	for index, value := range slRange {
		fmt.Println("index: ", index, " value: ", value)
	}

	mapRange := make(map[int]string)
	mapRange[1] = "Mummy"
	mapRange[2] = "Shanky"

	for key, value := range mapRange {
		fmt.Println("Key: ", key, " Value: ", value)
	}

	fmt.Println("Sum is :", sum(1, 2))

	value, err := sqrt(12.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	pArr := startStruct()
	for index, p := range pArr {
		fmt.Println("index: ", index, " Name: ", p.name, " Age: ", p.age)
	}

	x := 6
	pointerExample(&x)
	fmt.Println(x)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for -ve numbers")
	} else {
		return math.Sqrt(x), nil
	}
}

func startStruct() []person {
	p := person{name: "Dushyant", age: 35}
	p1 := person{name: "Sapra"}

	pArr := []person{p, p1}
	return pArr
}

func pointerExample(x *int) {
	*x++
}
