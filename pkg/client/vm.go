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
