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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	z "github.com/ShobuPrime/ShobuArch/pkg/shell"
	u "github.com/ShobuPrime/ShobuArch/pkg/util"
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
			log.Printf("Changing directory to %q\n", config_dir)
			if err := os.Chdir(config_dir); err != nil {
				log.Fatalln(err)
			}

			log.Println("Making Keepass directory")
			if err := os.MkdirAll("keepassxc", 0755); err != nil {
				log.Fatalln(err)
			}

			keepass_config := filepath.Join(config_dir, "keepassxc", "keepassxc.ini")

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
			f, err := os.OpenFile(keepass_config, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
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
			log.Printf("Changing directory to %q\n", user_home)
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

	cmd_list := []string{c.Pacman.AUR.Helper, `-Syyu`, `--needed`, `--noconfirm`}
	cmd_list = append(cmd_list, c.Pacman.AUR.Packages...)
	z.Arch_chroot(&cmd_list, true, c)

	for i := range c.Pacman.AUR.Packages {
		switch c.Pacman.AUR.Packages[i] {
		case "bcompare":
			log.Println("Service menus for Beyond Compare 4")
			cmd_list = []string{c.Pacman.AUR.Helper, `-Syyu`, `--needed`, `--noconfirm`}
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

	cmd_list := []string{}

	pwa_struct := &u.FIREFOX_PWA{}

	log.Println("Configuring separate profiles for each PWA...")
	create_profile := []string{`firefoxpwa`, `profile`, `create`, `--name`}

	for i := range c.PWA.Sites {
		cmd_list = append(create_profile, c.PWA.Sites[i].Name)
		z.Arch_chroot(&cmd_list, true, c)
	}

	log.Println("Reading PWA Config...")
	pwa_path := filepath.Join("/", "mnt", "home", c.User.Username, ".local", "share", "firefoxpwa")
	pwa_file := `config.json`
	firefoxpwa_config := u.ReadFile(&pwa_path, &pwa_file)

	_ = json.Unmarshal([]byte(strings.Join(*firefoxpwa_config, "\n")), pwa_struct)
	log.Println(u.PrettyJson(pwa_struct))

	log.Println("Installing PWAs into their respective profiles...")
	for _, profile := range pwa_struct.Profiles {
		cmd_list = []string{`firefoxpwa`, `site`, `install`}
		for i := range c.PWA.Sites {
			if profile.Name == c.PWA.Sites[i].Name {
				cmd_list = append(cmd_list,
					fmt.Sprintf(`%s/%s`, c.PWA.Sites[i].StartUrl, c.PWA.Sites[i].Manifest),
					`--start-url`,
					c.PWA.Sites[i].StartUrl,
					`--name`,
					c.PWA.Sites[i].Name,
					`--description`,
					c.PWA.Sites[i].Description,
					`--profile`,
					profile.Ulid,
				)
			}
		}
		z.Arch_chroot(&cmd_list, true, c)
	}
}

func UserVariables(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Configuring User Variables
	-------------------------------------------------------------------------
	`)

	log.Println("Configuring Environmental Variables")
	config_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config")
	environment_dir := filepath.Join(config_dir, "environment.d")

	log.Println("Making environment.d directory")
	if err := os.MkdirAll(environment_dir, 0755); err != nil {
		log.Fatalln(err)
	}

	environment_config := "environment.conf"

	config_settings := []string{
		`MOZ_ENABLE_WAYLAND=1`,
		`MOZ_DBUS_REMOTE=1`,
		`AMD_VULKAN_ICD=RADV`, // https://wiki.archlinux.org/title/Vulkan#Selecting_via_environment_variable
	}

	log.Println(`Creating "environment.conf"`)
	u.WriteFile(&environment_dir, &environment_config, &config_settings, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	log.Println("Configuring Electron")

	electron_config := "electron-flags.conf"

	config_settings = []string{
		`--enable-features=WaylandWindowDecorations`,
		`--ozone-platform-hint=auto`,
	}

	log.Println(`Creating "electron-flags.conf"`)
	u.WriteFile(&environment_dir, &electron_config, &config_settings, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	electron_config = "electron19-flags.conf"

	config_settings = []string{
		`--enable-features=UseOzonePlatform`,
		`--ozone-platform=wayland`,
	}
	log.Println(`Creating "electron19-flags.conf"`)
	u.WriteFile(&environment_dir, &electron_config, &config_settings, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "code":
			log.Println("Configuring Visual Studio Code")

			electron_config := "code-flags.conf"

			config_settings = []string{
				`--enable-features=UseOzonePlatform`,
				`--ozone-platform=wayland`,
			}
			log.Println(`Creating "code-flags.conf"`)
			u.WriteFile(&environment_dir, &electron_config, &config_settings, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		}
	}

	// To-Do:
	// For some applications, like bcompare, we need to make sure to export to PATH if not done already
	// ex) export PATH = /home/<username>/.cache/<aur_helper>/clone/bcompare/src/install/bin/bcompare:$PATH

	log.Println(`Enforcing permissions on home directory...`)
	//os.Chown(filepath.Base(keepass_config), 1000, 1000)
	cmd := []string{
		`chown`,
		`-R`,
		fmt.Sprintf(`%s:%s`, c.User.Username, c.User.Username),
		fmt.Sprintf(`/home/%s/`, c.User.Username),
	}
	z.Arch_chroot(&cmd, false, c)

	log.Println(`Adding .zshrc aliases...`)
	zshrc_dir := filepath.Join("/", "mnt", "home", c.User.Username)
	zshrc_file := ".zshrc"

	zshrc_contents := []string{
		fmt.Sprintf(`alias ARCH_UP='%s -Syyu --needed --noconfirm; flatpak update --assumeyes'`, c.Pacman.AUR.Helper),
		fmt.Sprintf(`alias DELETE_ORPHAN_PACKAGES='%s -Qtdq | %s -Rns -; flatpak uninstall --assumeyes --unused'`, c.Pacman.AUR.Helper, c.Pacman.AUR.Helper),
		fmt.Sprintf(`alias LIST_ALL_PACKAGES='%s -Qq; flatpak list --columns=application | tail -n +1'`, c.Pacman.AUR.Helper), // Ignore first line of output for Flatpak. Alternatives: awk '{if(NR>1)print}', sed -n '1!p'
		fmt.Sprintf(`alias LIST_BROKEN_PACKAGES="%s -Qk | grep -v ' 0 missing files'"`, c.Pacman.AUR.Helper),
		fmt.Sprintf(`#alias REPAIR_ALL_PACKAGES='for package in $(%s -Qq); do %s -S "$package" --noconfirm; done'`, c.Pacman.AUR.Helper, c.Pacman.AUR.Helper),
		fmt.Sprintf(`alias REPAIR_ALL_PACKAGES='%s -Qq | %s -S -; flatpak repair'`, c.Pacman.AUR.Helper, c.Pacman.AUR.Helper),
		fmt.Sprintf(`alias REPAIR_BROKEN_PACKAGES="%s -Qk | grep -v ' 0 missing files' | cut -d: -f1 | %s -S -"`, c.Pacman.AUR.Helper, c.Pacman.AUR.Helper),
	}
	u.WriteFile(&zshrc_dir, &zshrc_file, &zshrc_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
	u.ReadFile(&zshrc_dir, &zshrc_file)
}

func UserAutostart(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                Configuring User's AutoStart Programs
	-------------------------------------------------------------------------
	`)

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "keepassxc":
			log.Println(`Creating autostart for KeePassXC`)
			autostart_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "autostart")
			autostart_file := "org.keepassxc.KeePassXC.desktop"
			autostart_contents := []string{
				`[Desktop Entry]`,
				`Name=KeePassXC`,
				`GenericName=Password Manager`,
				`Exec=/usr/bin/keepassxc`,
				`TryExec=/usr/bin/keepassxc`,
				`Icon=keepassxc`,
				`StartupWMClass=keepassxc`,
				`StartupNotify=true`,
				`Terminal=false`,
				`Type=Application`,
				`Version=1.0`,
				`Categories=Utility;Security;Qt;`,
				`MimeType=application/x-keepass2;`,
				`X-GNOME-Autostart-enabled=true`,
				`X-GNOME-Autostart-Delay=2`,
				`X-KDE-autostart-after=panel`,
				`X-LXQt-Need-Tray=true`,
			}
			u.WriteFile(&autostart_dir, &autostart_file, &autostart_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
		}
	}

	for i := range c.Pacman.AUR.Packages {
		switch c.Pacman.AUR.Packages[i] {
		case "enpass", "enpass-bin":
			log.Println(`Creating autostart for Enpass`)
			autostart_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "autostart")
			autostart_file := "Enpass.desktop"
			autostart_contents := []string{
				`[Desktop Entry]`,
				`Type=Application`,
				`Name=Enpass`,
				`Exec= /opt/enpass/Enpass -minimize`,
				`Icon=enpass.png`,
				`Comment=The best password manager`,
				`X-GNOME-Autostart-Delay=12`,
				`X-GNOME-Autostart-enabled=true`,
			}
			u.WriteFile(&autostart_dir, &autostart_file, &autostart_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
		}
	}
}

func UserDotFiles(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                    Configuring User DotFiles
	-------------------------------------------------------------------------
	`)

	log.Println("Preparing environment for automatic systemd-nspawn scripts...")

	log.Println(`Creating AutoLogin for systemd-nspawn container...`)
	autologin_dir := filepath.Join("/", "mnt", "etc", "systemd", "system", "console-getty.service.d")
	err := os.MkdirAll(autologin_dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	autologin_file := "autologin.conf"
	autologin_contents := []string{
		`[Service]`,
		`ExecStart=`,
		fmt.Sprintf(`ExecStart=-/sbin/agetty -o '-p -f -- \\u' --noclear --keep-baud --autologin %s - 115200,38400,9600 $TERM`, c.User.Username),
	}
	u.WriteFile(&autologin_dir, &autologin_file, &autologin_contents, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	log.Println("Compiling VSCode extension commands...")
	cmd_list := []string{`sleep 3`}
	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "code":
			// Install extensions
			log.Println(`Installing VSCode Extensions`)
			code_extensions := []string{`shardulm94.trailing-spaces`,}
			for i := range c.Pacman.Packages {
				switch c.Pacman.Packages[i] {
				case "docker":
					code_extensions = append(code_extensions, `ms-azuretools.vscode-docker`)
				case "go":
					code_extensions = append(code_extensions, `golang.go`)
				case "python":
					code_extensions = append(code_extensions, `ms-python.python`)
				}
			}

			for i := range code_extensions {
				cmd_list = append(cmd_list, fmt.Sprintf(`code --install-extension %s`, code_extensions[i]),
				)
			}

			cmd_list := append(cmd_list, `code --list-extensions`)

			log.Println("Appending systemd-nspawn 'Get out of Jail for free' command...")
			cmd_list = append(cmd_list, `sudo poweroff`)

			log.Println("Ensuring VSCode extensions will automatically execute after mounting systemd-nspawn container...")
			systemd_autorun_dir := filepath.Join("/", "mnt", "etc", "profile.d")
			code_script := "install_code_extensions.sh"

			log.Println(`Creating VSCode script for systemd-nspawn container...`)
			u.WriteFile(&systemd_autorun_dir, &code_script, &cmd_list, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

			log.Println("Making executable...")
			cmd := fmt.Sprintf(`chmod +x %v`, filepath.Join(systemd_autorun_dir, code_script))
			z.Shell(&cmd)

			log.Println("Installing Code extensions via systemd-nspawn...")
			z.Systemd_nspawn(&[]string{}, true, c)

			log.Println("Cleaning up cruft...")
			log.Println("Deleting: ", autologin_dir)
			err = os.RemoveAll(autologin_dir)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("Deleting: ", filepath.Join(systemd_autorun_dir, code_script))
			err = os.Remove(filepath.Join(systemd_autorun_dir, code_script))
			if err != nil {
				log.Fatalln(err)
			}

			log.Println(`Creating config for VSCode...`)
			code_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "Code - OSS", "User")
			code_file := "settings.json"
			code_contents := []string{
				`{`,
				`	"files.autoSave": "afterDelay",`,
				`	"go.toolsManagement.autoUpdate": true,`,
				`	"git.autofetch": true,`,
				`	"launch": {`,
				`		"configurations": [`,
				`			{`,
				`				"name": "Launch Go",`,
				`				"type": "go",`,
				`				"request": "launch",`,
				`				"mode": "auto",`,
				`				"program": "${fileDirname}",`,
				`				"console": "integratedTerminal"`,
				`			},`,
				`			{`,
				`				"name": "Launch Python",`,
				`				"type": "python",`,
				`				"request": "launch",`,
				`				"program": "${file}",`,
				`				"console": "integratedTerminal"`,
				`			}`,
				`		],`,
				`		"compounds": []`,
				`	},`,
				`	"debug.allowBreakpointsEverywhere": true,`,
				`	"git.confirmSync": false,`,
				`	"git.enableSmartCommit": true,`,
				`	"diffEditor.ignoreTrimWhitespace": false,`,
				`	"go.formatTool": "gofmt",`,
				`	"go.formatFlags": [`,
				`		"[\"-s\", \"-w\"]"`,
				`	],`,
				`	"python.formatting.provider": "black",`,
				`	"editor.formatOnSave": true`,
				`}`,
			}
			u.WriteFile(&code_dir, &code_file, &code_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite

			code_dir = filepath.Join("/", "mnt", "home", c.User.Username, ".vscode-oss")
			code_file = "argv.json"

			// This configuration file allows you to pass permanent command line arguments to VS Code.
			// https://github.com/microsoft/vscode-python/issues/20247#issuecomment-1350342224`
			code_contents = []string{
				`{`,
				`	"enable-proposed-api": ["ms-python.python"]`,
				`}`,
			}
			u.WriteFile(&code_dir, &code_file, &code_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
		}
	}
}

func UserShell(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Configuring User Shell
	-------------------------------------------------------------------------
	`)

	log.Println(`Enforcing permissions`)
	//os.Chown(filepath.Base(keepass_config), 1000, 1000)
	cmd := []string{
		`chown`,
		`-R`,
		fmt.Sprintf(`%s:%s`, c.User.Username, c.User.Username),
		fmt.Sprintf(`/home/%s/`, c.User.Username),
	}
	z.Arch_chroot(&cmd, false, c)

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
