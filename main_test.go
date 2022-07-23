package main

import (
	"linkShortener/models"
	"testing"
)

func TestSaveLoad(t *testing.T) {
	models.InitDB()
	myId := models.SaveUrl("http://google.com")
	if myId != 1 {
		t.Log("error should be 1")
		t.Fail()
	}
	url := models.FetchUrl(myId)
	if url != "http://google.com" {
		t.Log("error should be 'url'")
		t.Fail()
	}
}
