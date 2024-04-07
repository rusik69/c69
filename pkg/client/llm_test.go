package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

// TestCreateLLM tests the CreateLLM function.
func TestCreateLLM(t *testing.T) {
	llm, err := client.CreateLLM(masterHost, masterPort, "test", "phi1.5")
	if err != nil {
		t.Error(err)
	}
	if llm.Name != "test" {
		t.Error("expected test, got ", llm.Name)
	}
}

// TestGetLLM tests the GetLLM function.
func TestGetLLM(t *testing.T) {
	llm, err := client.GetLLM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
	if llm.Name != "test" {
		t.Error("expected test, got ", llm.Name)
	}
}

// TestListLLMs tests the ListLLMs function.
func TestListLLMs(t *testing.T) {
	llms, err := client.ListLLMs(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(llms) != 1 {
		t.Error("expected 1 llm, got ", len(llms))
	}
}

// TestStopLLM tests the StopLLM function.
func TestStopLLM(t *testing.T) {
	err := client.StopLLM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestStartLLM tests the StartLLM function.
func TestStartLLM(t *testing.T) {
	err := client.StartLLM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteLLM tests the DeleteLLM function
func TestDeleteLLM(t *testing.T) {
	err := client.DeleteLLM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}
