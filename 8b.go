package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    delta := 1e-10
    z := 1.0
    old_z := z - 2*delta
    for math.Abs(z-old_z) > delta {
        old_z = z
        z = z - (z*z-x)/(2*z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}

