package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    count := make(map[string]int)
    fields := strings.Fields(s)
    for _,v := range fields {
        count[v]++
    }
    return count
}

func main() {
    wc.Test(WordCount)
}
