package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	if resp, err := http.Get("http://www.zol.com"); err != nil {
		panic("sssssss")
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			htmlDetermineEncoding := getHtmlDetermineEncoding(resp.Body)
			//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
			utf8Reader := transform.NewReader(resp.Body, htmlDetermineEncoding.NewDecoder())
			if content, err := ioutil.ReadAll(utf8Reader); err != nil {
				panic("sssssss")
			} else {
				fmt.Printf("%s \n", content)
			}
		}
	}

}

func getHtmlDetermineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
