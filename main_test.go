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

func BenchmarkSave(b *testing.B) {
	models.InitDB()
	for i := 0; i < b.N; i++ {
		models.SaveUrl("http://google.com")
	}
}

func BenchmarkLoad(b *testing.B) {
	models.InitDB()
	for i := 0; i < b.N; i++ {
		models.FetchUrl(uint((i % 100) + 1))
	}
}

func BenchmarkLoadCached(b *testing.B) {
	models.InitDB()
	for i := 0; i < b.N; i++ {
		models.FetchUrlCached(uint((i % 100) + 1))
	}
}
