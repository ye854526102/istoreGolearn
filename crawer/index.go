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
	"regexp"
)

func main() {
	if resp, err := http.Get("http://www.zhenai.com/zhenghun"); err != nil {
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
				//fmt.Printf("%s\n", content)
				getHtmlAllCityInfo(content)
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

func getHtmlAllCityInfo(str []byte) {
	regexpObj := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[A-Za-z0-9]+)"[^>]*>([^<]+)</a>`)
	res := regexpObj.FindAllSubmatch(str, -1)
	for _, v := range res {
		fmt.Printf("地址是:%s 城市名称是:%s\n", v[1], v[2])
	}
}
