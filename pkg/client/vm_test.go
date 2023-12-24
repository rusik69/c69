package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

var (
	VMID int
)

// TestCreateVM tests the CreateVM function.
func TestCreateVM(t *testing.T) {
	VMID, err := client.CreateVM("localhost", "7070", "test", "ubuntu22.04", "tiny")
	if err != nil {
		t.Error(err)
	}
	if VMID == 0 {
		t.Error("expected not 0, got ", VMID)
	}
}

// TestGETVM tests the GetVM function.
func TestGetVM(t *testing.T) {
	vm, err := client.GetVM("localhost", "7070", VMID)
	if err != nil {
		t.Error(err)
	}
	if vm.Name != "test" {
		t.Error("expected test, got ", vm.Name)
	}
}

// TestListVMs tests the ListVMs function.
func TestListVMs(t *testing.T) {
	vms, err := client.ListVMs("localhost", "7070")
	if err != nil {
		t.Error(err)
	}
	if len(vms) != 1 {
		t.Error("expected 1 vm, got ", len(vms))
	}
}

// TestDeleteVM tests the DeleteVM function.
func TestDeleteVM(t *testing.T) {
	err := client.DeleteVM("localhost", "7070", VMID)
	if err != nil {
		t.Error(err)
	}
}
