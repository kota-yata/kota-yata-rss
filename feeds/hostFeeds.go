package feeds

import (
	"fmt"
	"io"
	"net/http"
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
	// portName := ":" + os.Getenv("PORT")
	fmt.Println("RSS feed has been published at http://localhost:3432")
	err := http.ListenAndServe(":3432", nil)
	ErrorHandling(err)
}
