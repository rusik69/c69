package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

// TestCreateK8S tests the CreateK8S function.
func TestCreateK8S(t *testing.T) {
	k8s, err := client.CreateK8S(masterHost, masterPort, "test", "small")
	if err != nil {
		t.Error(err)
	}
	if k8s.Name != "test" {
		t.Error("expected test, got ", k8s.Name)
	}
}

// TestGetK8S tests the GetK8S function.
func TestGetK8S(t *testing.T) {
	k8s, err := client.GetK8S(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
	if k8s.Name != "test" {
		t.Error("expected test, got ", k8s.Name)
	}
}

// TestListK8Ss tests the ListK8Ss function.
func TestListK8Ss(t *testing.T) {
	k8ss, err := client.ListK8S(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(k8ss) != 1 {
		t.Error("expected 1 k8s, got ", len(k8ss))
	}
}

// TestStopK8S tests the StopK8S function.
func TestStopK8s(t *testing.T) {
	err := client.StopK8S(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestStartK8S tests the StartK8S function.
func TestStartK8S(t *testing.T) {
	err := client.StartK8S(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteK8S tests the DeleteK8S function.
func TestDeleteK8S(t *testing.T) {
	err := client.DeleteK8S(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}
