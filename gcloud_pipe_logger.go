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

func setupLogger() *logging.Logger {
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	logName := os.Args[1]

	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	logger := client.Logger(logName)

	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("Failed to close client: %v", err)
		}
	}()
	return logger
}

func main() {
	var entry logging.Entry
	scanner := bufio.NewScanner(os.Stdin)
	logger := setupLogger()
	for scanner.Scan() {
		entry = logging.Entry{Payload: scanner.Text()}
		logger.Log(entry)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
