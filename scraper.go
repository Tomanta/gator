package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tomanta/gator/internal/database"
	"time"
)

func scrapeFeeds(s *state) error {
	next_feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to get next feed: %w", err)
	}

	return scrapeFeed(s.db, next_feed)

}

func scrapeFeed(db *database.Queries, feed database.GetNextFeedToFetchRow) error {
	feed_update := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	}
	err := db.MarkFeedFetched(context.Background(), feed_update)
	if err != nil {
		return fmt.Errorf("Unable to mark feed as fetched: %w", err)
	}
	

	FeedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Unable to fetch feed %s: %w", feed.Url, err)
	}

	fmt.Printf("Feed fetched: %s.\n items:\n", FeedData.Channel.Title)
	for _, item := range FeedData.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	return nil

}
