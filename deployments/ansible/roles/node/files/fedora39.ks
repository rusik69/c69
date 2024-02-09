#version=Fedora39
# System authorization information
auth --enableshadow --passalgo=sha512

# Use network installation
url --url="http://download.fedoraproject.org/pub/fedora/linux/releases/39/Everything/x86_64/os/"

# Run the Setup Agent on first boot
firstboot --enable

# System keyboard
keyboard --vckeymap=us --xlayouts='us'

# System language
lang en_US.UTF-8

# Firewall configuration
firewall --enabled --ssh

# Network information
network  --bootproto=dhcp --device=eth0 --onboot=on

# Root password
rootpw root

# System timezone
timezone Europe/Prague --isUtc

# System bootloader configuration
bootloader --append=" crashkernel=auto" --location=mbr --boot-drive=sda

# Clear the Master Boot Record
zerombr

# Partition clearing information
clearpart --all --initlabel

# Disk partitioning information
part / --fstype="ext4" --grow --size=1

%packages
@^minimal-environment
@core
kexec-tools
%end

%addon com_redhat_kdump --enable --reserve-mb='auto'
%end