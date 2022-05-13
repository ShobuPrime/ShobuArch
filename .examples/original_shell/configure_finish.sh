#!/bin/bash

#---------------------------------------------------------------------------------

# umount /mnt/boot/efi
umount /dev/nvme1n1p1
zfs umount -a
zpool export zroot
dd if=/dev/nvme1n1p1  of=/dev/nvme2n1p1
systemctl reboot --firmware