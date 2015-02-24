package main

import (
    "fmt"
    //"bytes"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
    
    // trying out the bytes package
    /*
    var buffer bytes.Buffer
    for k, v := range ip {
        if k > 0 {
            buffer.WriteString(fmt.Sprintf("."))
        }
            
        buffer.WriteString(fmt.Sprintf("%d", v))
    }
    s := buffer.String()
    return fmt.Sprintf(s)
    */
}

func main() {
    addrs := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for n, a := range addrs {
        fmt.Printf("%v: %v\n", n, a)
    }
}
