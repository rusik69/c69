#!/usr/bin/env bash
set -x

timeout 5s ssh ubuntu@node0.rusik69.lol "sudo systemctl suspend"
timeout 5s ssh ubuntu@node1.rusik69.lol "sudo systemctl suspend"
timeout 5s ssh ubuntu@node2.rusik69.lol "sudo systemctl suspend"