package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/Abinet16/rss/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed){
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		db.CreatePost(context.Background(),
		database.CreatePostParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(), 
		UpdatedAt:  time.Now().UTC(),
		Title:      item.Channel.Title,
		Description: sql.NullString{
			String: item.Channel.Description,
			Valid:  true,
		},
	})   
		log.Println("Found post", item.Channel.Title, "on feed", feed.Name)
	}


	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}