package biz

import (
	"io/ioutil"
	"net/http"
)

func init() {

}

// GetUrlListByUrl 获取链接列表
func GetUrlListByUrl(url string) (urls []string, err error) {
	resp, err := http.Get(url)
	//初始化
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	urls = RegUrl(body)
	return urls, err
}
