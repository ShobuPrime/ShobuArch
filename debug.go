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
	//i "github.com/ShobuPrime/ShobuArch/pkg/install"
	u "github.com/ShobuPrime/ShobuArch/pkg/util"
)

func debug() {

	// Change values as you'd like here, and call relevant functions below
	cdb := conf.Config{
		Format:   "YAML",
		Bootloader: "",
		Kernel:   "linux-zen",
		Hostname: "",
		Timezone: "",
		Modules: []string{``},
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
			SystemDisk:    "",
			SystemDiskID:  "",
			MirrorInstall: false,
			MirrorDisk:    "",
			MirrorDiskID:  "",
			Filesystem:    "",
			EncryptionKey: "",
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
					Name:        "",
					Description: "",
					StartUrl:    "",
					Manifest:    "",
				},
			},
		},
	}

	_ = cdb

	// Call functions you would like to test
	fmt.Println(u.PrettyJson(u.ListUSB()))
	fmt.Println(u.PrettyJson(u.ListPCI()))
	fmt.Println(u.PrettyJson(u.SecureBootStatus()))
}
