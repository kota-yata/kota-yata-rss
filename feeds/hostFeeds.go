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
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/rss", func(writer http.ResponseWriter, req *http.Request) {
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
	http.HandleFunc("/google900b28595c041e06.html", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "google-site-verification: google900b28595c041e06.html")
		return
	})
}
