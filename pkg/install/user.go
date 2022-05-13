/*
* Automated Arch Linux tools
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

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	z "github.com/ShobuPrime/ShobuArch/pkg/shell"
)

func UserLogo() {

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
                            PROGRESS: User
                          SCRIPTHOME: ShobuArch
	-------------------------------------------------------------------------
	`)
}

func UserLocale(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                    Enforcing Desktop Environment Locale
	-------------------------------------------------------------------------
	`)

	// As other desktop environments are played around with, their locale configs will be modified as needed
	switch c.Desktop.Environment {
	case "kde":
		cmd_list := []string{
			`touch`,
			fmt.Sprintf(`/home/%s/.config/plasma-localerc`, c.User.Username),
		}
		z.Arch_chroot(&cmd_list, true, c)

		cmd_list = []string{
			`awk`,
			fmt.Sprintf(`BEGIN{ printf "[Formats]\nLANG=%s\n" >> "/home/%s/.config/plasma-localerc" }`, c.User.Language.Locale, c.User.Username),
		}
		z.Arch_chroot(&cmd_list, true, c)
	}
}

func UserPackages(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Installing AUR Packages
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{`yay`, `-Syu`, `--needed`, `--noconfirm`}
	for i := range c.Pacman.AUR.Packages {
		cmd_list = append(cmd_list, c.Pacman.AUR.Packages[i])
	}
	z.Arch_chroot(&cmd_list, true, c)

	for i := range c.Pacman.AUR.Packages {
		switch c.Pacman.AUR.Packages[i] {
		case "firefox-pwa", "firefox-pwa-bin":
			log.Println("Configuring Firefox for PWAs")
			cmd_list = []string{`firefoxpwa`, `runtime`, `install`}
			z.Arch_chroot(&cmd_list, true, c)
		case "openrazer-daemon":
			log.Println("Enabling OpenRazer Daemon")
			cmd_list = []string{`systemctl`, `--user`, `enable`, `openrazer-daemon.service`}
			z.Arch_chroot(&cmd_list, true, c)
		case "openrazer-driver-dkms":
			log.Println("Configuring Razer Drivers...")
			cmd_list = []string{`sudo`, `gpasswd`, `-a`, c.User.Username, `plugdev`}
			z.Arch_chroot(&cmd_list, true, c)
		case "openrazer-meta":
			log.Println("Configuring Razer Drivers...")
			cmd_list = []string{`sudo`, `gpasswd`, `-a`, c.User.Username, `plugdev`}
			z.Arch_chroot(&cmd_list, true, c)

			log.Println("Enabling OpenRazer Daemon")
			cmd_list = []string{`systemctl`, `--user`, `enable`, `openrazer-daemon.service`}
			z.Arch_chroot(&cmd_list, true, c)
		default:
		}
	}
}

func UserPWAs(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Installing Progressive Web Apps
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{`firefoxpwa`, `site`, `install`}

	for i := range c.PWA.Sites {
		cmd_list = append(cmd_list,
			fmt.Sprintf(`%s/%s`, c.PWA.Sites[i].StartUrl, c.PWA.Sites[i].Manifest),
			`--start-url`,
			c.PWA.Sites[i].StartUrl,
			`--name`,
			c.PWA.Sites[i].Name,
			`--description`,
			c.PWA.Sites[i].Description,
		)
	}

	z.Arch_chroot(&cmd_list, true, c)
}

func UserShell(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Configuring User Shell
	-------------------------------------------------------------------------
	`)

	// To-do: Add Shell Option during `Fresh Start`
	// Hard-coding ZSH shell
	c.User.Shell = "zsh"

	cmd_list := []string{`sudo`, `pacman`, `-Syyu`, `--needed`, `--noconfirm`}

	switch c.User.Shell {
	case "zsh":
		cmd_list = append(cmd_list,
			`zsh`,
			`grml-zsh-config`,
		)
		z.Arch_chroot(&cmd_list, true, c)

		cmd_list = []string{`sudo`, `chsh`, `--shell`, `/bin/zsh`, c.User.Username}
		z.Arch_chroot(&cmd_list, false, c)

		file_list := []string{
			`.bashrc`,
			`.bash_logout`,
			`.bash_profile`,
		}
		for file := range file_list {
			os.Remove(fmt.Sprintf(`/mnt/home/%s/%s`, c.User.Username, file_list[file]))
		}

	default:
		log.Println("User is configured with default shell")
	}
}
