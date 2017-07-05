package methods_interfaces

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y, Abs float64
	Z         string
}

//Method
//A method is a function with a receiver, "(t Test)" in this example.
//The receiver indicates the variable to inject the method into as a field
//The receiving struct/object is the scope of the function
func (v *Vertex) Calc() float64 {
	return v.X * v.Y
}

//Methods can have a pointer to as a receiver
//A receiver is pass by value, which means that operations done are done on a copy
//of the receiving object. Using the pointer as the receiver value allows the function
//to modify the original receiving object.
func (v *Vertex) Scale(scalar float64) {
	v.X = v.X * scalar
	v.Y = v.Y * scalar
}

type Float float64

//A method does not have to be linked to a struct. It can also be linked to a map/list
//Or even a type of a basic type
//The key is the "type" keyword.
func (f Float) ToInt() int {
	return int(f)
}

//This takes in a pointer to the original object instead of making and operating on
//a copy.
func Abs(v *Vertex) float64 {
	v.Abs = math.Sqrt(v.X*v.X + v.Y*v.Y)
	return v.Abs
}

//Interface
type VertInterface interface {
	Calc() float64
}

func RunCalc(v VertInterface) float64 {
	return v.Calc()
}

type EmptyInterface interface{} //Everything implements this.

//Type Switch
func ShowType(i interface{}) {
	//You can test the type of anything using a type switch, which is a special type of switch
	switch v := i.(type) { //Note the keyword 'type' instead of the concrete type.
	case int:
		fmt.Printf("Thrice %v is %v\n", v, v*3)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Not sure what this is\n")
	}
}

//Stringers
//Importance of the string interface
//Any type that implements func String() string implements the "Stringer" interface
//This is one of the most used interfaces. fmt looks for it, for example.
func (v Vertex) String() string {
	return fmt.Sprintf("%v, %v are the coordinates of this Vertex", v.X, v.Y)
}

//Errors
//Importance of the error interface
//Any type that implements func Error() string implements the Error interface
//This means that it IS an error, from Go's perspective
//This is one of the most used interfaces.
//The keyword "error" can be thought of as a name of an typed interface that contains this method:
// func Error() String
type VertexError struct {
	x, y float64
}

func (v VertexError) Error() string {
	return fmt.Sprintf("%v, %v could not do that\n", v.x, v.x)
}

func (v VertexError) String() string {
	return fmt.Sprintf("%v, %v could not do that\n", v.x, v.x)
}

func CouldError(x int) (int, error) {
	if x == 10 {
		return -1, VertexError{10, 5}
	}

	return x + 10, nil
}
