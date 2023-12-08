package node

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/rusik69/govnocloud/pkg/node/vm"
	"github.com/spf13/cobra"
)

// CreateVM creates a vm.
func CreateVM(host, port, name, image, flavor string) (int, error) {
	vm := vm.VM{
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

// DeleteVM creates a vm.
func DeleteVM(id int, host, port string) error {
	vm := vm.VM{
		ID: id,
	}
	idString := strconv.Itoa(id)
	url := "http://" + host + ":" + port + "/api/v1/vm/delete/" + idString
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
