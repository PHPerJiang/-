package controllers

import (
	"fmt"

	"github.com/my/repo/InformationCollection/biz"
)

func init() {

}

// Run 启动app
func Run(url string) {
	urls, err := biz.GetUrlListByUrl(url)
	if err != nil {
		fmt.Println("Get Urls Failed!")
		return
	}

	for _, v := range urls {
		fmt.Println(v)
	}
}
