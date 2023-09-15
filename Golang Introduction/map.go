package main

import "fmt"
func main() {
    age := map[string]int {"max": 24, "tom": 28}
    fmt.Println("map:", age)
    m := make(map[string]float64)
    m["E"]  = 2.7182818284
    m["Pi"] = 3.1415926535
    m["Phi"]= 1.6180339887

    for key, v := range m {
        fmt.Printf("Key: %v, Value: %v, Value: %v \n", key, v, m[key])
    }
    delete(m, "E") // does not return anything. It does nothing, if the key does not exist.
    fmt.Println("len:", len(m))
    fmt.Println("map:", m)

    _, ok := m["E"] // does the key exists?
    fmt.Println("ok:", ok)
}

// Output
// map: map[max:24 tom:28]
// Key: E, Value: 2.7182818284, Value: 2.7182818284 
// Key: Pi, Value: 3.1415926535, Value: 3.1415926535 
// Key: Phi, Value: 1.6180339887, Value: 1.6180339887 
// len: 2
// map: map[Phi:1.6180339887 Pi:3.1415926535]
// ok: false