package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Teste
func Titulo(urls ...string) <-chan string {
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
