package slack

import (
	"os"
	"testing"
)

func TestSendMessage(t *testing.T) {
	if err := SendImageURL(
		os.Getenv("SLACK_TEST_WEBHOOK"),
		"No joke I found something like this in our code base",
		"https://i.redd.it/r1os45ulkat41.png",
	); err != nil {
		t.Fatalf("Failed to send image: %s", err)
	}
}
