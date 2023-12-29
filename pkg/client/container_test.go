package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

var (
	ContainerID int
)

// TestCreateContainer tests the CreateContainer function.
func TestCreateContainer(t *testing.T) {
	ContainerID, err := client.CreateContainer(masterHost, masterPort, "test", "ubuntu22.04")
	if err != nil {
		t.Error(err)
	}
	if ContainerID == 0 {
		t.Error("expected not 0, got ", ContainerID)
	}
}

// TestGETContainer tests the GetContainer function.
func TestGetContainer(t *testing.T) {
	container, err := client.GetContainer(masterHost, masterPort, ContainerID)
	if err != nil {
		t.Error(err)
	}
	if container.Name != "test" {
		t.Error("expected test, got ", container.Name)
	}
}

// TestListContainers tests the ListContainers function.
func TestListContainers(t *testing.T) {
	containers, err := client.ListContainers(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(containers) != 1 {
		t.Error("expected 1 container, got ", len(containers))
	}
}

// TestStopContainer tests the StopContainer function.
func TestStopContainer(t *testing.T) {
	err := client.StopContainer(masterHost, masterPort, ContainerID)
	if err != nil {
		t.Error(err)
	}
}

// TestStartContainer tests the StartContainer function.
func TestStartContainer(t *testing.T) {
	err := client.StartContainer(masterHost, masterPort, ContainerID)
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteContainer tests the DeleteContainer function.
func TestDeleteContainer(t *testing.T) {
	err := client.DeleteContainer(masterHost, masterPort, ContainerID)
	if err != nil {
		t.Error(err)
	}
}
