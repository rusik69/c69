package client_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/rusik69/govnocloud/pkg/client"
)

var (
	masterHost string
	masterPort string
	nodeName   string
	nodeHost   string
	nodePort   string
)

// waitForMaster waits for the master to start.
func waitForMaster() {
	waitTime := 60
	for {
		_, err := http.Get("http://" + masterHost + ":" + masterPort + "/ping")
		if err == nil || waitTime == 0 {
			break
		}
		waitTime--
		time.Sleep(1 * time.Second)
	}
}

// TestMain is the main test function.
func TestMain(m *testing.M) {
	waitForMaster()
	masterHost = os.Getenv("TEST_MASTER_HOST")
	if masterHost == "" {
		masterHost = "localhost"
	}
	masterPort = os.Getenv("TEST_MASTER_PORT")
	if masterPort == "" {
		masterPort = "7070"
	}
	nodeName = os.Getenv("TEST_NODE_NAME")
	if nodeName == "" {
		nodeName = "localhost"
	}
	nodeHost = os.Getenv("TEST_NODE_HOST")
	if nodeHost == "" {
		nodeHost = "localhost"
	}
	nodePort = os.Getenv("TEST_NODE_PORT")
	if nodePort == "" {
		nodePort = "6969"
	}
	m.Run()
}

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
