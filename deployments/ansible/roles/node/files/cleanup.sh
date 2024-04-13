#!/usr/bin/env bash
docker stop test1-llm
docker rm test1-llm
docker stop test-llm
docker rm test-llm
sudo docker system prune -a -f
sudo virsh destroy test; sudo virsh undefine test || true
for i in $(seq 0 10); do
	sudo virsh destroy test$i
	sudo virsh undefine test$i
	sudo docker stop test$i
	sudo docker rm test$i
done
sudo virsh destroy test-k8s || true
sudo virsh undefine test-k8s || true
sudo virsh destroy test1-k8s || true
sudo virsh undefine test1-k8s || true
for i in $(sudo ls /var/lib/libvirt/images/); do
	sudo rm /var/lib/libvirt/images/$i 
done