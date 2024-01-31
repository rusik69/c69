package client_test

import (
	"fmt"
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

// TestCreateVM tests the CreateVM function.
func TestCreateVM(t *testing.T) {
	VMID, err := client.CreateVM(masterHost, masterPort, "test", "ubuntu22.04", "tiny")
	if err != nil {
		t.Error(err)
	}
	if VMID == 0 {
		t.Error("expected not 0, got ", VMID)
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

// RunVMs runs vms.
func RunVMs() {
	for i := 0; i < 10; i++ {
		vmName := "test" + fmt.Sprintf("%d", i)
		client.CreateVM(masterHost, masterPort, vmName, "ubuntu22.04", "tiny")
	}
}

// RemoveVMs removes vms.
func RemoveVMs() {
	for i := 0; i < 10; i++ {
		vmName := "test" + fmt.Sprintf("%d", i)
		client.DeleteVM(masterHost, masterPort, vmName)
	}
}
