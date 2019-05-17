package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	feedsList := os.Args[1:]

	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(feedsList[0])
	if err != nil {
		log.Println(err.Error())
	}

	render(feed.Items)
}

func render(items []*gofeed.Item) {
	for _, item := range items {
		if len(item.Title) > 70 {
			fmt.Printf("%v\n", item.Title)
			fmt.Printf("%70v %s\n", "", item.Link)
		} else {
			fmt.Printf("%-70v %s\n", item.Title, item.Link)
		}
	}
}
