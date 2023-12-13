package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	masterEnv "github.com/rusik69/govnocloud/pkg/master/env"
)

// AddNode adds a host.
func AddNode(host, port, name, nodeHost, nodePort string) error {
	node := masterEnv.Node{
		Name: name,
		Host: nodeHost,
		Port: nodePort,
	}
	url := "http://" + host + ":" + port + "/api/v1/node/add"
	body, err := json.Marshal(node)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
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

// DeleteNode deletes a node.
func DeleteNode(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/node/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
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

// ListNodes lists nodes.
func ListNodes(host, port string) ([]masterEnv.Node, error) {
	url := "http://" + host + ":" + port + "/api/v1/node/list"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var nodes []masterEnv.Node
	err = json.NewDecoder(resp.Body).Decode(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetNode gets a node.
func GetNode(host, port, name string) (masterEnv.Node, error) {
	url := "http://" + host + ":" + port + "/api/v1/node/" + name
	resp, err := http.Get(url)
	if err != nil {
		return masterEnv.Node{}, err
	}
	defer resp.Body.Close()
	var node masterEnv.Node
	err = json.NewDecoder(resp.Body).Decode(&node)
	if err != nil {
		return masterEnv.Node{}, err
	}
	return node, nil
}
