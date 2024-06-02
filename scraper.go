package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ayushrakesh/go-rssagg/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenFetch time.Duration) {

	log.Printf(`Fetching feeds with concurrency %v with time duration of %v`, concurrency, timeBetweenFetch)

	ticker := time.NewTicker(timeBetweenFetch)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("error fetching feeds- %v", err)
			continue
		}

		wg := sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, feed, &wg)
		}

		wg.Wait()

	}

}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	_, errr := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if errr != nil {
		log.Printf("Error marking feed as fetched- %v", errr)
		return
	}

	rssFeed, errrr := urlToFeed(feed.Url)
	if errrr != nil {
		log.Printf("Error fetching feed- %v", errrr)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %v collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
