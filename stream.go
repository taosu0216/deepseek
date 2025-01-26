package deepseek

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (sr *StreamReader) Close() error {
	return sr.body.Close()
}

func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request ChatCompletionRequest,
) (*StreamReader, error) {
	request.Stream = true
	//request.ToolChoice = ToolChoice{
	//	Type:     deepseek.ChatCompletionToolTypeFunc,
	//	Function: nil,
	//}
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

	return &StreamReader{
		scanner: bufio.NewScanner(resp.Body),
		body:    resp.Body, // 保存响应体引用
	}, nil
}

func (sr *StreamReader) Recv() (*ChatCompletionResponse, error) {
	for sr.scanner.Scan() {
		line := sr.scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		if !bytes.HasPrefix(line, []byte("data: ")) {
			continue
		}

		data := bytes.TrimPrefix(line, []byte("data: "))
		if string(data) == "[DONE]" {
			return nil, io.EOF
		}

		var response ChatCompletionResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal event data: %w", err)
		}

		return &response, nil
	}

	if err := sr.scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return nil, io.EOF
}
