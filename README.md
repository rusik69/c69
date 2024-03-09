# Govnocloud

## Description

Govnocloud is a simple and shitty cloud.
It currently supports running vms, docker containers and file storage

## Deployment

To install Govnocloud, simply run:
```bash
bin/govnocloud-deploy-linux-amd64 --master master_host --nodes node0_host, node1_host
```
Master node should be specified to run govnocloud master, database and web interface.
Nodes are used to run vms/containers

## Requirements

SSH key auth should be configured on master on nodes using root user, or another user can be specified using --user flag.

Ansible should be installed on the host used to deploy govnocloud.

## Usage

govnocloud-client can be used to interact with govnocloud.
Web interface is available on port 8080 of the master instance.

Container flavors:
```go
var ContainerFlavors = map[string]ContainerFlavor{
	"tiny": ContainerFlavor{
		ID:        0,
		MilliCPUs: 128,
		RAM:       256,
	},
	"small": ContainerFlavor{
		ID:        1,
		MilliCPUs: 256,
		RAM:       512,
	},
	"medium": ContainerFlavor{
		ID:        2,
		MilliCPUs: 512,
		RAM:       1024,
	},
	"large": ContainerFlavor{
		ID:        3,
		MilliCPUs: 1024,
		RAM:       2048,
	},
	"xlarge": ContainerFlavor{
		ID:        4,
		MilliCPUs: 2048,
		RAM:       4096,
	},
	"2xlarge": ContainerFlavor{
		ID:        5,
		MilliCPUs: 4096,
		RAM:       8192,
	},
}
```

VM flavors:
```go
var VMFlavors = map[string]VMFlavor{
	"small": VMFlavor{
		ID:        0,
		MilliCPUs: 512,
		RAM:       1024,
		Disk:      8,
	},
	"medium": VMFlavor{
		ID:        1,
		MilliCPUs: 1024,
		RAM:       2048,
		Disk:      8,
	},
	"large": VMFlavor{
		ID:        2,
		MilliCPUs: 2048,
		RAM:       4096,
		Disk:      16,
	},
	"xlarge": VMFlavor{
		ID:        3,
		MilliCPUs: 4096,
		RAM:       8192,
		Disk:      32,
	},
}
```