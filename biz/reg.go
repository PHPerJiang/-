package biz

import (
	"regexp"
)

func init() {

}

// RegUrl 正则匹配链接
func RegUrl(body []byte) (urls []string) {
	reg := `<a href="(/topics/\d+)"`
	regObj := regexp.MustCompile(reg)
	result := regObj.FindAllStringSubmatch(string(body), -1)
	for _, v := range result {
		urls = append(urls, "https://studygolang.com"+v[1])
	}
	return
}

func RegContent() {

}
