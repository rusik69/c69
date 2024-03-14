#!/usr/bin/env bash
ETCDCTL_API=3 etcdctl del --prefix ""
docker stop govnocloud-front || true
docker rm govnocloud-front || true