package main

import (
	"context"
	"fmt"
	"time"
	"strconv"
	"github.com/google/uuid"
	"github.com/tomanta/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	}

	/*****/
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("Could not get create feed: %w\n", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Could not follow feed: %w\n", err)
	}

	fmt.Printf("Feed created! %v\n", feed)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting feed list: %w\n", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))

	for _, feed := range feeds {
		fmt.Printf("%s: %s, added by: %s\n", feed.Name, feed.Url, feed.Creator)
	}

	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <url>\n", cmd.name)
	}

	feed_to_follow, err := getFeedByURL(s, cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Unable to look up feed: %w\n", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed_to_follow.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Could not follow feed: %w\n", err)
	}

	fmt.Printf("User %s is now following %s\n", feed_follow.UserName, feed_follow.FeedName)

	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <url>\n", cmd.name)
	}

	feed_to_unfollow, err := getFeedByURL(s, cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Unable to look up feed: %w\n", err)
	}

	deleteFeedFollow := database.DeleteFeedFollowParams{
		FeedID: feed_to_unfollow.ID,
		UserID: user.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), deleteFeedFollow)
	if err != nil {
		return fmt.Errorf("Could not delete feed follow: %w\n", err)
	}

	fmt.Printf("User %s is no longer following %s\n", user.Name, feed_to_unfollow.Name)

	return nil
}

func handlerListFollowedFeeds(s *state, cmd command, user database.User) error {

	feed_list, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Could not get feed list: %w\n", err)
	}

	if len(feed_list) == 0 {
		fmt.Printf("%s is not following any feeds.\n", user.Name)
		return nil
	}

	fmt.Printf("%s is following these feeds:\n", user.Name)
	for _, feed := range feed_list {
		fmt.Printf("* %s at %s\n", feed.FeedName, feed.FeedUrl)
	}
	return nil

}

func getFeedByURL(s *state, url string) (database.GetFeedByURLRow, error) {
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return database.GetFeedByURLRow{}, fmt.Errorf("Error retrieving feed: %w", err)
	}
	return feed, nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) == 1 {
		if customLimit, err := strconv.Atoi(cmd.arguments[0]); err == nil {
			limit = customLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	queryParams := database.GetPostsParams {
		UserID: user.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPosts(context.Background(), queryParams)
	if err != nil {
		return fmt.Errorf("Error browsing posts: %w", err)
	}

	fmt.Printf("%d posts found!\n", len(posts))
	for _, post := range posts {
		fmt.Printf("%s\n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}