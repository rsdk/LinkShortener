package models

import (
	"github.com/bluele/gcache"
)

var gc = gcache.New(10000).ARC().Build()

func SaveUrl(url string) uint {
	var link Link
	link.Url = url
	DB.Create(&link)
	return link.ID
}

func FetchUrlCached(id uint) string {
	url, err := gc.Get(id)
	if err != nil {
		url := FetchUrl(id)
		gc.Set(id, url)
		return url
	}
	urlStr, ok := url.(string)
	if !ok {
		panic("Could not convert url to string")
	}
	return urlStr
}

func FetchUrl(id uint) string {
	var link Link
	DB.First(&link, id)
	return link.Url
}
