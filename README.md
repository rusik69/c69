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