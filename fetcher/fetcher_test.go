package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	url := " http://album.zhenai.com/u/1141336149"
	fetch, err := Fetch(url)
	if err != nil {
		fmt.Printf("fetch url err: %v\n", err)
	}
	if len(fetch) == 0 {
		t.Errorf("fetch url no content")
	}
}
