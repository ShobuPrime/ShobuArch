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
package main

import (
	"fmt"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	// i "github.com/ShobuPrime/ShobuArch/pkg/install"
	u "github.com/ShobuPrime/ShobuArch/pkg/util"
)

func debug() {

	// Change values as you'd like here, and call relevant functions below
	cdb := conf.Config{
		Format:     "YAML",
		Bootloader: "",
		Kernel:     "linux-zen",
		Hostname:   "",
		Timezone:   "",
		Modules:    []string{``},
		User: conf.User{
			Username: "shobuprime",
			Password: "demo",
			Language: conf.Language{
				Locale:   "en_US.UTF-8",
				CharSet:  "UTF-8",
				Keyboard: "us",
			},
		},
		Storage: conf.Storage{
			SystemDisk:     "",
			SystemDiskID:   "",
			MirrorInstall:  false,
			MirrorDisk:     "",
			MirrorDiskID:   "",
			Filesystem:     "",
			EncryptionKey:  "",
			EncryptionUUID: "",
		},
		Desktop: conf.Desktop{
			Environment: "",
			InstallType: "",
		},
		Pacman: conf.Pacman{
			AUR: conf.AURs{
				Helper:   "",
				Packages: []string{``},
			},
			Packages: []string{
				`keepassxc`,
			},
		},
		Flatpak: conf.Flatpaks{
			Packages: []string{
				`com.brave.Browser`,
				`com.microsoft.Edge`,
				`com.synology.SynologyDrive`,
			},
		},
		PWA: conf.PWAs{
			Sites: []conf.Site{
				{
					Name:        "Messages",
					Description: "Simple, helpful messaging by Google",
					StartUrl:    "https://messages.google.com/web",
					Manifest:    "manifest.json",
				},
				{
					Name:        "YouTube",
					Description: "Share your videos with friends, family, and the world.",
					StartUrl:    "https://www.youtube.com",
					Manifest:    "manifest.webmanifest",
				},
			},
		},
	}

	_ = cdb

	// Call functions you would like to test
	fmt.Println(u.PrettyJson(u.ListUSB()))
	// u.BiometricIDs()
	// i.SetupBiometrics(&cdb)
	// fmt.Println(u.PrettyJson(u.ListPCI()))
	// fmt.Println(u.PrettyJson(u.SecureBootStatus()))
	// i.UserPWAs(&cdb)
}

/*
	--Generate Keyfile: https://wiki.archlinux.org/title/Dm-crypt/Device_encryption#Keyfiles
		dd if=/dev/random of=/.luks/.ata-<drive>-part1 bs=32 count=1
		cryptsetup luksAddKey /dev/sda1 /.luks/<key_file>.key

	--Get UUID of parent partition
		lsblk -dno UUID /dev/sda1

	--Add entry to /etc/crypttab: https://kifarunix.com/automount-luks-encrypted-device-in-linux/
		# <name>       <device>                                     <password>              <options>
		luks_DATA      UUID=<uuid>    /.luks/ata-<drive>-part1

	--Get UUID of encryped partition (Open device if not done so already)
		cryptsetup open /dev/disk/by-id/ata-<drive>-part1 luks_DATA --key-file /.luks/ata-<drive>-part1
		mount -o rw,noatime,compress=zstd:3,ssd,discard,space_cache=v2,commit=120,subvol=@ /dev/mapper/luks_DATA /mnt/ata-<drive>-part1
		lsblk -dno UUID /dev/mapper/luks_DATA

	--Print new fstab and verify if it looks ok before saving
		genfstab -U /

	--Unmount encrypted drive and run the following to verify if done correctly
		mount -av


	VSCode install extensions from CMD line (online):
	code --list-extensions

	golang.go
	ms-python.python
	ms-azuretools.vscode-docker
	shardulm94.trailing-spaces

	code --install-extension <name>

	https://stackoverflow.com/questions/37071388/how-can-i-install-visual-studio-code-extensions-offline#38866913
	https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ms-python/vsextensions/python/2022.19.13351014/vspackage
*/
