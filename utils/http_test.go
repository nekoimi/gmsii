package utils

import (
	"fmt"
	"testing"
)

func TestParseQueryString(t *testing.T) {
	params := parseQueryString("?id=1681133829352484853&wfr=spider&for=pc")

	fmt.Println(params)

	reParse := parseToQueryString(params)

	fmt.Println(reParse)
}

func TestGet(t *testing.T) {
	if response, err := Get("http://www.baidu.com", nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(response))
	}
}
