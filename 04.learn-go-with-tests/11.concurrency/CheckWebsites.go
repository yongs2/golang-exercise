package concurrency

import (
	//"fmt"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
		//fmt.Println(url, "=", results[url])
	}
	return results
}
