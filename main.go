package main

import (
	"fmt"
	"net/http"
	"io"
	"strings"
	"time"
)

var (
	urlsToScrape = []string{
		"https://go.dev/",
		"https://www.rust-lang.org/",
		"https://elixir-lang.org/",
		"https://github.com/",
		"https://www.google.com/",
	}
	
	// Channel (used for safe communication between goroutines)
	resultChannel = make(chan string) 
)

// Function to scrape a single URL concurrently
func scrapeURL(url string) {
	fmt.Printf("Scraping started: %s\n", url)
	
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		resultChannel <- fmt.Sprintf("❌ Error (%s): %v", url, err)
		return
	}
	// Ensure the response body is closed when the function finishes
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resultChannel <- fmt.Sprintf("❌ Read Error (%s): %v", url, err)
		return
	}

	// Simple title extraction logic
	bodyStr := string(body)
	start := strings.Index(bodyStr, "<title>")
	end := strings.Index(bodyStr, "</title>")

	title := "Title Not Found"
	if start != -1 && end != -1 && start < end {
		title = bodyStr[start+7 : end]
	}

	// Send the successful result to the channel
	resultChannel <- fmt.Sprintf("✅ TITLE: %s | URL: %s", title, url)
}

func main() {
	fmt.Println("--- Starting Concurrent Web Scraper ---")

	// STEP 1: Start a goroutine for each URL
	for _, url := range urlsToScrape {
		// The 'go' keyword starts the function in a new lightweight concurrent process (goroutine).
		go scrapeURL(url) 
	}
	
	// STEP 2: Collect results from the channel
	// We wait for exactly the number of results we expect.
	for i := 0; i < len(urlsToScrape); i++ {
		// This line blocks (waits) until a value is sent to the channel.
		result := <-resultChannel 
		fmt.Println(result)
	}

	fmt.Println("--- All Scrapes Completed. ---")
}
