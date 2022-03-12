package main

import (
	"fmt"
	"regexp"
)

var (
	// 非贪婪模式
	reg  = regexp.MustCompile(`background-image: url\(.*\/th.*?\)`)
	reg2 = regexp.MustCompile(`=.*?&`)
)

func ParseImageUrl(str string) []string {

	imageUrls := reg.FindAllString(str, -1)
	length := len(imageUrls)
	if length == 0 {
		fmt.Println(length)
	}
	return imageUrls
}

func extractImageFilename(str string) []string {

	ret := []string{}
	//fmt.Println(str)
	for _, param := range reg2.FindStringSubmatch(str) {
		ret = append(ret, param[1:len(param)-1])
		continue
	}
	return ret
}
