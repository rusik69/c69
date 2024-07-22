#!/usr/bin/env bash
set -x

ssh ubuntu@node0.rusik69.lol "sudo systemctl suspend"
ssh ubuntu@node1.rusik69.lol "sudo systemctl suspend"
ssh ubuntu@node2.rusik69.lol "sudo systemctl suspend"