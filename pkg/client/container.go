package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/rusik69/govnocloud/pkg/types"
)

// CreateContainer creates a container.
func CreateContainer(host, port, name, image string) (int, error) {
	container := types.Container{
		Name:  name,
		Image: image,
	}
	url := "http://" + host + ":" + port + "/api/v1/container/create"
	body, err := json.Marshal(container)
	if err != nil {
		return 0, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New(string(bodyText))
	}
	err = json.Unmarshal(bodyText, &container)
	if err != nil {
		return 0, err
	}
	idInt, err := strconv.Atoi(container.ID)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}

// StartContainer starts a container.
func StartContainer(host, port string, id int) error {
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/container/start/" + idString
	resp, err := http.Get(url)
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
func StopContainer(host, port string, id int) error {
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/container/stop/" + idString
	resp, err := http.Get(url)
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
	url := "http://" + host + ":" + port + "/api/v1/container/list"
	resp, err := http.Get(url)
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
func GetContainer(host, port string, id int) (types.Container, error) {
	idString := strconv.Itoa(id)
	container := types.Container{
		ID: idString,
	}
	url := "http://" + host + ":" + port + "/api/v1/container/" + idString
	resp, err := http.Get(url)
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
func DeleteContainer(host, port string, id int) error {
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/container/" + idString
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
	return nil
}
