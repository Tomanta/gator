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

	feed_update := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: next_feed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), feed_update)
	if err != nil {
		return fmt.Errorf("Unable to mark feed as fetched: %w", err)
	}

	feedDetails, err := fetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		return fmt.Errorf("Unable to fetch feed %s: %w", next_feed.Url, err)
	}

	fmt.Printf("Feed fetched: %s.\n items:\n", feedDetails.Channel.Title)
	for _, item := range feedDetails.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	return nil

}
