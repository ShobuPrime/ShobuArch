#!/bin/bash

echo "Installing Arch on ZFS\n"

# Ensure network connectivity
# ip a

# List block devices
# lsblk

# List devices by id
# ls -lah /dev/disk/by-id/

# Manual Partition disk (using values from 5700g ITX)
# gdisk /dev/nvme1n1
# gdisk /dev/nvme2n1

# Command: n
# Partition number: \n (1)
# First sector: \n
# Last sector: +600M
# Hex code or GUID: ef00

# Command: n
# Partition number: \n (2)
# First sector: \n
# Last sector: \n
# Hex code or GUID: bf00

# Command: p

# Command: w
# Command: Y

# sgdisk --zap-all /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>
# sgdisk -n1:0:+600M -t1:ef00 /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>
# sgdisk -n2:0:0 -t2:bf00 /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>

wipefs -a /dev/nvme1n1
sgdisk --zap-all /dev/nvme1n1
sgdisk -o /dev/nvme1n1
sgdisk -n 1::+600M -t 1:ef00 /dev/nvme1n1
sgdisk -n 2:: -t 2:bf00 /dev/nvme1n1

# sgdisk --zap-all /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>
# sgdisk -n1:0:+600M -t1:ef00 /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>
# sgdisk -n2:0:0 -t2:bf00 /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>

wipefs -a /dev/nvme2n1
sgdisk --zap-all /dev/nvme2n1
sgdisk -o /dev/nvme2n1
sgdisk -n 1::+600M -t 1:ef00 /dev/nvme2n1
sgdisk -n 2:: -t 2:bf00 /dev/nvme2n1

# Make EFI partition fat32
# mkfs.vfat /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>-part1
mkfs.vfat /dev/nvme1n1p1

# mkfs.vfat /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>-part1
mkfs.vfat /dev/nvme2n1p1

# Load zfs modules
modprobe zfs

# Check that modules are loaded
lsmod | grep -i zfs

# Create zpool
# zpool create -f -o ashift=12 \
#     -O acltype=posixacl \
#     -O relatime=on \
#     -O xattr=sa \
#     -O dnodesize=legacy \
#     -O normalization=formD \
#     -O mountpoint=none \
#     -O canmount=off \
#     -O devices=off \
#     -R /mnt \
#     -O compression=zstd \
#     zroot mirror /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>-part2 /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>-part2

zpool create -f -o ashift=12 \
    -O acltype=posixacl \
    -O relatime=on \
    -O xattr=sa \
    -O dnodesize=legacy \
    -O normalization=formD \
    -O mountpoint=none \
    -O canmount=off \
    -O devices=off \
    -R /mnt \
    -O compression=zstd \
    -O dedup=on \
    zroot mirror /dev/nvme1n1p2 /dev/nvme2n1p2

zpool status
zfs create -o mountpoint=none zroot/data
zfs create -o mountpoint=none zroot/ROOT
zfs create -o mountpoint=/ -o canmount=noauto zroot/ROOT/default
zfs create -o mountpoint=/home zroot/data/home
zfs create -o mountpoint=/var -o canmount=off zroot/var
# zfs create zroot/var/log
zfs create -o mountpoint=/var/lib -o canmount=off zroot/var/lib
zfs create zroot/var/lib/libvirt
zfs create zroot/var/lib/docker
zpool export zroot
# zpool import -d /dev/disk/by-id/nvme-Samsung_970_PRO_512GB_<Serial>-part2 -R /mnt zroot -N
# zpool import -d /dev/nvme1n1p2 -R /mnt zroot -N
zpool import -d /dev/nvme1n1p2 -R /mnt zroot -N
zfs mount zroot/ROOT/default
zfs mount -a

zpool set bootfs=zroot/ROOT/default zroot
zpool set cachefile=/etc/zfs/zpool.cache zroot
mkdir -p /mnt/{etc/zfs,boot}
cp /etc/zfs/zpool.cache /mnt/etc/zfs/zpool.cache

mount /dev/nvme1n1p1 /mnt/boot

pacman -Syu --noconfirm

# amd/intel-ucode depending on system
pacstrap /mnt base base-devel dkms git amd-ucode linux linux-firmware linux-headers nano vim --needed --noconfirm

genfstab -U -p /mnt >> /mnt/etc/fstab

chmod +x ./configure_arch.sh
cp ./configure_arch.sh /mnt/configure_arch.sh
chmod +x ./configure_user.sh
cp ./configure_user.sh /mnt/configure_user.sh
zpool status
# echo "Running 'arch-chroot /mnt' -- please run './configure_arch.sh' next"
# arch-chroot /mnt

# ( arch-chroot /mnt /usr/bin/runuser /configure_arch.sh )|& tee 2-arch.log
# ( arch-chroot /mnt /usr/bin/runuser -u shobuprime -- /configure_user.sh )|& tee 2-user.log

arch-chroot /mnt /configure_arch.sh
# arch-chroot /mnt /usr/bin/runuser -u shobuprime -- /configure_user.sh