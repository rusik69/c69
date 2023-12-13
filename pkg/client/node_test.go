package client_test

import (
	"testing"

	"github.com/rusik69/ds0/pkg/client"
)

// TestAddNode tests the AddNode function.
func TestAddNode(t *testing.T) {
	err := client.AddNode("localhost", "7070", "localhost", "localhost", "6969")
	if err != nil {
		t.Error(err)
	}
}

// TestListNodes tests the ListNodes function.
func TestListNodes(t *testing.T) {
	nodes, err := client.ListNodes("localhost", "7070")
	if err != nil {
		t.Error(err)
	}
	if len(nodes) != 1 {
		t.Error("expected 1 node, got ", len(nodes))
	}
}

// TestGetNode tests the GetNode function.
func TestGetNode(t *testing.T) {
	node, err := client.GetNode("localhost", "7070", "localhost")
	if err != nil {
		t.Error(err)
	}
	if node.Name != "localhost" {
		t.Error("expected localhost, got ", node.Name)
	}
}

// TestDeleteNode tests the DeleteNode function.
func TestDeleteNode(t *testing.T) {
	err := client.DeleteNode("localhost", "7070", "localhost")
	if err != nil {
		t.Error(err)
	}
}
