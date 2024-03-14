#!/usr/bin/env bash
ETCDCTL_API=3 etcdctl del --prefix ""
docker stop govnocloud-front
docker rm govnocloud-front