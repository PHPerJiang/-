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

// RegTitles 正则匹配内容
func RegTitles(body string) (titles []*dao.Titles) {
	reg := `[\d]+、[\s\S]+?\[([\s\S]+?)\]\(([\s\S]+?)\)`
	regObj := regexp.MustCompile(reg)
	result := regObj.FindAllStringSubmatch(body, -1)
	for _, v := range result {
		titles = append(titles, &dao.Titles{
			Title: v[1],
			Link:  v[2],
		})
	}
	return
}
