package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("test CheckWebsites", func(t *testing.T) {
		t.Helper()
		websites := []string{
			"https://google.com",
			"http://glob.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		want := map[string]bool{
			"https://google.com":         true,
			"http://glob.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		got := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	b.Run("benchmark Checkwebsites", func(b *testing.B) {
		b.Helper()
		urls := make([]string, 100)
		for i := 0; i < len(urls); i++ {
			urls[i] = "a url"
		}

		for i := 0; i < b.N; i++ {
			// CheckWebsites(slowStubWebsiteChecker, urls)
			CheckWebsites(slowStubWebsiteChecker, urls)
		}
	})
}
