package main

import (
    "fmt"
    "errors"
    "sync"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
    urls map[string]error
    sync.Mutex
}{urls: make(map[string]error)}

var loading = errors.New("loading")

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    fetched.Lock()
    if _, ok := fetched.urls[url]; ok {
        fetched.Unlock()
        fmt.Println("already fetched: ", url)
        return
    }
    fetched.urls[url] = loading
    fetched.Unlock()

    body, urls, err := fetcher.Fetch(url)
    fetched.Lock()
    fetched.urls[url] = err
    fetched.Unlock()

    if err != nil {
        fmt.Printf("%v failed: %v\n", url, err)
        return
    }

    done := make(chan bool)
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        fmt.Printf("Crawling child of %v: %v\n", url, u)
        go func(url string) {
            Crawl(url, depth-1, fetcher)
            done<-true
        }(u)
    }
    fmt.Printf("Done with %v\n", url)
    for _, u := range urls {
        fmt.Printf("Waiting for %v\n", u)
        <-done
    }
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)

    fmt.Println("\nFetching")
    for url, err := range fetched.urls {
        if err != nil {
            fmt.Printf("error fetching: %v\n", err)
        } else {
            fmt.Printf("%v fetched\n", url)
        }
    }
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
