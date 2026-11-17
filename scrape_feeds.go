package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/montacore/blog_aggregator/internal/database"
	"time"
)

func scrapeFeeds(s *state) {
	// NOTE: Fetch next feed
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("[error] error fetching next feed: %s\n", err)
		return
	}
	// NOTE: Mark as fetched
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("[error] could not mark feed fetched: %s", err)
		return
	}
	result, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("[error] could not fetch feed data: %s\n", err)
		return
	}
	for _, entry := range result.Channel.Item {

		parsedTime, dateErr := time.Parse(time.RFC1123Z, entry.PubDate)
		if dateErr != nil {
			fmt.Printf("[error] could not parse publication date %q: %v\n", entry.PubDate, dateErr)
		}
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     entry.Title,
			Url:       entry.Link,
			Description: sql.NullString{
				String: entry.Description,
				Valid:  entry.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  parsedTime,
				Valid: dateErr == nil,
			},
			FeedID: feed.ID,
		})
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			} else {
				fmt.Printf("[error] error encountered when attempting to create post: %s", err)
			}
		}
	}

}
