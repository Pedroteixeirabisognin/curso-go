package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Google I/O - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			//Duas formas de fazer
			r, _ := regexp.Compile(`<title>(.*?)</title>`)
			//r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) //É como se você estivesse já mandando o parâmetro pra função anônima
		//Como se fosse um func("https://www.google.com")
	}
	return c
}
func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://jovemnerd.com.br/")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
