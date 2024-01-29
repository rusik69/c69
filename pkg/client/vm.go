package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rusik69/govnocloud/pkg/types"
)

// CreateVM creates a vm.
func CreateVM(host, port, name, image, flavor string) (int, error) {
	vm := types.VM{
		Name:   name,
		Image:  image,
		Flavor: flavor,
	}
	url := "http://" + host + ":" + port + "/api/v1/vms"
	body, err := json.Marshal(vm)
	if err != nil {
		return 0, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New(string(bodyText))
	}
	err = json.Unmarshal(bodyText, &vm)
	if err != nil {
		return 0, err
	}
	return vm.ID, nil
}

// DeleteVM deletes a vm.
func DeleteVM(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/vm/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	return err
}

// StartVM starts a vm.
func StartVM(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/vmstart/" + name
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	return err
}

// StopVM stops a vm.
func StopVM(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/vmstop/" + name
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	return err
}

// ListVMs lists vms.
func ListVMs(host, port string) ([]types.VM, error) {
	url := "http://" + host + ":" + port + "/api/v1/vms"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var vms []types.VM
	err = json.NewDecoder(resp.Body).Decode(&vms)
	return vms, err
}

// GetVM gets a vm.
func GetVM(host, port, name string) (types.VM, error) {
	vm := types.VM{}
	url := "http://" + host + ":" + port + "/api/v1/vm/" + name
	resp, err := http.Get(url)
	if err != nil {
		return vm, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&vm)
	return vm, err
}
