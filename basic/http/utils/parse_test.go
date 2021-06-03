package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	parse, err := url.Parse("wss://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	fmt.Print(parse)
}
