package main

import (
	"errors"
	"fmt"
)
func div(a, b float64) (float64, error) {
    if b == 0 {
        return -1, errors.New("can't perform division by zero")
    }
    return a / b, nil
}
func main() {
    for i := 2.; i >= -2; i = i - 0.5 {
        if x, err := div(3, i); err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("3/%v=%v\n", i, x)
        }
    }
}
