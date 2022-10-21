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
	"os"
	"strings"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	z "github.com/ShobuPrime/ShobuArch/pkg/shell"
)

func PostInstallLogo() {

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
                            PROGRESS: Post-Install
                          SCRIPTHOME: ShobuArch
	-------------------------------------------------------------------------
	`)
}

func PostInstallCleanup(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                		Post-Install Cleanup
	-------------------------------------------------------------------------
	`)

	cmd := []string{`sudo`, `sed`, `-i`, `s/%wheel ALL=(ALL:ALL) NOPASSWD: ALL/# %wheel ALL=(ALL:ALL) NOPASSWD: ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`sudo`, `sed`, `-i`, `s/%sudo ALL=(ALL:ALL) ALL/# %sudo ALL=(ALL:ALL) ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)

	log.Println("Saving config to user profile")
	config_bytes, err := os.ReadFile(fmt.Sprintf("./shobuarch_config.%s", strings.ToLower(c.Format)))
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(fmt.Sprintf("/mnt/home/%s/Desktop/shobuarch_config.%s", c.User.Username, strings.ToLower(c.Format)), config_bytes, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Saving log to user profile")
	log_bytes, err := os.ReadFile("./shobuarch.log")
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(fmt.Sprintf("/mnt/home/%s/Desktop/shobuarch.log", c.User.Username), log_bytes, 0755)
	if err != nil {
		log.Fatalln(err)
	}
}

func PostInstallUnmount(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                		Unmounting Drives
	-------------------------------------------------------------------------
	`)

	var disk1_part1, disk1_part2, disk2_part1 string
	var cmd_list []string

	switch strings.HasPrefix(c.Storage.SystemDisk, "/dev/nvme") {
	case true:
		disk1_part1 = fmt.Sprintf(`%vp1`, c.Storage.SystemDisk)
		disk1_part2 = fmt.Sprintf(`%vp2`, c.Storage.SystemDisk)
	case false:
		disk1_part1 = fmt.Sprintf(`%v1`, c.Storage.SystemDisk)
		disk1_part2 = fmt.Sprintf(`%v2`, c.Storage.SystemDisk)
	}

	switch c.Storage.MirrorInstall {
	case true:
		switch strings.HasPrefix(c.Storage.MirrorDisk, "/dev/nvme") {
		case true:
			disk2_part1 = fmt.Sprintf(`%vp1`, c.Storage.MirrorDisk)
		case false:
			disk2_part1 = fmt.Sprintf(`%v1`, c.Storage.MirrorDisk)
		}

		log.Println("Cloning EFI Partitions")
		cmd_list = append(cmd_list,
			fmt.Sprintf(`dd if=%v of=%v`, disk1_part1, disk2_part1),
		)
	}

	cmd_list = append(cmd_list,
		fmt.Sprintf(`umount -A %s`, disk1_part1),
	)

	switch c.Storage.Filesystem {
	case "luks":
		cmd_list = append(cmd_list,
			`umount -A /dev/mapper/luks_ROOT`,
		)
	case "zfs":
		cmd_list = append(cmd_list,
			`zfs umount -a`,
			`zpool export zroot`,
		)
	default:
		cmd_list = append(cmd_list,
			fmt.Sprintf(`umount -A %s`, disk1_part2),
		)
	}

	for i := range cmd_list {
		z.Shell(&cmd_list[i])
	}
}
