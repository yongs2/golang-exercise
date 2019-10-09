package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
    // a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 함수를 고쳐서, 같은 URL을 두번 가져오는 중복을 피하면서 URL들을 병렬로 패치하게 고쳐보세요.
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
    // TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("[", depth, "]", err)
		return
	}
	fmt.Printf("[ %d ] Crawl.found: %s %q\n", depth, url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	fmt.Println("Fetch: ", url, ", f[url]=", (*f)[url])
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("Fetch.not f: %s", url)
}

// fetcher is a populated fakeFetcher
var fetcher = &fakeFetcher{
	"http://golang.org/" : &fakeResult {
		"The Go Progamming Language",
		[]string {
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/" : &fakeResult {
		"Packages",
		[]string {
			"http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string {
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
	},
	"http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string {
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}