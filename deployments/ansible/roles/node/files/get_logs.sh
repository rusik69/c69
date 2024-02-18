#!/usr/bin/env bash
journalctl _SYSTEMD_INVOCATION_ID=`systemctl show -p InvocationID --value govnocloud-node.service`