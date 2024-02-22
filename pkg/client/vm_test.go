package client_test

import (
	"fmt"
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/sirupsen/logrus"
)

// TestCreateVM tests the CreateVM function.
func TestCreateVM(t *testing.T) {
	vm, err := client.CreateVM(masterHost, masterPort, "test", "ubuntu22.04", "small")
	if err != nil {
		t.Error(err)
	}
	logrus.Println(vm)
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
	fmt.Println("RunVMs")
	for i := 0; i < 8; i++ {
		vmName := "test" + fmt.Sprintf("%d", i)
		_, err := client.CreateVM(masterHost, masterPort, vmName, "ubuntu22.04", "small")
		if err != nil {
			fmt.Println(err)
		}
	}
}
