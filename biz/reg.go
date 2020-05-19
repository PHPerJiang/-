package biz

import (
	"regexp"

	"github.com/my/repo/InformationCollection/dao"
)

func init() {

}

// RegUrl 正则匹配链接
func RegUrl(body string) (urls []string) {
	reg := `<a href="(/topics/\d+)"`
	regObj := regexp.MustCompile(reg)
	result := regObj.FindAllStringSubmatch(body, -1)
	for _, v := range result {
		urls = append(urls, "https://studygolang.com"+v[1])
	}
	return
}

func RegTitles(body string) (titles []dao.Titles, err error) {
	reg := `<a href="(/topics/\d+)"`
	regObj := regexp.MustCompile(reg)
	_ = regObj.FindAllStringSubmatch(body, -1)
	return
}
