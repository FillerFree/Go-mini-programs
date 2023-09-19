package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func main() {
	bucketMap := map[string]Bucket{
		"limited":   {capacity: 3, name: "limited", lastUsed: 0},
		"unlimited": {capacity: 3, name: "unlimited", lastUsed: 0},
	}
	http.HandleFunc("/limited", func(w http.ResponseWriter, r *http.Request) {
		rateLimiter, _ := bucketMap[r.URL.Path]
		if rateLimiter.isRateLimited(time.Now().Hour()*60 + time.Now().Minute()) {
			w.WriteHeader(503)
			fmt.Fprintf(w, "You got rate limited Noob !!, %q", html.EscapeString(r.URL.Path))
		}
		fmt.Println()
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return
		}
	})

	http.HandleFunc("/unlimited", func(w http.ResponseWriter, r *http.Request) {
		rateLimiter, _ := bucketMap[r.URL.Path]
		if rateLimiter.isRateLimited(time.Now().Hour()*60 + time.Now().Minute()) {
			fmt.Fprintf(w, "You got rate limited Noob !!, %q", html.EscapeString(r.URL.Path))
		}
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
