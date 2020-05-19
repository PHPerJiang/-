package biz

import (
	"io/ioutil"
	"net/http"
)

func init() {

}

// GetUrlListByUrl 获取链接列表
func GetUrlListByUrl(url string) (urls []string, err error) {
	body, err := GetBodyByUrl(url)
	urls = RegUrl(body)
	return urls, err
}

// GetBodyByUrl 获取响应体
func GetBodyByUrl(url string) (body string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	body = string(respBody)
	return body, err
}
