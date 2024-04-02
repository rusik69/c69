package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// CreateK8S creates a k8s cluster.
func CreateK8S(host, port, name, flavor string) (types.K8S, error) {
	k8s := types.K8S{
		Name:   name,
		Flavor: flavor,
	}
	url := "http://" + host + ":" + port + "/api/v1/k8s"
	body, err := json.Marshal(k8s)
	if err != nil {
		return types.K8S{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return types.K8S{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.K8S{}, err
	}
	if resp.StatusCode != 200 {
		return types.K8S{}, errors.New(string(bodyText))
	}
	logrus.Println("bodyText", string(bodyText))
	var newK8S types.K8S
	err = json.Unmarshal(bodyText, &newK8S)
	logrus.Println("new K8S", newK8S)
	if err != nil {
		return types.K8S{}, err
	}
	return k8s, nil
}

// GetK8S gets a k8s cluster.
func GetK8S(host, port, name string) (types.K8S, error) {
	url := "http://" + host + ":" + port + "/api/v1/k8s/" + name
	resp, err := http.Get(url)
	if err != nil {
		return types.K8S{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.K8S{}, err
	}
	if resp.StatusCode != 200 {
		return types.K8S{}, errors.New(string(bodyText))
	}
	var k8s types.K8S
	err = json.Unmarshal(bodyText, &k8s)
	if err != nil {
		return types.K8S{}, err
	}
	return k8s, nil
}

// DeleteK8S deletes a k8s cluster.
func DeleteK8S(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/k8s/" + name
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
	return nil
}

// ListK8S lists k8s clusters.
func ListK8S(host, port string) ([]types.K8S, error) {
	url := "http://" + host + ":" + port + "/api/v1/k8s"
	resp, err := http.Get(url)
	if err != nil {
		return []types.K8S{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return []types.K8S{}, err
	}
	if resp.StatusCode != 200 {
		return []types.K8S{}, errors.New(string(bodyText))
	}
	var k8sList []types.K8S
	err = json.Unmarshal(bodyText, &k8sList)
	if err != nil {
		return []types.K8S{}, err
	}
	return k8sList, nil
}

// StartK8S starts a k8s cluster.
func StartK8S(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/k8sstart/" + name
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
	return nil
}

// StopK8S stops a k8s cluster.
func StopK8S(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/k8sstop/" + name
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
	return nil
}

// GetKubeconfig gets the kubeconfig of a k8s cluster.
func GetKubeconfig(host, port, name string) (string, error) {
	url := "http://" + host + ":" + port + "/api/v1/k8s/" + name + "/kubeconfig"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(string(bodyText))
	}
	bodyJSON := map[string]string{}
	err = json.Unmarshal(bodyText, &bodyJSON)
	if err != nil {
		return "", err
	}
	decoded, err := base64.StdEncoding.DecodeString(bodyJSON["kubeconfig"])
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
