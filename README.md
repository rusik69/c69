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

All nodes and masters should be connected using tailscale.

## Usage

govnocloud-client can be used to interact with govnocloud.
Web interface is available on port 8080 of the master instance.

## Container flavors:
| Flavor  | ID  | MilliCPUs | RAM  |
| ------- | --- | --------- | ---- |
| tiny    |  0  | 128       | 256  |
| small   |  1  | 256       | 512  |
| medium  |  2  | 512       | 1024 |
| large   |  3  | 1024      | 2048 |
| xlarge  |  4  | 2048      | 4096 |
| 2xlarge |  5  | 4096      | 8192 |

## VM flavors:
| Flavor | ID  | MilliCPUs | RAM  | Disk |
| ------ | --- | --------- | ---- | ---- |
| small  |  0  | 512       | 1024 | 8    |
| medium |  1  | 1024      | 2048 | 8    |
| large  |  2  | 2048      | 4096 | 16   |
| xlarge |  3  | 4096      | 8192 | 32   |