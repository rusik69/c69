#!/bin/bash
set -xe
# Mount the original ISO
mkdir /mnt/iso
mount -o loop /var/lib/libvirt/images/Fedora-Server-netinst-x86_64-39-1.5.iso /mnt/iso

# Copy the contents of the ISO to a new directory
mkdir /tmp/newiso
cp -a /mnt/iso/* /tmp/newiso

# Add your Kickstart file to the new directory
cp /var/lib/libvirt/images/fedora39.ks /tmp/newiso/ks.cfg

# Create a new ISO
mkisofs -o /var/lib/libvirt/images/new-Fedora-Server-netinst-x86_64-39-1.5.iso -b isolinux/isolinux.bin -c isolinux/boot.cat -no-emul-boot -boot-load-size 4 -boot-info-table -R -J -v -T /tmp/newiso

# Clean up
umount /mnt/iso
rm -rf /mnt/iso /tmp/newiso
mv /var/lib/libvirt/images/new-Fedora-Server-netinst-x86_64-39-1.5.iso /var/lib/libvirt/images/Fedora-Server-netinst-x86_64-39-1.5.iso
