package _select

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayeyServer(20 * time.Millisecond)
	fastServer := makeDelayeyServer(0)

	t.Run("race", func(t *testing.T) {
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		assert.NoError(t, err)
		assert.Equal(t, want, got)

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("returns error if a serve doesn't exist", func(t *testing.T) {
		server := makeDelayeyServer(11 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10 * time.Millisecond)

		assert.EqualError(t, err, ErrTimeOut.Error())
	})
}

func makeDelayeyServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
