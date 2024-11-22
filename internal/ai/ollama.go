package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type OllamaRequest struct {
	Model   string        `json:"model"`
	System  string        `json:"system"`
	Prompt  string        `json:"prompt"`
	Options OllamaOptions `json:"options"`
	Stream  bool          `json:"stream"`
}

type OllamaOptions struct {
	NumCtx      int     `json:"num_ctx"`
	Temperature float32 `json:"temperature"`
	NumKeep     int     `json:"num_keep"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func QueryOllama(prompt, system, host, model string, options OllamaOptions) (string, error) {
	url := fmt.Sprintf("%s/api/generate", host)

	requestBody, err := json.Marshal(OllamaRequest{
		Model:   model,
		System:  system,
		Prompt:  prompt,
		Options: options,
		Stream:  false,
	})
	if err != nil {
		return "", fmt.Errorf("could not marshal request body: %w", err)
	}

	slog.Debug("Request", "requestBody", string(requestBody))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("could not make POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body: %w", err)
	}

	var ollamaResp OllamaResponse
	err = json.Unmarshal(body, &ollamaResp)
	if err != nil {
		return "", fmt.Errorf("could not unmarshal response body: %w", err)
	}

	return ollamaResp.Response, nil
}
