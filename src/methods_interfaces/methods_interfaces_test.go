package methods_interfaces_test

import (
	"fmt"
	mi "methods_interfaces"
	"testing"
)

var test = mi.Vertex{
	10,
	6,
	10,
	"z",
}

func TestVertCalc(t *testing.T) {
	result := test.Calc()
	fmt.Println("Calc returns", result)
}

func TestVertScale(t *testing.T) {
	test.Scale(0.6)
	fmt.Println("X", test.X, "Y", test.Y)
}

func TestFloat(t *testing.T) {
	var f mi.Float = 5.78
	fmt.Println("to int of 5.78", f.ToInt())
	fmt.Println("f is still ", f)
}

func TestAbs(t *testing.T) {
	mi.Abs(&test)

	fmt.Println("Abs", test.Abs)
}

func TestInterface(t *testing.T) {
	result := mi.RunCalc(&test) //Sicne Vertex.Calc is assigned to *Vertex, we have to pass RunCalc a pointer
	fmt.Println("result", result)

	// 	var nilI mi.VertInterface //Trying to run this without assigning it to a type that implements Calc will cause a runtime error

}

func TestTypeAssertions(t *testing.T) {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) //Returns 0, false, because i's underlying type is not a float64
	fmt.Println(f, ok)

	// 	f = i.(float64) // panics because it is not TESTING whether i is a float64
	// 	fmt.Println(f)
}

func TestShowType(t *testing.T) {
	mi.ShowType(50)
	mi.ShowType("hello")
	mi.ShowType(true)
}

func TestStringer(t *testing.T) {
	fmt.Println(test)
}

func TestError(t *testing.T) {
	_, err := mi.CouldError(10)
	if err != nil {
		fmt.Printf(err.Error()) //Even though VertexError has a Stringer, the error interface is ignorant of that fact and we have to call the Error() function to get the error message
	}

}
