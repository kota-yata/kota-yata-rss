package feeds

import (
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"sort"
	"time"
)

type individualFeed struct {
	service string
	url string
}

func ErrorHandling(errMessage error) {
	if errMessage != nil {
		log.Fatal(errMessage)
		os.Exit(1)
	}
}

func GetAllItem()[]*gofeed.Item {
	parser := gofeed.NewParser()
	var combinedFeedArray []*gofeed.Item
	layout := "20060102150405"
	timeZoneJST, err := time.LoadLocation("Asia/Tokyo")
	ErrorHandling(err)
	individualFeedArray := []*individualFeed{
		{
			service: "owned",
			url:     "https://blog.kota-yata.com/rss.xml",
		},
		{
			service: "zenn",
			url:     "https://zenn.dev/kota_yata/feed",
		},
		{
			service: "note",
			url:     "https://note.com/kotay/rss",
		},
	}
	for i := 0; i < len(individualFeedArray); i++ {
		feed, err := parser.ParseURL(individualFeedArray[i].url)
		ErrorHandling(err)
		// フィード内の記事一つ一つの公開日時をソート用にフォーマットする
		for _, item := range feed.Items {
			item.Published = formatTime(individualFeedArray[i].service, item.Published, timeZoneJST, layout)
		}
		combinedFeedArray = append(combinedFeedArray, feed.Items...)
	}
	sortByTime(combinedFeedArray, layout)
	return combinedFeedArray
}

// 畜生みたいな日付フォーマットをさばいていく
func formatTime(service string, publishedDate string, timeZone *time.Location, layout string) string {
	originalLayout := "Mon, 02 Jan 2006 15:04:05 GMT" // RFC1123
	if service == "note" {
		originalLayout = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123Z
	}
	formattedTime, err := time.ParseInLocation(originalLayout, publishedDate, timeZone)
	ErrorHandling(err)
	formattedTimeString := formattedTime.Format(layout)
	return formattedTimeString
}

func sortByTime(combinedFeedArray []*gofeed.Item, layout string) {
	sort.Slice(combinedFeedArray, func(i, j int) bool {
		timei, err := time.Parse(layout, combinedFeedArray[i].Published)
		ErrorHandling(err)
		timej, err := time.Parse(layout, combinedFeedArray[j].Published)
		ErrorHandling(err)
		return timei.After(timej)
	})
}
