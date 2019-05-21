package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/mmcdole/gofeed"
)

type FeedItem struct {
	Link      string
	Title     string
	Published string
}

type FeedPage struct {
	FeedTitle string
	Items     []*FeedItem
}

func main() {
	feedsList := os.Args[1:]

	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(feedsList[0])
	if err != nil {
		log.Println(err.Error())
	}

	data := FeedPage{
		FeedTitle: feed.Title,
		Items:     convert(feed.Items),
	}

	tmpl := template.Must(template.ParseFiles("./web/feed.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	log.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err.Error())
	}
}

func convert(gofeedItems []*gofeed.Item) (feedItems []*FeedItem) {
	for _, item := range gofeedItems {
		feedItems = append(feedItems, &FeedItem{
			Link:      item.Link,
			Title:     item.Title,
			Published: item.PublishedParsed.Format("Jan _2 2006"),
		})
	}

	return
}
