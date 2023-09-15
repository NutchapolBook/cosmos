package main

// %v for a value, which will be converted into a string with default options.
// %T for the type of a value
// %x for the hex encoding
// %d for integer
// %f for float, %e and %E for scientific notation
// %s for string
// %p for the pointer address of the variable

import "fmt"
func main() {
    a, b := 2, 3
    c := float64(a + b)
    fmt.Printf("%v + %v = %f = %v, stored as %T", a, b, c, c, c)
}

// Output: 2 + 3 = 5.000000 = 5, stored as float64