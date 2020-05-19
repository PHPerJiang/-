package controllers

import (
	"fmt"

	"github.com/my/repo/InformationCollection/biz"
)

func init() {

}

// Run 启动app
func Run(url string) {
	//获取urls
	urls, err := biz.GetUrlListByUrl(url)
	if err != nil {
		fmt.Println("Get Urls Failed!")
		return
	}
	//获取文章正文并提取标题与连接
	for _, v := range urls {
		body, _ := biz.GetBodyByUrl(v)
		titles := biz.RegTitles(body)
		for _, v := range titles {
			fmt.Printf("title:%s\n", v.Title)
			fmt.Printf("link:%s\n", v.Link)
		}
	}
}
