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

func Fetch(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("wrong http request: %s", err.Error())
		return nil, fmt.Errorf("wrong http request: %s", err.Error())
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	//request.Header.Set("cookie", "hcbm_tk=MjZvkSuzEcIn8AOHsM9c12k0gXxOgSNsOhW1lbGIMPPjFIx6Qq0Aijmj8pCIlEkRisx7emipV/qMD5ZgK0ebvcKe8Vvl3Mm/HBB6RCQFAbEwP9AvMplbY1UDmWlSB3jx7dk3cCHS9bO4Jkp/tD57ezjp0wpOK+ZNGx/uNT3SW2JbwCLEX/ht/rDbmXiZbXSsAY8fXdWTgpsbwY4veOJZBQxUTywQapA7JS9Sohcdlj8COoywhRK/jwC4hqe5fC2NJgY4+o+VZa0ee3HG002+CuGM6HXTpwqV80iosVBuKsZqGi9eTY260KahUKqGozyZR83h9Btwg6XrWX5nJ43753rJ/HpaQGViC0RcUwSl3XLe+qaz0qV2w8+B0PKqnLj4Ez97x82MpZG+PD5q5P3mW+7INQ2v2cx+u12jpPe6bBPZ_RERXV4AmIa4VjY7e; FSSBBIl1UgzbN7NO=5GtCu6XVZWt_7MV.TyBkrhuTnPylYigdUm9oppwdmmEZpw1gcqD2gP1CF6jBtWf5h6vCYsPNZvh_Op_E1WGvouG; _efmdata=1MqNEMq5%2FXluNc%2FrwcaVK1WsQWPSXmiUTyPpd04bn4nv2sRb8LtEHeUm4kKSFTXdOXiA363g%2Fp8PY4vpycfPu28nvJz9iW%2B2EjRwVcn9iCQ%3D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1660298194; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1660298194; sid=dfeffb4b-827f-40e9-80a5-ced6210fafea; FSSBBIl1UgzbN7NP=53B3OCKhCEelqqqDc9USUBqQphL.GFwikcOKSiNEumfkMIkfIhFbB0eIZ2ytlzr81NdBUgVhIqkctZSY0cIO694Sh4ZsAKMD9roaXQ1c4wA9mJJZGbE1seIDGgU8TgnTdm5rJxIQJDchxpv9_fMrbf1wIk7QIMvs0CZkB6VqEaC5pUxtB0Zd881EN2wIbvByzkDd7LrUx4n8onxHBKcyreGEpiLNa5wlfdqapL5g5vLER6fqSg2xjLTGdRuMQ9L35o8.H9ldQcWSSKJGriwQj59QetRk7ogQ5Gsfr.A5GYVbNy9L8a14uvmpe7_n49Wo1E")
	request.Header.Set("x-client-data", "CIu2yQEIpbbJAQjBtskBCKmdygEIndPKAQiWocsBCKC8zAEI+rzMAQizwcwBCMXBzAEI1sHMAQ==")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	//bodyReader := bufio.NewReader(resp.Body)
	//e := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	//return ioutil.ReadAll(utf8Reader)
	return ioutil.ReadAll(resp.Body)
}

func Fetch1(url string) ([]byte, error) {
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
