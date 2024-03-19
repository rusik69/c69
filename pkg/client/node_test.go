package client_test

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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
	nodes      []string
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

// waitForNodes waits for the nodes to start.
func waitForNodes() {
	waitTime := 60
	for _, node := range nodes {
		s := strings.Split(node, ":")
		host := s[1]
		port := s[2]
		for {
			_, err := http.Get("http://" + host + ":" + port + "/ping")
			if err == nil || waitTime == 0 {
				break
			}
			waitTime--
			time.Sleep(1 * time.Second)
		}
	}
}

// TestMain is the main test function.
func TestMain(m *testing.M) {
	waitForMaster()
	waitForNodes()
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
	nodesString := os.Getenv("TEST_NODES")
	if nodesString == "" {
		nodes = []string{}
	}
	nodesSplit := strings.Split(nodesString, ",")
	nodes = append(nodes, nodesSplit...)
	addNodes()
	m.Run()
	addNodes()
	UploadFiles()
}

// TestAddNode tests the AddNode function.
func TestAddNode(t *testing.T) {
	err := client.AddNode(masterHost, masterPort, nodeName, nodeHost, nodePort)
	if err != nil {
		t.Error(err)
	}
}

// addNode adds node
func addNodes() {
	for _, node := range nodes {
		s := strings.Split(node, ":")
		host := s[0]
		port := s[1]
		name := strings.Split(host, ".")[0]
		err := client.AddNode(masterHost, masterPort, name, host, port)
		if err != nil {
			fmt.Println("Error adding node: ", err.Error())
			continue
		}
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
	if node.Name != "node0" {
		t.Error("expected node0, got ", node.Name)
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
	err := client.AddNode(masterHost, masterPort, "node0", "node0.govno.cloud", "6969")
	if err != nil {
		t.Error(err)
	}
}
