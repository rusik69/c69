#!/usr/bin/env bash
sudo docker system prune -a -f
sudo virsh destroy test; sudo virsh undefine test || true
for i in $(seq 0 10); do
	sudo virsh destroy test$i
	sudo virsh undefine test$i
	sudo docker stop test$i
	sudo docker rm test$i
done
for i in $(sudo ls /var/lib/libvirt/images/); do
	sudo rm /var/lib/libvirt/images/$i 
done
sudo umount /nbd0 || true