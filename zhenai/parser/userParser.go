package parser

import (
	"go_crawler/model"
	"regexp"
)

const userUrlReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseUserList1(contents []byte) model.User {
	re := regexp.MustCompile(userUrlReg)
	matches := re.FindAllSubmatch(contents, -1)

	user := model.User{}

	for _, match := range matches {
		user.UserUrl = append(user.UserUrl, string(match[1]))
		user.UserName = append(user.UserName, string(match[2]))
	}

	//for _, v := range user.UserUrl {
	//	fmt.Println(v)
	//}

	return user
}
