package basics

//This is a "factored" import and allows multiple imports with only one "import" label
import (
	"fmt"
)

//Declaring new variables
//var <name>[, <name>...] <type> = <value>[, <value>...]
var Exported string = "This is exported"

//Constant. Can not use shorthand declaration
//Only basic types can be constants (no structs, maps, functions, etc)
const Pi = 3.14
const Npi float32 = -3.14

/**
None of these variables are exported
*/
//Basic types basically amount to boolean, string, and numeric types (there are a lot of these)
var i int = 1
var b bool = true
var s string = "asdf"

//Variables that are not initialized are given their "zero value"
var ii int    //0
var bb bool   //false
var ss string //"" <- empty string

//This function takes in two ints as parameters and returns two named results
//(an int and an error, if that should be necessary)
func Add(x, y int) (z int, err error) {
	fmt.Println("Running Add")
	z = x + y
	return z, nil
}

/*
	Flow controls
*/

//For. Take in n. Return n+10 and n+20
//There is no while loop. for performs that role too
func ForFunc(n int) (x, y int, err error) {
	// for pre-statement; comparison; post-statement; {}
	// pre is executed once before the beginning of the loop
	//post is executed after every iteration
	y = n
	x = n
	for i := 0; i < 10; i++ {
		x++
		y++
	}
	for i <= 100 { //This is Go's while loop
		//IF
		if i++; i <= 20 { //Go's if condition allows a short statement before
			continue //Continue skips all following code in the nearest loop
		} else if i > 30 {
			break //break out of the nearest loop
		}
		y++
	}

	return x, y, nil
}

//Switch
func SwitchFunc(in string) (string, error) {
	switch in {
	case "1":
		return "First String", nil
	case "2":
		//asdf
		return "Second String", nil
	case "3":
		//qwer
		return "Third String", nil
	default:
		//tyui
		return "Default", nil
	}
}

//Defer
func DeferFunc() {
	//This will be evaluated after the enclosing function finishes executes
	//The parameters are evaluated immediately
	x := "Deferred"
	defer fmt.Println(x)
	x = "Not Deferred"
	fmt.Println(x)
}

/*
	Data types and objects
*/
//Pointers
var p = &ii   //'&' Returns the address of ii. Can't take pointers of constants.
var pval = *p //This returns the value at p. Known as "dereferencing"

//Structs

type Circle struct { //A struct is a typed collection of fields.
	Radius float64 //They're useful for grouping data.
	Diam   float64
}

func (c Circle) Area() float64 { //You can include a function on a type
	return c.Radius * c.Radius * 3.14
}

var c = Circle{10, 20}      //This struct literal initiates the attributes in order
var c2 = Circle{Radius: 20} //Diam is initiated to the zero value (0, in this case)
var c3 = Circle{}           //Both attributes are initiated at the zero value
var cp = &c                 //You can return the pointer to a struct like normal
//Accessing the values of the attributes of a struct when you only have a pointer
//Can be done normally, without the need for (*cp)
var c_radius = cp.Radius

//Arrays
func Arrays() {
	//Arrays can not be resized.
	var array [2]string //This declares an array. Values are the zero value
	//var array2 = [...]string{"Yoda", "Windu"} //The compiler will figure out the length of the array
	fmt.Println("Zero value: " + array[0]) //Prints ""
	array[0] = "first"
	array[1] = "second"
	fmt.Println(array) //Prints space separated values
	primes := [5]int{1, 2, 3, 5, 7}
	fmt.Println(primes)
}

//Slices
func Slices() {
	//Slices are dynamically sized arrays.
	var primes = [5]int{1, 2, 3, 5, 7} //Array
	var slice1 []int = primes[:3]      //Notice the slice is declared by not defining the length of the array in the signature
	fmt.Println(slice1)
	var slice2 []int //No need to give any content. This defaults to nil, with length 0 and no content
	fmt.Println("len", len(slice2), "cap", cap(slice2))
	slice2 = primes[1:4]                                // [x:y] x defaults to 0, y to the length of the slice/array
	fmt.Println("len", len(slice2), "cap", cap(slice2)) //Length is the number of items in the slice. Capacity is the size of the underlying array
	slice2 = slice2[:2]                                 //Modify the LENGTH (not capacity) by reslicing. Must have capacity for this
	fmt.Println("len", len(slice2), "cap", cap(slice2)) //Length is 2, capacity remains at 4
	slice2 = append(slice2, 11, 13, 17, 19, 23)         //Append will grow an array when needed.
	fmt.Println("len", len(slice2), "cap", cap(slice2))
	slice2copy := make([]int, 20) //Allocates a zeroed array
	copy(slice2copy, slice2)      //copy(dst, src). Controls for the smallest index.
	fmt.Println("len", len(slice2copy), "cap", cap(slice2copy), slice2copy)
}

//Range
func Range() {
	var primes = [5]int{1, 2, 3, 5, 7}

	for index, value := range primes { //Iterates over a slice, map, or array
		fmt.Println(index, value)
	}
}

//Maps
func Maps() {

	map1 := make(map[string]Circle) //Returns an initialized map. This is not nil.
	// var map2 map[string]Circle      //This is a nil map. Attempts to write will cause a panic
	// map2["circle1"] = Circle{10, 20} //This will cause a panic
	map1["circle1"] = Circle{10, 20}           //Add
	map1["circle2"] = Circle{20, 40}           //
	delete(map1, "circle2")                    //Delete
	fmt.Println("Circle 1: ", map1["circle1"]) //Retrieve
	fmt.Println("Circle 2: ", map1["circle2"]) //Will return the zero value for the element if that key DNE
	value, ok := map1["circle2"]               //Returns value and true if exists. Zero value and false if not
	fmt.Println(value, ok)

	literal := map[string]Circle{
		"small circle": Circle{
			10,
			20,
		},
		"big circle": { //If the value is just a type name, you do not have to include it
			100,
			200,
		},
	}

	fmt.Println(literal)

}

//Functions as values
func TakesValue(fn func(int, int) int) {
	fmt.Println(fn(10, 20))
}

var AsValue = func(x, y int) int {
	return x * y
}

//Closures
func Closures() func() int {
	x, y := 10, 4
	fn := func() int {
		return x*3 + y //x and y are available in the returned function, even if Closure has finished executing
	}
	x = 13 //The value of referenced variables will be the last value assigned before referencing.
	//Thus, if several closures reference the same variable/object, it can be used to communicate between them
	return fn
}
