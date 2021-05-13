package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Fast vs slow url", func(t *testing.T) {
		t.Helper()

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not except an errob tu got one %v", err)
		}

		if got != want {
			fmt.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within specific time", func(t *testing.T) {
		t.Helper()

		server := makeDelayedServer(25000 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			fmt.Errorf("Expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
