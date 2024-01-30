package client_test

import (
	"fmt"
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

var (
	ContainerID string
)

// TestCreateContainer tests the CreateContainer function.
func TestCreateContainer(t *testing.T) {
	ContainerID, err := client.CreateContainer(masterHost, masterPort, "test", "busybox")
	if err != nil {
		t.Error(err)
	}
	if ContainerID == "" {
		t.Error("expected not 0, got ", ContainerID)
	}
}

// TestGETContainer tests the GetContainer function.
func TestGetContainer(t *testing.T) {
	container, err := client.GetContainer(masterHost, masterPort, "test")
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
	if len(containers) != 3 {
		t.Error("expected 3 container, got ", len(containers))
	}
}

// TestStopContainer tests the StopContainer function.
func TestStopContainer(t *testing.T) {
	err := client.StopContainer(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestStartContainer tests the StartContainer function.
func TestStartContainer(t *testing.T) {
	err := client.StartContainer(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteContainer tests the DeleteContainer function.
func TestDeleteContainer(t *testing.T) {
	err := client.DeleteContainer(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// RunContainers runs containers.
func TestRunContainers(t *testing.T) {
	for i := 1; i < 10; i++ {
		containerName := "test" + fmt.Sprintf("%d", i)
		_, err := client.CreateContainer(masterHost, masterPort, containerName, "nginx")
		if err != nil {
			t.Error(err)
		}
	}
}
