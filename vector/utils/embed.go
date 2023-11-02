package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}
type EmbeddingRequest struct {
	Input any    `json:"input"`
	Model string `json:"model"`
}
type EmbeddingResponse struct {
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
	Model  string      `json:"model"`
	Usage  Usage       `json:"usage"`
}

func CreateEmbeddings(
	inputs []string,
	model string,
	apiKey string,
) (*[][]float32, error) {
	url := "https://api.openai.com/v1/embeddings"

	payload := EmbeddingRequest{
		Input: inputs,
		Model: model,
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var embeddingResponse EmbeddingResponse

	json.NewDecoder(resp.Body).Decode(&embeddingResponse)

	var embeddings [][]float32
	for _, v := range embeddingResponse.Data {
		embeddings = append(embeddings, v.Embedding)
	}
	return &embeddings, nil
}
