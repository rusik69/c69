package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rusik69/govnocloud/pkg/types"
)

// AddNode adds a host.
func AddNode(host, port, name, nodeHost, nodePort string) error {
	node := types.Node{
		Name: name,
		Host: nodeHost,
		Port: nodePort,
	}
	url := "http://" + host + ":" + port + "/api/v1/nodes"
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

// ListNodes lists nodes.
func ListNodes(host, port string) (map[string]types.NodeStats, error) {
	url := "http://" + host + ":" + port + "/api/v1/nodes"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var nodes map[string]types.NodeStats
	err = json.NewDecoder(resp.Body).Decode(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetNode gets a node.
func GetNode(host, port, name string) (types.Node, error) {
	url := "http://" + host + ":" + port + "/api/v1/node/" + name
	resp, err := http.Get(url)
	if err != nil {
		return types.Node{}, err
	}
	defer resp.Body.Close()
	var node types.Node
	err = json.NewDecoder(resp.Body).Decode(&node)
	if err != nil {
		return types.Node{}, err
	}
	return node, nil
}

// GetNodeStats gets a node stats.
func GetNodeStats(host, port string) (types.NodeStats, error) {
	url := "http://" + host + ":" + port + "/api/v1/stats"
	resp, err := http.Get(url)
	if err != nil {
		return types.NodeStats{}, err
	}
	defer resp.Body.Close()
	var nodeStats types.NodeStats
	err = json.NewDecoder(resp.Body).Decode(&nodeStats)
	if err != nil {
		return types.NodeStats{}, err
	}
	return nodeStats, nil
}
