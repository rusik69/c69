package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rusik69/govnocloud/pkg/types"
)

// CreateLLM creates a llm cluster.
func CreateLLM(host, port, name, model string) (types.LLM, error) {
	llm := types.LLM{
		Name:  name,
		Model: model,
	}
	url := "http://" + host + ":" + port + "/api/v1/llm"
	body, err := json.Marshal(llm)
	if err != nil {
		return types.LLM{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return types.LLM{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.LLM{}, err
	}
	if resp.StatusCode != 200 {
		return types.LLM{}, errors.New(string(bodyText))
	}
	var newLLM types.LLM
	err = json.Unmarshal(bodyText, &newLLM)
	if err != nil {
		return types.LLM{}, err
	}
	return llm, nil
}

// GetLLM gets a llm cluster.
func GetLLM(host, port, name string) (types.LLM, error) {
	url := "http://" + host + ":" + port + "/api/v1/llm/" + name
	resp, err := http.Get(url)
	if err != nil {
		return types.LLM{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.LLM{}, err
	}
	if resp.StatusCode != 200 {
		return types.LLM{}, errors.New(string(bodyText))
	}
	var llm types.LLM
	err = json.Unmarshal(bodyText, &llm)
	if err != nil {
		return types.LLM{}, err
	}
	return llm, nil
}

// DeleteLLM deletes a llm cluster.
func DeleteLLM(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/llm/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	return err
}

// ListLLMs lists all llm clusters.
func ListLLMs(host, port string) ([]types.LLM, error) {
	url := "http://" + host + ":" + port + "/api/v1/llm"
	resp, err := http.Get(url)
	if err != nil {
		return []types.LLM{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return []types.LLM{}, err
	}
	if resp.StatusCode != 200 {
		return []types.LLM{}, errors.New(string(bodyText))
	}
	var llms []types.LLM
	err = json.Unmarshal(bodyText, &llms)
	if err != nil {
		return []types.LLM{}, err
	}
	return llms, nil
}

// StartLLM starts a llm cluster.
func StartLLM(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/llmstart/" + name
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(bodyText))
	}
	return nil
}

// StopLLM stops a llm cluster.
func StopLLM(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/llmstop/" + name
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(bodyText))
	}
	return nil
}

// GenerateLLM generates a response from llm.
func GenerateLLM(host, port, ip, name, input string) (string, error) {
	url := "http://" + host + ":" + port + "/api/v1/llmgenerate/" + name
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(input)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(string(bodyText))
	}
	return string(bodyText), nil
}
