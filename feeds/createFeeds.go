package feeds

import (
	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
	"time"
)

func CreateFeeds(items []*gofeed.Item) []string {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "kota-yata integrated RSS",
		Link:        &feeds.Link{
			Href: "https://feed.kota-yata.com",
		},
		Description: "The integrated RSS feed of blogs kota-yata wrote. Including Zenn, Qiita, blog.kota-yata.com, note",
		Author:       &feeds.Author{
			Name: "Kota Yatagai",
			Email: "kota@yatagai.com",
		},
		Created:     now,
		Items: consistFeedItems(items),
	}
	rss, err := feed.ToRss()
	ErrorHandling(err)
	api, err := feed.ToJSON()
	ErrorHandling(err)
	returnedArray := []string{rss, api}
	return returnedArray
}

func consistFeedItems(items []*gofeed.Item) []*feeds.Item {
	var resultItems []*feeds.Item
	for _, item := range items {
		publishedTime, err := time.Parse("20060102150405", item.Published)
		ErrorHandling(err)
		result := &feeds.Item{
			Title: item.Title,
			Link: &feeds.Link{
				Href:item.Link,
			},
			Author: &feeds.Author{
				Name: "Kota Yatagai",
				Email: "kota@yatagai.com",
			},
			Description: item.Description,
			Updated: publishedTime,
			Created: publishedTime,
			Id: item.GUID,
			Content: item.Content,
		}
		resultItems = append(resultItems, result)
	}
	return resultItems
}
