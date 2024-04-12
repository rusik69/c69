package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
)

var llm types.LLM

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

// TestGenerateLLM tests the GenerateLLM function
func TestGenerateLLM(t *testing.T) {
	res, err := client.GenerateLLM(masterHost, masterPort, "test", "hello")
	if err != nil {
		t.Error(err)
	}
	if res == "" {
		t.Error("expected not empty string")
	}
}

// TestGetLLM tests the GetLLM function.
func TestGetLLM(t *testing.T) {
	res, err := client.GetLLM(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
	if res.Name != "test" {
		t.Error("expected test, got ", res.Name)
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
