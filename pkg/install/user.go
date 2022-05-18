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
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

func UserKeyring(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                    Initializing System Keyring
	-------------------------------------------------------------------------
	`)

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "keepassxc":
			log.Println("Configuring KeePassXC as System Keyring")

			pwd, err := os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}

			config_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config")
			log.Printf("Changing directory to %q", config_dir)
			if err := os.Chdir(config_dir); err != nil {
				log.Fatalln(err)
			}

			log.Printf("Making Keepass directory")
			if err := os.MkdirAll("keepass", 0755); err != nil {
				log.Fatalln(err)
			}

			keepass_config := filepath.Join(config_dir, "keepass", "keepassxc.ini")

			config_settings := []string{
				`[General]`,
				`ConfigVersion=2`,
				`MinimizeAfterUnlock=true`,
				``,
				`[Browser]`,
				`CustomProxyLocation=`,
				``,
				`[FdoSecrets]`,
				`Enabled=true`,
				``,
				`[GUI]`,
				`AdvancedSettings=true`,
				`MinimizeOnClose=true`,
				`MinimizeOnStartup=true`,
				`ShowExpiredEntriesOnDatabaseUnlockOffsetDays=0`,
				`ShowTrayIcon=true`,
				`TrayIconAppearance=colorful`,
				``,
				`[PasswordGenerator]`,
				`AdditionalChars=`,
				`ExcludedChars=`,
				``,
				`[SSHAgent]`,
				`Enabled=true`,
			}
			log.Println(`Creating "keepassxc.ini"`)
			f, err := os.OpenFile(keepass_config, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(`Saving settings`)
			if _, err := f.Write([]byte(strings.Join(config_settings, "\n"))); err != nil {
				log.Fatalln(err)
			}
			log.Println(`Closing file`)
			if err := f.Close(); err != nil {
				log.Fatalln(err)
			}

			// keepassxc-cli db-create <name>.kbdx --set-key-file <filepath>
			// keepasscx-cli db-info <name>.kbdx --key-file <filepath> --no-password
			// Default password `shobuarch`
			database := `A9mimmf7S7UAAAQAAhAAAAAxwfLmv3FDUL5YBSFq/Fr/AwQAAAABAAAABCAAAAA6gIwPGKUKHY7tiznPSf+tOKAaofJx0/dXwJwzhUDPiAcQAAAAlUosWDY0KACeIQAEQBY/9QuLAAAAAAFCBQAAACRVVUlEEAAAAO9jbd+MKURLkfeppAPjCgwFAQAAAEkIAAAACgAAAAAAAAAFAQAAAE0IAAAAAAAABAAAAAAEAQAAAFAEAAAAEAAAAEIBAAAAUyAAAAAjkwIV6OEByQOujHV5OOkm8zx2RArSVpyMCAuCgFsmbQQBAAAAVgQAAAATAAAAAAAEAAAADQoNCmLMN98TJfsNuFrKZSspbQtP/nBNuEQ4I6CO83eMhDlss9yWOR4sedvhPxYpjwla3NtoTlvTeY3fhqnwDk4COnkQgLyATEl23l5cLDmPJcfIwJBvNUljPmII41wK0a+HXAAEAACZjhMpkaJb6J7OLSLoAjsDREHabM3+HlQ3uKnpOKsRVbAgaup63fR97nEDfVdSdGBONSpNdKSpfaZBsg3qJrVfJZ308sQ+zSVVFkOo78NBjCLhNMerGJ2ALzFnJHli0ESOsCDRO4SynI5i1aEJ/gKCjEk/DHyDZYyXehIBhcsU5a61p8B2xOgpvOZ/m6ilWJKDBiN18R5uxZAABxlaIMaU1/f0UwAsoe60GHZ8DAN1hYHE7CTCyTjn/vMlKRq8pyOJAJ9f0MoZ0eq3TrYmYdGR+G2zVVWhMBF9uYYcEcBEaDx6MFf+/J/2vrRkmZM1GjSV+gn+yFByqrY83fjBYtHuQSUO65/KS1cQPSOCjTQ1vuGlhPSEgDS8Wi3QpwNWgvYDWmwRc0fUGeL+of10hnAOq8f5aGGuvGr3nI4uzDufE3QD1338cSQ3krEY/85xNE7lCViT85eC+GZDTArWAVoJfKIsdEL5hgUey7rMOLhHpA+TihbTNvI3XYi43/1UIWeqQZxQYBDhFplDlp5qSTCHi55jlyCxgRXXeZyRhrn5DhLQCteFhVjApRKgyzZU6uk0KsrSY1oNSFDprO1OrKOiJOEDTHUJ9083oInYMZXAqVdGUc0m5PxnNxYDFyXS6OXp0PUb/MSezVLhf1/xd4jvNeOIfVVX8TvNdaKRtJ72KJxSH1uH23ETSmOeN8vxKKJKJCEq62QZbRTdnI4xmuoO1foZS312lWhyLfU2dU5CzQp+nHBXf3sS32zXFD5V+5TOdOSWGZNM45V4NWiFA8PUzFWsexknJ5SpFp+jap9Gi0fcq310dw6teDLzE41IzEkg4X27H2SqZA4d/J+9AGSa6r0lRWUBeF8PXFR8iFhIjG95eUYxUa4M/UUMTb6z/1qjEPqIUjLPybK98vWfAEICBA5YT+1WXsOzA0zZ9ex+wTRSIp0t8xs4XMlF3ONKVgyG5S9uOxdaBGiRXg+TDk8kcVzT5GOoAT7BdIE/hgedGbNB8QNrk7B62jCZUOSgd4lkHKh1fOLk3qGTVPg2g9OiLrJBZsiDHPodi613dbBZTpJOQjdAMUh4D5M3W8rLmN/ydXWrd9euYx/KH0b6UOuiDRZoe0/tygeSgGn0QPFQadrOVDMsy4mTwYgqeBvUGQh10YBfdmgvoF6ATvt/WnE9g6RLE3Eip7iUQL4Osb7brzouDgroXw8pD++4hI0e3D+xpOmRzpAxo0C5/T45WcP9X2US1264TTwkh8+hI8p7+Kjr5TM5n+BZvfHd2rVF2cul8wc10HQqr8S22vB223AUilKccnVVHs0o64i7N4vZNkizxsEPVwBJyr0NdhGOqjUOUR67Up8d2pv2BFhGxxFKUWbm8MgcyDtUl64Kc+APpOPSDAZ3q012Yv26qUO7N6UAAAAA`
			dec, err := base64.StdEncoding.DecodeString(database)
			if err != nil {
				log.Fatalln(err)
			}

			user_home := filepath.Join("/", "mnt", "home", c.User.Username)
			log.Printf("Changing directory to %q", user_home)
			if err := os.Chdir(user_home); err != nil {
				log.Fatalln(err)
			}

			log.Println(`Creating "system-keyring" database`)
			f, err = os.Create("system-keyring.kdbx")
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()

			log.Println(`Writing data to database`)
			if _, err := f.Write(dec); err != nil {
				log.Fatalln(err)
			}

			log.Println(`Waiting for data to synchronize`)
			if err := f.Sync(); err != nil {
				log.Fatalln(err)
			}

			log.Println(`Enforcing permissions`)
			//os.Chown(filepath.Base(keepass_config), 1000, 1000)
			cmd := []string{
				`chown`,
				`-R`,
				fmt.Sprintf(`%s:%s`, c.User.Username, c.User.Username),
				fmt.Sprintf(`/home/%s/`, c.User.Username),
			}
			z.Arch_chroot(&cmd, false, c)

			log.Println("Done!")
			// Return to original directory
			if err := os.Chdir(pwd); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func UserPackages(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Installing AUR Packages
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{`yay`, `-Syyu`, `--needed`, `--noconfirm`}
	for i := range c.Pacman.AUR.Packages {
		cmd_list = append(cmd_list, c.Pacman.AUR.Packages[i])
	}
	z.Arch_chroot(&cmd_list, true, c)

	for i := range c.Pacman.AUR.Packages {
		switch c.Pacman.AUR.Packages[i] {
		case "bcompare":
			log.Println("Service menus for Beyond Compare 4")
			cmd_list = []string{`yay`, `-Syyu`, `--needed`, `--noconfirm`}
			switch c.Desktop.Environment {
			case "cinnamon":
				cmd_list = append(cmd_list, `bcompare-cinnamon`)
				z.Arch_chroot(&cmd_list, true, c)
			case "gnome":
				cmd_list = append(cmd_list, `bcompare-nautilus`)
				z.Arch_chroot(&cmd_list, true, c)
			case "mate":
				cmd_list = append(cmd_list, `bcompare-mate`)
				z.Arch_chroot(&cmd_list, true, c)
			case "kde":
				cmd_list = append(cmd_list, `bcompare-kde5`)
				z.Arch_chroot(&cmd_list, true, c)
			case "xfce":
				cmd_list = append(cmd_list, `bcompare-thunar`)
				z.Arch_chroot(&cmd_list, true, c)
			}
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
