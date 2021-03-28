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
	portName := os.Getenv("PORT")
	if portName == "" {
		portName = "3432"
	}
	fmt.Println("RSS feed has been published at http://localhost:3432")
	err := http.ListenAndServe(":" + portName, nil)
	ErrorHandling(err)
}
