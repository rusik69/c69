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
		nodeName = "node0"
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
	RunContainers()
	RunVMs()
	UploadFiles()
}

// TestAddNode tests the AddNode function.
func TestAddNode(t *testing.T) {
	err := client.AddNode(masterHost, masterPort, nodeName, nodeHost, nodePort)
	if err != nil {
		t.Error(err)
	}
}

// TestListNodes tests the ListNodes function.
func TestListNodes(t *testing.T) {
	nodes, err := client.ListNodes(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(nodes) != 2 {
		t.Error("expected 2 nodes, got ", len(nodes))
	}
}

// TestGetNode tests the GetNode function.
func TestGetNode(t *testing.T) {
	node, err := client.GetNode(masterHost, masterPort, nodeName)
	if err != nil {
		t.Error(err)
	}
	if node.Name != "x220" {
		t.Error("expected x220, got ", node.Name)
	}
}

// TestDeleteNode tests the DeleteNode function.
func TestDeleteNode(t *testing.T) {
	err := client.DeleteNode(masterHost, masterPort, nodeName)
	if err != nil {
		t.Error(err)
	}
}

func TestAddNodes(t *testing.T) {
	client.AddNode(masterHost, masterPort, "x220", "x220.rusik69.lol", "6969")
}
