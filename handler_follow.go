package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VoluteTech/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	url := cmd.Args[0]

	dbUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get the user: %w", err)
	}

	dbFeed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed with this url: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: dbUser.ID,
		FeedID: dbFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create the feed follow record: %w", err)
	}

	fmt.Printf("User '%s' is now following the feed '%s'", dbUser.Name, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("coulnd't get the user: %w", err)
	}

	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get the feed followed for user: %w", err)
	}
	if len(feedsFollowed) == 0 {
		fmt.Println("No feeds followed for the current user")
		return nil
	}

	fmt.Printf("User: %s\n", user.Name)
	for _, feed := range feedsFollowed {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}

