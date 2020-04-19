package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SendImageURL will send the message to slack.
func SendImageURL(slackHook string, text string, imageURL string) error {
	if slackHook == "" {
		return fmt.Errorf("Slack Webhook is empty")
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	data, err := genBlock(text, imageURL)
	if err != nil {
		return fmt.Errorf("Unable to create slack block: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, slackHook, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("Malformed Request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to send request to Slack: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error Occured sending message to slack. Failed to get body. Status Code %v", resp.StatusCode)
		}
		return fmt.Errorf("Error occured sending message to slack: %s", string(data))
	}

	return nil
}

// map[string]interface{} is aliased because I'm lazy.
type m map[string]interface{}

func genBlock(text string, imageURL string) ([]byte, error) {
	return json.Marshal(m{
		"blocks": []m{
			{
				"type": "section",
				"text": m{
					"type": "mrkdwn",
					"text": fmt.Sprintf(
						`"*%s*"`,
						text,
					),
				},
			},
			{
				"type":      "image",
				"image_url": imageURL,
				"alt_text":  text,
			},
			{ // This part will be hardcoded to fit our usecase.
				"type": "section",
				"text": m{
					"type": "mrkdwn",
					"text": "*ᵀᵃᵏᵉⁿ ᶠʳᵒᵐ ᴾʳᵒᵍʳᵃᵐᵐᵉʳᴴᵘᵐᵒʳ ᵒⁿ ᴿᵉᵈᵈᶦᵗ*",
				},
			},
		},
	})
}
