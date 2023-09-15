package main

import (
	"fmt"
)
func main() {
    v1 := [3]float64{7, 5, 4}
    var v2 [3]float64
    v2 = [3]float64{2, 4, 6}
    for v3,i := [...]float64{0, 0, 0}, 0; i < len(v3); i++ {
        v3[i] = v1[(i + 1) % 3] * v2[(i + 2) % 3] - v1[(i + 2) % 3] * v2[(i + 1) % 3]
        defer fmt.Printf("%v\n", v3)
    }
}
