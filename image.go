package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fetchImageUrl(url string) []string {
	str := httpGet(url)
	//fmt.Println(str)

	imageUrls := ParseImageUrl(str)
	ret := []string{}
	for _, txt := range imageUrls {
		ret = extractImageFilename(txt)
	}
	return ret
}

func httpGet(url string) string {

	res, err := http.Get(url)
	if err != nil {
		return ""
		//panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
		// handle error
	}

	// fmt.Println(string(body))
	// fmt.Println(res.Status)

	return string(body)
}

func httpGetBinaryToFile(url string, path string) (string, bool) {

	fmt.Printf("url: %s \npath: [%s]\n", url, path)

	res, err := http.Get(url)
	if err != nil {
		return err.Error(), false
	}
	defer res.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return "", true
}
