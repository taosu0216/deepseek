package deepseek

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request *ChatCompletionRequest,
) (*ChatCompletionResponse, error) {
	request.Stream = false

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	defer func(Body io.ReadCloser) {
		deferErr := Body.Close()
		if deferErr != nil {
			fmt.Println(deferErr)
		}
	}(resp.Body)

	var response ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
