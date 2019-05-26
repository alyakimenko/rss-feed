package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

type FeedItem struct {
	Link      string
	Title     string
	Published string
}

type FeedPage struct {
	FeedTitle       string
	FeedDescription string
	Items           []*FeedItem
}

func main() {
	feedsList := os.Args[1:]

	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(feedsList[0])
	if err != nil {
		log.Println(err.Error())
	}

	data := &FeedPage{
		FeedTitle:       feed.Title,
		FeedDescription: feed.Description,
		Items:           convert(feed.Items),
	}

	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "feed.html", data)
	}).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", router)

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
