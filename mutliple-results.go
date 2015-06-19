package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func swap(x, y string) (newX, newY string) {
    newX = y
    newY = x
    return
}

func main() {
    a := "Hello"
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter name: ")
    b, _ := reader.ReadString('\n')
    b2 := strings.TrimSpace(b)
    fmt.Println(a, b2)
    a2, b3 := swap(a, b2)
    fmt.Println(a2, b3)
}
