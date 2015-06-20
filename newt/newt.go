package main

import (
    "fmt"
    "math"
)

func Newt(x float64) (float64, int) {
    var z, delta float64 = 100, 0.05
    var prev float64
    for i:=0; ; i++ {
        prev = z
        z = z-(math.Pow(z,2)-x)/(2*z)
        if prev - z < delta {
            return z, i
        }
    }
}

func main() {
    var eval float64 = 100
    result, loops := Newt(eval)
    fmt.Printf("Est:\t%f\n", result)
    fmt.Printf("Found in %d loops\n", loops)
    fmt.Printf("Real:\t%v", math.Sqrt(eval))
}
