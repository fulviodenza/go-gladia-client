package gladia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const gladiaHeaderKey = "x-gladia-key:"
const uploadEndpoint = "v2/upload"
const transcribeEndpoint = "v2/pre-recorded"

// UploadResponse represents the response from the Gladia upload API
type UploadResponse struct {
	AudioURL      string        `json:"audio_url"`
	AudioMetadata AudioMetadata `json:"audio_metadata"`
}

// AudioMetadata contains information about the uploaded audio file
type AudioMetadata struct {
	ID               string  `json:"id"`
	Filename         string  `json:"filename"`
	Extension        string  `json:"extension"`
	Size             int64   `json:"size"`
	AudioDuration    float64 `json:"audio_duration"`
	NumberOfChannels int     `json:"number_of_channels"`
}

// UploadFile uploads an audio file to Gladia API and returns the audio URL that can be used for transcription
func (c *Client) UploadFile(ctx context.Context, filePath string) (*UploadResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("audio", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+uploadEndpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(gladiaHeaderKey, c.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %s, body: %s", resp.Status, string(bodyBytes))
	}

	var uploadResponse UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &uploadResponse, nil
}

func (s *Client) Transcribe(ctx context.Context, audioURL string) error {
	reqBody := TranscriptionRequest{AudioURL: audioURL}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.BaseURL+transcribeEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(gladiaHeaderKey, s.APIKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	return nil
}

func (s *Client) GetResult(ctx context.Context, id string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.BaseURL+transcribeEndpoint+id, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(gladiaHeaderKey, s.APIKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	p := []byte{}
	_, err = resp.Body.Read(p)
	if err != nil {
		return "", err
	}

	message := ""
	err = json.Unmarshal(p, &message)
	if err != nil {
		return "", err
	}

	return message, nil
}
