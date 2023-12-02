package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rusik69/govnocloud/pkg/node/vm"
	"github.com/spf13/cobra"
)

// CreateVM creates a vm.
func CreateVM(cmd *cobra.Command) error {
	vmName := cmd.PersistentFlags().Lookup("name").Value.String()
	if vmName == "" {
		return errors.New("vm name is required")
	}
	vmImage := cmd.PersistentFlags().Lookup("image").Value.String()
	if vmImage == "" {
		return errors.New("vm image is required")
	}
	vmFlavor := cmd.PersistentFlags().Lookup("flavor").Value.String()
	if vmFlavor == "" {
		return errors.New("vm flavor is required")
	}
	host := cmd.PersistentFlags().Lookup("host").Value.String()
	if host == "" {
		return errors.New("host is required")
	}
	port := cmd.PersistentFlags().Lookup("port").Value.String()
	if port == "" {
		return errors.New("port is required")
	}
	vm := vm.VM{
		Name:   vmName,
		Image:  vmImage,
		Flavor: vmFlavor,
	}
	url := "http://" + host + ":" + port + "/api/v1/vm/create"
	body, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	return err
}

// DeleteVM creates a vm.
func DeleteVM(cmd *cobra.Command) error {
	vmID := cmd.PersistentFlags().Lookup("id").Value.String()
	if vmID == "" {
		return errors.New("vm id is required")
	}
	host := cmd.PersistentFlags().Lookup("host").Value.String()
	if host == "" {
		return errors.New("host is required")
	}
	port := cmd.PersistentFlags().Lookup("port").Value.String()
	if port == "" {
		return errors.New("port is required")
	}
	vm := vm.VM{
		ID: vmID,
	}
	url := "http://" + host + ":" + port + "/api/v1/vm/delete/" + vmID
	body, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	return err
}

// ListVMs creates a vm.
func ListVMs(cmd *cobra.Command) ([]vm.VM, error) {
	host := cmd.PersistentFlags().Lookup("host").Value.String()
	if host == "" {
		return nil, errors.New("host is required")
	}
	port := cmd.PersistentFlags().Lookup("port").Value.String()
	if port == "" {
		return nil, errors.New("port is required")
	}
	url := "http://" + host + ":" + port + "/api/v1/vm/list"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var vms []vm.VM
	err = json.NewDecoder(resp.Body).Decode(&vms)
	return vms, err
}
