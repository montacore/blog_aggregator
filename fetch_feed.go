package main

import (
	"encoding/xml"
	"net/http"
	"fmt"
	"log"
	"context"
	"io"
)

type RSSFeed struct {
	Title			string		`xml:"title"`
	Link			string		`xml:"link"`
	Description		string		`xml:"description"`
	Item			[]RSSItem	`xml:"item"`
} `xml:"channel"`

type RSSItem struct {
	Title			string		`xml:"title"`
	Link			string		`xml:"link"`
	Description		string		`xml:"description"`
	PubDate			string		`xml:"pubDate"`
}

/*TODO: Write fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error)
// which should fetch a feed from the given URL, and assuming that nothing
// goes wrong, return a filled-out RSSFeed struct.
// helpful resources:
// 	http.NewRequestWithContext
// 	http.Client.Do
// 	set User-Agent header to gator in the request with request.Header.Set
// 	Which is a common practice to identify your program to the server. 
// 		io.Readall
// 		xml.Unmarshal(works just like json.Unmarshal)
// 		html.UnescapeString - decodes escaped html entities like `&ldquo;`
// 		This will be needed on the Title and Description fields(of both 
// 		the entire channel and individual items)
*/


// TODO: Add an `agg` command, this will ultimately become the long-running
// aggregator service. For now, we'll just use it to fetch a single feed
// and ensure our parsing is working. 
// It should fetch the feed found at: https://www.wagslane.dev/index.xml
// and print the entire struct to the console.


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("[error] encountered an error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[error] there was a problem with your request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[error] there was a problem reading response body: %w", err)
	}

	var result RSSFeed
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("[error] there was an error unmarshaling the xml response: %w", err)
	}

	return &result, nil
}
