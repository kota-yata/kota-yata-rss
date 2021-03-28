package main

import (
	"integrated-rss/feeds"
)

func main() {
	items := feeds.GetAllItem()
	rss := feeds.CreateFeeds(items)
	feeds.SetHost(rss)
}
