package main

// if you want to import one package
// import "fmt"

// importing multiple packages
import (
	"errors"
	"fmt"
	"math"
)

// struct is like an object in GO
// unforunately, you cannot define default value
type person struct {
	name  string
	age   int
	funny bool
}

// main func, similar to one in java
func main() {
	// define a var x (default value = 0)
	var x int
	// you could initialize it with some val
	var y int = 10

	// when you initialize, you don't have to put type
	// go can infer the type
	var z = 10

	// short hand syntax
	zz := 2

	zzz := z + zz

	// printing
	fmt.Println("x+y:", x+y)

	// if statements
	if zzz > 19 {
		fmt.Println("zzz:", zzz)
	}

	// defining arrays (fixed length)
	var arr [5]int
	fmt.Println(arr)

	// shortcut for defining array with values
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Arr2:", arr2)

	// to define array without length, we use "slices"
	// think of it as arraylist in java
	slices := []int{1, 2}
	fmt.Println("slices:", slices)

	// to add an item to a slice
	slices = append(slices, 12)
	fmt.Println("updated slices:", slices)

	// using maps
	// string = key; int = value
	table := make(map[string]int)

	// setting key-value pair
	table["age"] = 12
	table["value"] = 3

	fmt.Println("Table:", table)

	// to access specific key
	fmt.Println("Age:", table["age"])

	// to delete a key
	delete(table, "value")
	fmt.Println("Table:", table)

	// go only has for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value at i:", i)
	}

	fmt.Println("-------------")

	j := 0
	// could also function as a while loop
	for j < 4 {
		fmt.Println("Value at j:", j)
		j++
	}

	// looping over array
	stringArr := []string{"Keshav", "Narula"}

	// for looping, special variables
	// index and value using the range
	for index, value := range stringArr {
		fmt.Println("i:", index, "val:", value)
	}

	// can do the same with map
	// but instead of index, you use "key"
	for key, value := range table {
		fmt.Println("key:", key, "val:", value)
	}

	result := sum(7, 9)
	fmt.Println("Result:", result)

	// since we returned two things from the func,
	// we can extract that
	sqrtVal, err := sqrt(64)

	if err != nil {
		fmt.Println("sqrt(64):", err)
	} else {
		fmt.Println("sqrt(64):", sqrtVal)
	}

	// using structs
	// you can instantiate them like classes
	p := person{name: "Keshav", age: 21}
	fmt.Println("Person:", p)

	// Lets talk about pointers
	pt := 6
	fmt.Println("pt VALUE:", pt)
	fmt.Println("pt POINTER:", &pt)

	inc(pt) // did not update the val
	// it made a copy
	fmt.Println("pt:", pt)

	// to allow it to modify the actual pt at that address
	// we use &.
	incValue(&pt)
	fmt.Println("pt:", pt)
}

// can define helper functions
func sum(x int, y int) int {
	return x + y
}

// you can return multiple things
// in this case a float value, followed by
// an error
// GO does NOT have expections
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Cannot get sqrt of negative numbers")
	}

	// can return nil if no error
	return math.Sqrt(x), nil
}

// increment function
// creates a copy of the variable
// does not mutate it
func inc(x int) {
	x++
}

// increment function, but actually
// modifies the value at the address
// does not make a copy
func incValue(x *int) {
	*x++
}
