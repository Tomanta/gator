package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"github.com/tomanta/gator/internal/database"
	"github.com/google/uuid"	
	"time"
)

func scrapeFeeds(s *state) error {
	next_feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to get next feed: %w", err)
	}

	return scrapeFeed(s.db, next_feed)

}

func savePost(db *database.Queries, post RSSItem, feedID uuid.UUID) error {

	pubdate, err := time.Parse(time.RFC1123Z, post.PubDate)
	if err != nil {
		return fmt.Errorf("Error saving post, cannot create pubtime: %w", err)
	}

	postToCreate := database.CreatePostParams {
		ID: uuid.New(),
		Title: sql.NullString{
			String: post.Title,
			Valid: true,
		},
		Url: post.Link,
		Description: sql.NullString{
			String: post.Description,
			Valid: true,
		},
		PublishedAt: sql.NullTime {
			Time: pubdate,
			Valid: true,
		},
		FeedID: feedID,
	}

	_, err = db.CreatePost(context.Background(), postToCreate)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		}
		return fmt.Errorf("Error saving post: %w", err)
	}
	return nil
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

	for _, item := range FeedData.Channel.Item {
		err := savePost(db, item, feed.ID)
		if err != nil {
			return fmt.Errorf("Error saving post: %w", err)
		}
	}
	return nil

}
