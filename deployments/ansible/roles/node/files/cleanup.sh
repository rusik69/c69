#!/usr/bin/env bash
docker compose -f ~/docker-compose-`hostname`.yml down
docker system prune -a -f
sudo virsh destroy test; sudo virsh undefine test || true
for i in $(seq 0 10); do
	sudo virsh destroy test$i
	sudo virsh undefine test$i
done
docker ps -aq | xargs docker stop | xargs docker rm || true
rm -f /var/lib/libvirt/images/test* || true