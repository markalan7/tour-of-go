package main

import (
    "log"
    "net/http"
    "fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct string
    Who string
}

func (s Struct) String() string {
    return fmt.Sprintf("%s%s %s", s.Greeting, s.Punct, s.Who)
}

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed know."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
