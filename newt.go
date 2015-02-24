/* author: Mark Silvis
 * Uses Newton's method to estimate the square root of a number
*/

package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

func Newt(x, z float64) float64 {
    return z-((math.Pow(z, 2)-x)/(2*z))
}

func main() {
    var num float64 = 121           // find sqrt estimate for num
    var delta float64 = 0.0000005   // delta
        // when calculated value changes by less than
        // this amount between loops, return the value
    var z float64 = 100 // starting point
    
    fmt.Printf("Square root: %g\n", Sqrt(num))  // print exact sqrt
    
    prev := 0.0
    for i := 0; ; i++ {
        z = Newt(num, z)
        if prev == 0 {
            prev = z
        } else if (prev-z) < delta {
            fmt.Printf("Estimate: %g\n", z)
            fmt.Printf("Found in %d loops\n", i)
            break
        } else {
            prev = z
        }
        
        // Optional
        // Stops calculation after 100 loops
        /*if i > 100 {
            fmt.Println(i)
            break
        }*/
    }
}
