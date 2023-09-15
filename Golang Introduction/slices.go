package main

import "fmt"

func main() {
	vectors := []struct {
		x, y, z float64
	}{
		{1, 2, 3},
		{3.2, 4, 6},
		{4, 3, 1},
	}
	fmt.Printf("type %T and value %v\n", vectors, vectors)

	vectors = append(vectors, struct{ x, y, z float64 }{7, 7, 7})

	fmt.Printf("type %T and value %v\n", vectors[3:], vectors[3:])
	fmt.Printf("type %T and value %v\n", vectors[3], vectors[3])

	// print slices[index] value
	for i, v := range vectors {
		fmt.Println(i, " : ", v)
	}

	numbers := make([]int, 10, 10) // create a slice with an underlying array
	fmt.Println(numbers)

	// numbers[index] = index
	for i := range numbers {
		numbers[i] = i
	}
	// print numbers
	fmt.Println(numbers)
}

// Out put
// type []struct { x float64; y float64; z float64 } and value [{1 2 3} {3.2 4 6} {4 3 1}]
// type []struct { x float64; y float64; z float64 } and value [{7 7 7}]
// type struct { x float64; y float64; z float64 } and value {7 7 7}
// 0  :  {1 2 3}
// 1  :  {3.2 4 6}
// 2  :  {4 3 1}
// 3  :  {7 7 7}
// [0 0 0 0 0 0 0 0 0 0]
// [0 1 2 3 4 5 6 7 8 9]
