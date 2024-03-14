#!/usr/bin/env bash
ETCDCTL_API=3 etcdctl del --prefix ""
docker rm -f $(docker ps -a -q)