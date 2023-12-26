package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/rusik69/govnocloud/pkg/types"
)

// CreateVM creates a vm.
func CreateVM(host, port, name, image, flavor string) (int, error) {
	vm := types.VM{
		Name:   name,
		Image:  image,
		Flavor: flavor,
	}
	url := "http://" + host + ":" + port + "/api/v1/vm/create"
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
func DeleteVM(host, port string, id int) error {
	vm := types.VM{
		ID: id,
	}
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/vm/" + idString
	body, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	return err
}

// StartVM starts a vm.
func StartVM(host, port string, id int) error {
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/vm/start/" + idString
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
func StopVM(host, port string, id int) error {
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/vm/stop/" + idString
	body, err = http.Get(url)
	if err != nil {
		
	return err
}

// ListVMs lists vms.
func ListVMs(host, port string) ([]types.VM, error) {
	url := "http://" + host + ":" + port + "/api/v1/vm/list"
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
func GetVM(host, port string, id int) (types.VM, error) {
	vm := types.VM{
		ID: id,
	}
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/vm/" + idString
	resp, err := http.Get(url)
	if err != nil {
		return vm, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&vm)
	return vm, err
}
