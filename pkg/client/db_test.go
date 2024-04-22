package client_test

import (
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

// TestCreateDB tests the CreateDB function.
func TestCreateDB(t *testing.T) {
	db, err := client.CreateDB(masterHost, masterPort, "test", "mysql")
	if err != nil {
		t.Error(err)
	}
	if db.Name != "test" {
		t.Error("expected test, got ", db.Name)
	}
}

// TestGetDB tests the GetDB function.
func TestGetDB(t *testing.T) {
	db, err := client.GetDB(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
	if db.Name != "test" {
		t.Error("expected test, got ", db.Name)
	}
}

// TestListDBs tests the ListDBs function.
func TestListDBs(t *testing.T) {
	dbs, err := client.ListDBs(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(dbs) != 1 {
		t.Error("expected 1 db, got ", len(dbs))
	}
}

// TestStopDB tests the StopDB function.
func TestStopDB(t *testing.T) {
	err := client.StopDB(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestStartDB tests the StartDB function.
func TestStartDB(t *testing.T) {
	err := client.StartDB(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteDB tests the DeleteDB function.
func TestDeleteDB(t *testing.T) {
	err := client.DeleteDB(masterHost, masterPort, "test")
	if err != nil {
		t.Error(err)
	}
}
