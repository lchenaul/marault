package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type chatReq struct {
	Message string `json:"message"`
}

type chatRes struct {
	Reply string `json:"reply"`
}

const maraultSystemPrompt = `
You are the official AI assistant for Marault Intelligence.

Your role is to explain services clearly, conversationally, and professionally.
Do not repeat website copy verbatim.
Explain in plain English.
Use short answers.
Ask clarifying questions when needed.
Never provide pricing or fixed timelines.
Direct serious inquiries to /inquire.
`

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req chatReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Call AI
	reply, err := callOpenAI(maraultSystemPrompt, req.Message)
	if err != nil {
		fmt.Println("AI error:", err) // server-side log
		http.Error(w, "AI request failed", http.StatusBadGateway)
		return
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(chatRes{Reply: reply})
}

func callOpenAI(systemPrompt, userMsg string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("missing OPENAI_API_KEY env var")
	}

payload := map[string]any{
	"model": "gpt-4.1-mini",
	"max_output_tokens": 300,
	"input": []map[string]any{
		{
			"role": "system",
			"content": []map[string]any{
				{"type": "input_text", "text": systemPrompt},
			},
		},
		{
			"role": "user",
			"content": []map[string]any{
				{"type": "input_text", "text": userMsg},
			},
		},
	},
}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", "https://api.openai.com/v1/responses", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 25 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("provider error (%d): %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("bad json from provider: %w", err)
	}

	// Try to extract text from `output`
	text := extractResponseText(result)
	if text == "" {
		return "", fmt.Errorf("no text found in provider response: %s", string(respBody))
	}

	return text, nil
}

// Defensive parsing: avoids crashing if response shape varies slightly
func extractResponseText(result map[string]any) string {
	// Preferred: output[].content[].text
	if out, ok := result["output"].([]any); ok {
		for _, item := range out {
			m, _ := item.(map[string]any)
			if m == nil {
				continue
			}
			content, _ := m["content"].([]any)
			for _, c := range content {
				cm, _ := c.(map[string]any)
				if cm == nil {
					continue
				}
				if t, ok := cm["text"].(string); ok && t != "" {
					return t
				}
			}
		}
	}

	// Fallback: some responses include output_text
	if t, ok := result["output_text"].(string); ok && t != "" {
		return t
	}

	return ""
}
