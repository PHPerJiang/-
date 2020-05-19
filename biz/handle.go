package biz

import (
	"fmt"
	"strconv"

	"github.com/my/repo/InformationCollection/dao"
)

func init() {

}

// HandleGitContent 拼接提交git的正文内容
func HandleGitContent(titles []*dao.Titles) {
	content := ""
	for k, v := range titles {
		line := strconv.Itoa(k+1) + ". " + v.Title + "  " + v.Link + "\n"
		content += line
	}
	fmt.Printf("%v", content)
}
