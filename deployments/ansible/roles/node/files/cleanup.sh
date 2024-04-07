#!/usr/bin/env bash
sudo docker system prune -a -f
docker stop test-llm
docker rm test-llm
sudo virsh destroy test; sudo virsh undefine test || true
for i in $(seq 0 10); do
	sudo virsh destroy test$i
	sudo virsh undefine test$i
	sudo docker stop test$i
	sudo docker rm test$i
done
sudo virsh destroy test-k8s || true
sudo virsh undefine test-k8s || true
for i in $(sudo ls /var/lib/libvirt/images/); do
	sudo rm /var/lib/libvirt/images/$i 
done