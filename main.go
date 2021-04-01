package main

import (
	"integrated-rss/feeds"
)

func main() {
	items := feeds.GetAllItem()
	feedArray := feeds.CreateFeeds(items)
	feeds.SetHost(feedArray)
}
