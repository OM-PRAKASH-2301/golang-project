package main

import "fmt"

func main() {
	// integer()
	// float()
	// String()
	// Boolean()
	// Byte()
	// Array()
	// Slice()
	// Slice()
	// Constructcall()
	// mapcall()
	interfaceCall()
}

//Integer code started here
func integer() {
	var a int = 100                   // Platform-dependent size
	var b int8 = 127                  // 8-bit signed integer
	var c int16 = 32767               // 16-bit signed integer
	var d int32 = 2147483647          // 32-bit signed integer
	var e int64 = 9223372036854775807 // 64-bit signed integer

	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int16:", c)
	fmt.Println("int32:", d)
	fmt.Println("int64:", e)
}

// integer code ended here

// float code started here
func float() {
	var c float32 = 10.5455456546664554334
	var d float64 = 20.12345
	fmt.Println("Float values:", c, d)
}

// float code ended here

//String code started here
func String() {
	var name string = "Om Prakash"
	fmt.Println("String value:", name)
}

// String code ended here

// Boolean code started here
func Boolean() {
	var isGoFun bool = false
	fmt.Println("Boolean value:", isGoFun)
	if isGoFun != true {
		isGoFun = true
	}
	fmt.Println("Boolean value after condition check:", isGoFun)
}

// Boolean code ended here

// Byte code started here
func Byte() {
	var a byte = 'A' // ASCII value of 'A' is 65
	var b byte = 'B'
	var c byte = 'C'
	var y byte = 'Y'
	var z byte = 'Z'
	fmt.Println("Byte value:", a, b, c, y, z)
}

// Byte code ended here

// Array code started here
// Array => An array in Go is a collection of elements of the same type, stored in contiguous memory locations. Arrays have a fixed size, which must be defined at the time of declaration. Once defined, the size of an array cannot be changed.
func Array() {
	// Example 1: Basic Array
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Example 1 - Basic Array:", arr1)

	// Example 2: Default Values
	var arr2 [3]int // Uninitialized elements will be 0
	fmt.Println("Example 2 - Default Values:", arr2)

	// Example 3: Partial Initialization
	arr3 := [5]int{10, 20}
	fmt.Println("Example 3 - Partial Initialization:", arr3)

	// Example 4: Implicit Length
	arr4 := [...]int{100, 200, 300}
	fmt.Println("Example 4 - Implicit Length:", arr4)

	// Example 5: Accessing Elements
	arr5 := [3]string{"Go", "is", "fun"}
	fmt.Println("Example 5 - First Element:", arr5[0])
	fmt.Println("Example 5 - Last Element:", arr5[2])

	// Example 6: Modifying Elements
	arr6 := [3]int{1, 2, 3}
	arr6[1] = 42
	fmt.Println("Example 6 - Modified Array:", arr6)

	// Example 7: Iterating Over an Array
	arr7 := [4]string{"A", "B", "C", "D"}
	fmt.Println("Example 7 - Iterating:")
	for i, v := range arr7 {
		fmt.Printf("Index %d: %s\n", i, v)
	}
}

// Array code ended here

// Slice code started here
func Slice() {
	var slice []int = []int{1, 2, 3, 4, 5, 7, 8}
	slice = append(slice, 4) // Add an element
	fmt.Println("Slice values:", slice)
}

// Slice code ended here

// struct code started here
type Person struct {
	Name string
	Age  int
}

func Constructcall() {
	var p Person = Person{Name: "Om Prakash", Age: 31}
	fmt.Println("Struct value:", p)
}

// struct code ended here

// interface code started here

// What is an Interface?
// An interface in Go is a type that specifies a collection of method signatures. If a type implements all the methods declared in an interface, the type is said to "satisfy" the interface.
func interfaceCall() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	fmt.Println("Rectangle Details:")
	PrintShapeDetails(r)

	fmt.Println("\nCircle Details:")
	PrintShapeDetails(c)
}

func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

// interface code ended here

// map code started here
func mapcall() {
	students := map[int]string{
		101: "Om Prakash",
		102: "Amit Kumar",
		103: "Rahul Sharma",
	}

	// Map को iterate करना
	for id, name := range students {
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
