/*
* Automated Arch Linux Tools
* Copyright (C) 2022  Anthony Dardano (ShobuPrime)
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* ShobuArch is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package install

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	z "github.com/ShobuPrime/ShobuArch/pkg/shell"
	u "github.com/ShobuPrime/ShobuArch/pkg/util"
)

func PreInstallLogo() {

	log.Println(`
	-------------------------------------------------------------------------
	███████╗██╗  ██╗ ██████╗ ██████╗ ██╗   ██╗ █████╗ ██████╗  ██████╗██╗  ██╗
	██╔════╝██║  ██║██╔═══██╗██╔══██╗██║   ██║██╔══██╗██╔══██╗██╔════╝██║  ██║
	███████╗███████║██║   ██║██████╔╝██║   ██║███████║██████╔╝██║     ███████║
	╚════██║██╔══██║██║   ██║██╔══██╗██║   ██║██╔══██║██╔══██╗██║     ██╔══██║
	███████║██║  ██║╚██████╔╝██████╔╝╚██████╔╝██║  ██║██║  ██║╚██████╗██║  ██║
	╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
	-------------------------------------------------------------------------
                        Automated Arch Linux (Installer)
                            PROGRESS: Pre-Install
                          SCRIPTHOME: ShobuArch
	-------------------------------------------------------------------------
	`)
}

func Prerequisites(c *conf.Config) {
	fmt.Println(`
	-------------------------------------------------------------------------
                    Installing Prerequisites
	-------------------------------------------------------------------------
	`)

	prereqs := `pacman -Syy --noconfirm --needed btrfs-progs glibc gptfdisk sbctl sed sudo`
	z.Shell(&prereqs)
}

func Mirrors(c *conf.Config) {

	// cmd := `(curl -4 ifconfig.co/country-iso)`
	// iso, _ := z.Shell(&cmd)

	// Hard-coded ISO
	// To-do: Fix receiving ISO from ifconfig
	iso := "US"

	log.Printf(`
		-------------------------------------------------------------------------
					Setting up %s mirrors for faster downloads
		-------------------------------------------------------------------------
	`, iso)

	cmd_list := []string{}

	cmd_list = append(cmd_list,
		`timedatectl set-ntp true`,
		`pacman -Syy --noconfirm archlinux-keyring`,
		`pacman -Syy --noconfirm --needed pacman-contrib terminus-font`,
		`sed -n 's/^#ParallelDownloads/ParallelDownloads/g' /etc/pacman.conf`,
		`pacman -Syy --noconfirm --needed reflector rsync`,
		`cp /etc/pacman.d/mirrorlist /etc/pacman.d/mirrorlist.bak`,
		fmt.Sprintf(`reflector --protocol https --country %s --latest 20 --sort rate --ipv6 --fastest 5 --save /etc/pacman.d/mirrorlist`, iso),
	)

	for i := range cmd_list {
		z.Shell(&cmd_list[i])
	}
}

func FormatDisks(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
				Creating Filesystems
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{}

	var mount_options, disk1_part1, disk1_part2, disk2_part1, disk2_part2 string

	cmd_list = append(cmd_list,
		fmt.Sprintf(`wipefs -a %v`, c.Storage.SystemDisk),
		fmt.Sprintf(`sgdisk --zap-all %v`, c.Storage.SystemDisk),
		fmt.Sprintf(`sgdisk -o %v`, c.Storage.SystemDisk),
		fmt.Sprintf(`sgdisk -n 1::+550M --typecode=1:ef00 --change-name=1:'EFIBOOT' %v`, c.Storage.SystemDisk),
		fmt.Sprintf(`sgdisk -n 2:: --typecode=2:bf00 --change-name=2:'ROOT' %v`, c.Storage.SystemDisk),
	)

	if strings.HasPrefix(c.Storage.SystemDisk, "/dev/nvme") {
		disk1_part1 = fmt.Sprintf(`%vp1`, c.Storage.SystemDisk)
		disk1_part2 = fmt.Sprintf(`%vp2`, c.Storage.SystemDisk)
	} else {
		disk1_part1 = fmt.Sprintf(`%v1`, c.Storage.SystemDisk)
		disk1_part2 = fmt.Sprintf(`%v2`, c.Storage.SystemDisk)
	}

	cmd_list = append(cmd_list, fmt.Sprintf(`mkfs.vfat -F32 -n "EFIBOOT" %v`, disk1_part1))

	if c.Storage.MirrorInstall {

		cmd_list = append(cmd_list,
			fmt.Sprintf(`wipefs -a %v`, c.Storage.MirrorDisk),
			fmt.Sprintf(`sgdisk --zap-all %v`, c.Storage.MirrorDisk),
			fmt.Sprintf(`sgdisk -o %v`, c.Storage.MirrorDisk),
			fmt.Sprintf(`sgdisk -n 1::+550M --typecode=1:ef00 --change-name=1:'EFIBOOT' %v`, c.Storage.MirrorDisk),
			fmt.Sprintf(`sgdisk -n 2:: --typecode=2:bf00 --change-name=2:'ROOT' %v`, c.Storage.MirrorDisk),
		)

		if strings.HasPrefix(c.Storage.SystemDisk, "/dev/nvme") {
			disk2_part1 = fmt.Sprintf(`%vp1`, c.Storage.MirrorDisk)
			disk2_part2 = fmt.Sprintf(`%vp2`, c.Storage.MirrorDisk)
		} else {
			disk2_part1 = fmt.Sprintf(`%v1`, c.Storage.MirrorDisk)
			disk2_part2 = fmt.Sprintf(`%v2`, c.Storage.MirrorDisk)
		}

		cmd_list = append(cmd_list, fmt.Sprintf(`mkfs.vfat -F32 -n "EFIBOOT" %v`, disk2_part1))
	}

	log.Println("Detecting SSD...")
	if c.Storage.SystemDiskRota || strings.Contains(c.Storage.SystemDiskID, "SSD") {
		// MOUNT_OPTIONS if SSD: noatime,compress=zstd,ssd,commit=120
		mount_options = `discard=async,noatime,compress=zstd,ssd,commit=120`
		log.Printf("SSD detected! MOUNT_OPTIONS=%v\n", mount_options)
	} else {
		// MOUNT_OPTIONS if HDD: noatime,compress=zstd,commit=120
		mount_options = `noatime,compress=zstd,commit=120`
		log.Printf("HDD detected! MOUNT_OPTIONS=%v\n", mount_options)
	}

	switch c.Storage.Filesystem {
	case "btrfs":
		switch c.Storage.MirrorInstall {
		case true:
			cmd_list = append(cmd_list,
				fmt.Sprintf(`mkfs.btrfs -L ROOT -m raid1 -d raid1 %v %v -f`, disk1_part2, disk2_part2),
			)
		case false:
			cmd_list = append(cmd_list,
				fmt.Sprintf(`mkfs.btrfs -L ROOT %v -f`, disk1_part2),
			)
		}
		cmd_list = append(cmd_list,
			fmt.Sprintf(`mount -t btrfs %v /mnt`, disk1_part2),
			`btrfs subvolume create /mnt/@`, // createsubvolumes
			`btrfs subvolume create /mnt/@home`,
			`btrfs subvolume create /mnt/@var`,
			`btrfs subvolume create /mnt/@tmp`,
			`btrfs subvolume create /mnt/@.snapshots`,
			`umount /mnt`,
			fmt.Sprintf(`mount -o %v,subvol=@ %v /mnt`, mount_options, disk1_part2),
			`mkdir -p /mnt/{home,var,tmp,.snapshots}`,
			fmt.Sprintf(`mount -o %v,subvol=@home %v /mnt/home`, mount_options, disk1_part2),
			fmt.Sprintf(`mount -o %v,subvol=@tmp %v /mnt/tmp`, mount_options, disk1_part2),
			fmt.Sprintf(`mount -o %v,subvol=@var %v /mnt/var`, mount_options, disk1_part2),
			fmt.Sprintf(`mount -o %v,subvol=@.snapshots %v /mnt/.snapshots`, mount_options, disk1_part2),
			`mkdir -p /mnt/boot/EFI`,
			fmt.Sprintf(`mount -t vfat %s /mnt/boot/`, disk1_part1),
		)
	case "ext4":
		cmd_list = append(cmd_list,
			fmt.Sprintf(`mkfs.ext4 -L ROOT %v -f`, disk1_part2),
			fmt.Sprintf(`mount -t ex4 %v /mnt`, disk1_part2),
		)
	case "luks":
		cmd_list = append(cmd_list,
			fmt.Sprintf(`echo -n "%v" | cryptsetup -y -v luksFormat --type=luks2 --label=ROOT %v -`, c.Storage.EncryptionKey, disk1_part2),
			fmt.Sprintf(`echo -n "%v" | cryptsetup -v luksOpen %v luks_ROOT -`, c.Storage.EncryptionKey, disk1_part2),
			`lsblk`,
			// fmt.Sprintf(`mkfs.btrfs -L ROOT %v -f`, disk1_part2),
			`mkfs.btrfs /dev/mapper/luks_ROOT -f`,
			// fmt.Sprintf(`mount -t btrfs %v /mnt`, disk1_part2),
			`mount /dev/mapper/luks_ROOT /mnt`,
			`btrfs subvolume create /mnt/@`, // subvolumesetup
			`btrfs subvolume create /mnt/@home`,
			`btrfs subvolume create /mnt/@var`,
			`btrfs subvolume create /mnt/@tmp`,
			`btrfs subvolume create /mnt/@.snapshots`,
			`umount /mnt`,
			fmt.Sprintf(`mount -o %v,subvol=@ /dev/mapper/luks_ROOT /mnt`, mount_options), // mountallsubvol
			`mkdir -p /mnt/{home,var,tmp,.snapshots}`,
			fmt.Sprintf(`mount -o %v,subvol=@home /dev/mapper/luks_ROOT /mnt/home`, mount_options),
			fmt.Sprintf(`mount -o %v,subvol=@tmp /dev/mapper/luks_ROOT /mnt/tmp`, mount_options),
			fmt.Sprintf(`mount -o %v,subvol=@var /dev/mapper/luks_ROOT /mnt/var`, mount_options),
			fmt.Sprintf(`mount -o %v,subvol=@.snapshots /dev/mapper/luks_ROOT /mnt/.snapshots`, mount_options),
			`mkdir -p /mnt/boot/EFI`,
			fmt.Sprintf(`mount -t vfat %s /mnt/boot/`, disk1_part1),
		)
	case "zfs":
		zfs_cmd := `zpool create -f -o ashift=12	\
						-O acltype=posixacl			\
						-O relatime=on				\
						-O xattr=sa					\
						-O dnodesize=legacy			\
						-O normalization=formD		\
						-O mountpoint=none			\
						-O canmount=off				\
						-O devices=off				\
						-R /mnt						\
						-O compression=zstd			\
						-O dedup=on					\
					zroot`

		switch c.Storage.MirrorInstall {
		case true:
			cmd_list = append(cmd_list,
				`modprobe zfs`,
				fmt.Sprintf(`%v mirror %v %v`, zfs_cmd, disk1_part2, disk2_part2),
			)
		case false:
			cmd_list = append(cmd_list,
				`modprobe zfs`,
				fmt.Sprintf(`%v %v`, zfs_cmd, disk1_part2),
			)
		}

		cmd_list = append(cmd_list,
			`zpool status`,
			`zfs create -o mountpoint=none zroot/data`,
			`zfs create -o mountpoint=none zroot/ROOT`,
			`zfs create -o mountpoint=/ -o canmount=noauto zroot/ROOT/default`,
			`zfs create -o mountpoint=/home zroot/data/home`,
			`zfs create -o mountpoint=/var -o canmount=off zroot/var`, // Future: Figure out how to add zfs create zroot/var/log without freeze on shutdown
			`zfs create -o mountpoint=/var/lib -o canmount=off zroot/var/lib`,
			`zfs create zroot/var/lib/libvirt`,
			`zfs create zroot/var/lib/docker`,
			`zpool export zroot`,
			fmt.Sprintf(`zpool import -d %v -R /mnt zroot -N`, disk1_part2),
			`zfs mount zroot/ROOT/default`,
			`zfs mount -a`,
			`zpool set bootfs=zroot/ROOT/default zroot`,
			`zpool set cachefile=/etc/zfs/zpool.cache zroot`,
			`mkdir -p /mnt/{etc/zfs,boot}`,
			`cp /etc/zfs/zpool.cache /mnt/etc/zfs/zpool.cache`,
			fmt.Sprintf(`mount %v /mnt/boot`, disk1_part1),
		)
	}

	cmd_list = append(cmd_list, `df -HT`)

	log.Println("Formatting Disks...")
	for i := range cmd_list {
		z.Shell(&cmd_list[i])
	}
}

func ArchInstall(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
				Arch Install on Main Drive
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{}

	// Temp: Adding openssl-1.1 to fix libcrypto.so.1.1 error
	// pacman: error while loading shared libraries: libcrypto.so.1.1: cannot open shared object file: No such file or directory
	cmd_list = append(cmd_list,
		fmt.Sprintf(`pacstrap /mnt --noconfirm --needed archlinux-keyring base base-devel dkms libnewt %s linux-firmware %s-docs %s-headers nano openssl-1.1 packagekit sbctl sed sudo vim sudo wget`, c.Kernel, c.Kernel, c.Kernel),
		`echo "keyserver hkp://keyserver.ubuntu.com" >> /mnt/etc/pacman.d/gnupg/gpg.conf`,
		`cp /etc/pacman.d/mirrorlist /mnt/etc/pacman.d/mirrorlist`,
	)

	switch c.Storage.Filesystem {
	case "zfs":
		log.Printf("Generating fstab for %s\n", c.Storage.Filesystem)
		cmd_list = append(cmd_list,
			`genfstab -L /mnt >> /mnt/etc/fstab`,
		)
		log.Println("Commenting out zroot entries in fstab. ZFS handles this by itself")
		cmd_list = append(cmd_list,
			`sed -i 's/zroot/# zroot/g' /mnt/etc/fstab`,
			`sed -i 's/# # zroot/# zroot/g' /mnt/etc/fstab`,
			`cat /mnt/etc/fstab`,
		)
	default:
		log.Printf("Generating fstab for %s\n", c.Storage.Filesystem)
		cmd_list = append(cmd_list,
			`genfstab -U -p /mnt >> /mnt/etc/fstab`,
		)
	}

	for i := range cmd_list {
		z.Shell(&cmd_list[i])
	}

	fstab_path := filepath.Join("/", "mnt", "etc")
	fstab_file := `fstab`
	u.ReadFile(&fstab_path, &fstab_file)
}
