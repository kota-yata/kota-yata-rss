package feeds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SetHost(rssFeeds string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		reader := strings.NewReader(rssFeeds)
		_, err := io.Copy(writer, reader)
		ErrorHandling(err)
		req.Header.Set("Content-Type", "application/rss+xml")
		return
	})
	uploadCertChallenge()
	portName := os.Getenv("PORT")
	if portName == "" {
		portName = "3432"
	}
	fmt.Println("RSS feed has been published at http://localhost:" + portName)
	err := http.ListenAndServe(":" + portName, nil)
	ErrorHandling(err)
}

func uploadCertChallenge() {
	http.HandleFunc("/.well-known/acme-challenge/o7AhrXD3WFtRBrHpZdQEiJCVoEuSZ6nJcqKkgA6DwCo", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "o7AhrXD3WFtRBrHpZdQEiJCVoEuSZ6nJcqKkgA6DwCo.cNJeOkhi1FI7f7y1QEqdrZvDa29r3dp3Fz7mQrtdr2E")
		return
	})
}
