package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

// TestCreateVM tests the CreateVM function.
func TestCreateVM(t *testing.T) {
	vm, err := client.CreateVM(masterHost, masterPort, "test", "ubuntu24.04", "medium")
	if err != nil {
		t.Error(err)
	}
	if vm.Name != "test" {
		t.Error("expected test, got ", vm.Name)
	}
	if vm.Image != "ubuntu24.04" {
		t.Error("expected ubuntu24.04, got ", vm.Image)
	}
	if vm.Flavor != "medium" {
		t.Error("expected medium, got ", vm.Flavor)
	}
	if vm.Committed != true {
		t.Error("expected true, got ", vm.Committed)
	}

}

// TestGETVM tests the GetVM function.
func TestGetVM(t *testing.T) {
	vm, err := client.GetVM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
	if vm.Name != "test" {
		t.Error("expected test, got ", vm.Name)
	}
}

// TestListVMs tests the ListVMs function.
func TestListVMs(t *testing.T) {
	vms, err := client.ListVMs(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(vms) != 1 {
		t.Error("expected 1 vm, got ", len(vms))
	}
}

// TestStopVM tests the StopVM function.
func TestStopVM(t *testing.T) {
	err := client.StopVM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestStartVM tests the StartVM function.
func TestStartVM(t *testing.T) {
	err := client.StartVM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteVM tests the DeleteVM function.
func TestDeleteVM(t *testing.T) {
	err := client.DeleteVM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}
