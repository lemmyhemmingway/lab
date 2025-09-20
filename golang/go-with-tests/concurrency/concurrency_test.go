package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockWebsiteChecker(url string) bool {
	return url != "waat://gundi.goods"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://example.com",
		"waat://gundi.goods",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://example.com": true,
		"waat://gundi.goods": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}

func BenchmarkCheckWebhasites(b *testing.B) {
	urls := make([]string, 100)

	for i:= 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
