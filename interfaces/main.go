package main

import (
	"fmt"
	"math"
	"strings"
)

//In order to successfully implement an interface,
//you need to implement all the methods declared by the interface with exact signatures.

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Empty interface
type MyString string

func explain(i interface{}) {
	fmt.Printf("Type is: %T, Value is: %v\n", i, i)
}

func explainWithTypeCheck(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("Type is String, Value is ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("Type is int, Value is ", i)
	default:
		fmt.Printf("Type is %T, Value is %v\n", i, i)
	}
}

// Multiple interfaces
type Area interface {
	Area() float64
}

type Volume interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * c.side * c.side
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

// Embedding interfaces
type Bark interface {
	sound()
}

type Bite interface {
	Bite()
}

type Animal interface {
	Bark
	Bite
}

type Dog struct {
	name string
}

func (d Dog) Sound() {
	fmt.Printf("Dog %v bark \n", d.name)
}

func (d Dog) Bite() {
	fmt.Printf("Dog %v bite \n", d.name)
}

// Pointer to interface
type CollegeFee struct {
	admissionFee int
	yearlyFee    int
}

type Student interface {
	studentDetail()
}

func (c *CollegeFee) studentDetail() {
	fmt.Printf("Admission Fee: %v, Yearly Fee: %v\n", c.admissionFee, c.yearlyFee)
}

func main() {
	//Here as rectangle is implementing Shape interface, so we can use Rectangle variable to call interface methods directly.
	rectangle := Rectangle{7.0, 5.0}
	fmt.Printf("Rectangle Area: %v, Perimeter: %v\n", rectangle.Area(), rectangle.Perimeter())

	//When a type implements an interface, a variable of that type can also be represented as the type of an interface.
	//We can confirm that by creating a nil interface s of type Shape and assign a struct of type Rect.
	var s Shape
	s = Rectangle{5.0, 4.0}
	fmt.Println("Rectangle Area ", s.Area())
	fmt.Println("Rectangle Perimeter ", s.Perimeter())

	var circleShape Shape
	circleShape = Circle{5.0}
	fmt.Println("Circle Area ", circleShape.Area())
	fmt.Println("Circle Perimeter ", circleShape.Perimeter())

	//Empty interface
	ms := MyString("Hello From My String")
	explain(s)
	explain(ms)

	//Multiple interfaces
	//A type can implement multiple interfaces.
	c := Cube{3}
	var area Area = c
	var volume Volume = c
	fmt.Printf("Area fetched using struct reference: %v, Volume fetched using struct reference: %v\n", c.Area(), c.Volume())
	fmt.Println("Cube Area is ", area.Area())
	fmt.Println("Cube Volume is ", volume.Volume())

	//Type assertion
	cube := area.(Cube)
	fmt.Printf("Area fetched using Type assertion: %v, Volume fetched using Type assertion: %v\n", cube.Area(), cube.Volume())

	//Check if dynamic type of i implements interface Type?
	_, ok := area.(Shape)
	fmt.Println(ok)

	explainWithTypeCheck("Hello World")
	explainWithTypeCheck(cube)
	explainWithTypeCheck(52)

	//Embedding interfaces
	//In Go, an interface cannot implement other interfaces or extend them,
	//but we can create a new interface by merging two or more interfaces.
	d := Dog{"Puppy"}
	d.Sound()
	d.Bite()

	//Pointer to interface
	collegeFee := CollegeFee{100, 200}
	var stu Student = &collegeFee
	stu.studentDetail()
}
