package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func title(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := title("https://google.com", "https://amazon.com")
	t2 := title("https://youtube.com", "https://linkedin.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
