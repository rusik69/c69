package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/rusik69/govnocloud/pkg/types"
)

// CreateDB creates a database.
func CreateDB(host, port, name, dbType string) (types.DB, error) {
	db := types.DB{
		Name: name,
		Type: dbType,
	}
	url := "http://" + host + ":" + port + "/api/v1/db"
	body, err := json.Marshal(db)
	if err != nil {
		return types.DB{}, err
	}
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return types.DB{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return types.DB{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.DB{}, err
	}
	if resp.StatusCode != 200 {
		return types.DB{}, errors.New(string(bodyText))
	}
	var newDB types.DB
	err = json.Unmarshal(bodyText, &newDB)
	if err != nil {
		return types.DB{}, err
	}
	return db, nil
}

// GetDB gets a database.
func GetDB(host, port, name string) (types.DB, error) {
	url := "http://" + host + ":" + port + "/api/v1/db/" + name
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.DB{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return types.DB{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.DB{}, err
	}
	if resp.StatusCode != 200 {
		return types.DB{}, errors.New(string(bodyText))
	}
	var db types.DB
	err = json.Unmarshal(bodyText, &db)
	if err != nil {
		return types.DB{}, err
	}
	return db, nil
}

// DeleteDB deletes a database.
func DeleteDB(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/db/" + name
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(bodyText))
	}
	return nil
}

// ListDBs lists databases.
func ListDBs(host, port string) ([]types.DB, error) {
	url := "http://" + host + ":" + port + "/api/v1/db"
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []types.DB{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []types.DB{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return []types.DB{}, err
	}
	if resp.StatusCode != 200 {
		return []types.DB{}, errors.New(string(bodyText))
	}
	var dbs []types.DB
	err = json.Unmarshal(bodyText, &dbs)
	if err != nil {
		return []types.DB{}, err
	}
	return dbs, nil
}

// StartDB starts a database.
func StartDB(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/dbstart/" + name
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(bodyText))
	}
	return nil
}

// StopDB stops a database.
func StopDB(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/dbstop/" + name
	client := &http.Client{
		Timeout: 300 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(bodyText))
	}
	return nil
}
