package feeds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SetHost(feeds []string) {
	rssFeeds := feeds[0]
	apiFeeds := feeds[1]
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		reader := strings.NewReader(rssFeeds)
		_, err := io.Copy(writer, reader)
		ErrorHandling(err)
		req.Header.Set("Content-Type", "application/rss+xml")
		return
	})
	http.HandleFunc("/api", func(writer http.ResponseWriter, req *http.Request) {
		reader := strings.NewReader(apiFeeds)
		_, err := io.Copy(writer, reader)
		ErrorHandling(err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Access-Control-Allow-Headers", "*")
		req.Header.Set("Access-Control-Allow-Origin", "*")
		req.Header.Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
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
	http.HandleFunc("/google900b28595c041e06.html", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "google-site-verification: google900b28595c041e06.html")
		return
	})
}
