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
		req.Header.Set("Content-Type", "application/rss+xml")
		writer.Header().Set("Content-Type", "application/rss+xml")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
		if (*req).Method == "OPTIONS" {
			return
		}
		reader := strings.NewReader(rssFeeds)
		_, err := io.Copy(writer, reader)
		ErrorHandling(err)
		return
	})
	http.HandleFunc("/api", func(writer http.ResponseWriter, req *http.Request) {
		req.Header.Set("Content-Type", "application/json")
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
		if (*req).Method == "OPTIONS" {
			return
		}
		reader := strings.NewReader(apiFeeds)
		_, err := io.Copy(writer, reader)
		ErrorHandling(err)
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
	http.HandleFunc("/.well-known/pki-validation/A60C1FCD3DD9405F854273F29AFC7954.txt", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "C5DB15EF8C4A35769AC8BF6521B9DF5479B6382667C4B039BABBAD0A3FDE4EAA\ncomodoca.com\n3a49b7dc5e3ff77")
		return
	})
}
