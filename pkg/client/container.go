package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/rusik69/govnocloud/pkg/types"
)

// CreateContainer creates a container.
func CreateContainer(host, port, name, image, flavor string) (types.Container, error) {
	container := types.Container{
		Name:   name,
		Image:  image,
		Flavor: flavor,
	}
	url := "http://" + host + ":" + port + "/api/v1/containers"
	body, err := json.Marshal(container)
	if err != nil {
		return types.Container{}, err
	}
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return types.Container{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return types.Container{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Container{}, err
	}
	if resp.StatusCode != 200 {
		return types.Container{}, errors.New(string(bodyText))
	}
	err = json.Unmarshal(bodyText, &container)
	if err != nil {
		return types.Container{}, err
	}
	return container, nil
}

// StartContainer starts a container.
func StartContainer(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/containerstart/" + name
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
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
	return nil
}

// StopContainer stops a container.
func StopContainer(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/containerstop/" + name
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
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
	return nil
}

// ListContainers lists containers.
func ListContainers(host, port string) ([]types.Container, error) {
	url := "http://" + host + ":" + port + "/api/v1/containers"
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var containers []types.Container
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(bodyText))
	}
	err = json.Unmarshal(bodyText, &containers)
	if err != nil {
		return nil, err
	}
	return containers, nil
}

// GetContainer gets a container.
func GetContainer(host, port, name string) (types.Container, error) {
	container := types.Container{
		ID: name,
	}
	url := "http://" + host + ":" + port + "/api/v1/container/" + name
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return container, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return container, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return container, err
	}
	if resp.StatusCode != 200 {
		return container, errors.New(string(bodyText))
	}
	err = json.Unmarshal(bodyText, &container)
	if err != nil {
		return container, err
	}
	return container, nil
}

// DeleteContainer deletes a container.
func DeleteContainer(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/container/" + name
	client := &http.Client{
		Timeout: 300 * time.Second,
	}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
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
	return nil
}
