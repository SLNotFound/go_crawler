package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

//func Fetch(url string) ([]byte, error) {
//	client := http.Client{}
//	request, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		fmt.Printf("wrong http request: %s", err.Error())
//		return nil, fmt.Errorf("wrong http request: %s", err.Error())
//	}
//
//	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
//
//	resp, err := client.Do(request)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
//	}
//
//	//bodyReader := bufio.NewReader(resp.Body)
//	//e := determineEncoding(bodyReader)
//	//utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
//	//return ioutil.ReadAll(utf8Reader)
//	return ioutil.ReadAll(resp.Body)
//}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 根据页面的charset 自动转换中文
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
