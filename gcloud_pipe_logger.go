package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// Imports the Stackdriver Logging client package.
	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

func setupLogger() *logging.Client {
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func main() {
	var entry logging.Entry
	logName := os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)
	client := setupLogger()
	logger := client.Logger(logName)
	for scanner.Scan() {
		entry = logging.Entry{Payload: scanner.Text()}
		logger.Log(entry)
	}
	logger.Flush()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("Failed to close client: %v", err)
		}
	}()
}
