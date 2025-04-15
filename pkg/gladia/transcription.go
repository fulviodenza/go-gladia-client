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

const uploadEndpoint = "v2/upload"
const transcribeEndpoint = "v2/pre-recorded/"

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

	resp, err := c.sendFormRequest(ctx, uploadEndpoint, body, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %s, body: %s", resp.Status, string(bodyBytes))
	}

	var uploadResponse UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &uploadResponse, nil
}

func (s *Client) Transcribe(ctx context.Context, audioURL string) (*TranscriptionResponse, error) {
	reqBody := TranscriptionRequest{AudioURL: audioURL}
	var result TranscriptionResponse

	err := s.sendJSONRequest(ctx, http.MethodPost, transcribeEndpoint, reqBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTranscriptionResult retrieves the result of a transcription by its ID
func (c *Client) GetTranscriptionResult(ctx context.Context, transcriptionID string) (*CompletedTranscriptionResult, error) {
	var result CompletedTranscriptionResult

	err := c.sendJSONRequest(ctx, http.MethodGet, transcribeEndpoint+transcriptionID, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTranscriptionStatus checks the status of a transcription by its ID
func (c *Client) GetTranscriptionStatus(ctx context.Context, transcriptionID string) (*GetTranscriptionStatus, error) {
	var status GetTranscriptionStatus

	err := c.sendJSONRequest(ctx, http.MethodGet, transcribeEndpoint+transcriptionID, nil, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (c *Client) sendJSONRequest(ctx context.Context, method, endpoint string, reqBody interface{}, respBody interface{}) error {
	var bodyReader io.Reader

	if reqBody != nil {
		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+endpoint, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set(gladiaHeaderKey, c.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("received non-200 response: %s, body: %s", resp.Status, string(bodyBytes))
	}

	if respBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// sendFormRequest sends a multipart form request to the Gladia API
func (c *Client) sendFormRequest(ctx context.Context, endpoint string, formData *bytes.Buffer, contentType string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+endpoint, formData)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set(gladiaHeaderKey, c.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return resp, nil
}
