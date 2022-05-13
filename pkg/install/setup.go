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
	"runtime"
	"strconv"
	"strings"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	z "github.com/ShobuPrime/ShobuArch/pkg/shell"
	u "github.com/ShobuPrime/ShobuArch/pkg/util"
)

func SetupLogo() {

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
                            PROGRESS: Setup
                          SCRIPTHOME: ShobuArch
	-------------------------------------------------------------------------
	`)
}

func SetupHostname(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                            Setting Hostname
	-------------------------------------------------------------------------
	`)

	cmd := []string{
		`awk`,
		fmt.Sprintf(`BEGIN{ printf "%s\n" >> "/etc/hostname" }`, c.Hostname),
	}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{
		`awk`,
		`BEGIN{ printf "127.0.0.1    localhost\n" >> "/etc/hosts" }`,
	}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{
		`awk`,
		`BEGIN{ printf "::1          localhost\n" >> "/etc/hosts" }`,
	}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{
		`awk`,
		fmt.Sprintf(`BEGIN{ printf "127.0.1.1    %s.local ShobuLANlord\n" >> "/etc/hosts" }`, c.Hostname),
	}

	z.Arch_chroot(&cmd, false, c)
}

func SetupNetwork(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                            Network Setup
	-------------------------------------------------------------------------
	`)

	cmd := []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`, `networkmanager`, `dhclient`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`systemctl`, `enable`, `--now`, `NetworkManager`}
	z.Arch_chroot(&cmd, false, c)

	log.Println("Add parallel downloading")
	cmd = []string{`sed`, `-i`, `s/^#ParallelDownloads/ParallelDownloads/g`, `/etc/pacman.conf`}
	z.Arch_chroot(&cmd, false, c)

	// This specific command is giving errors on Arch-Chroot. Using host shell with hard-coded mount point
	log.Println("Enable multi-lib libraries")
	bash := `sed -i "/\[multilib\]/,/Include/"'s/^#//' /mnt/etc/pacman.conf`
	z.Shell(&bash)

	cmd = []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`}
	z.Arch_chroot(&cmd, false, c)
}

func SetupMirrors(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
			Setting up mirrors for optimal download 
	-------------------------------------------------------------------------
	`)

	cmd := []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`, `pacman-contrib`, `curl`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`, `reflector`, `rsync`, `grub`, `arch-install-scripts`, `git`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`cp`, `/etc/pacman.d/mirrorlist`, `/etc/pacman.d.mirrorlist.bak`}
	z.Arch_chroot(&cmd, false, c)
}

func SetupResources(c *conf.Config) {
	log.Printf(`
	-------------------------------------------------------------------------
				You have "%v" cores.
			Changing the makeflags and compression
			   settings for %v cores.
	-------------------------------------------------------------------------
	`, runtime.NumCPU(), runtime.NumCPU())

	log.Println("Checking system total memory...")
	cmd := `(cat /proc/meminfo | grep -i 'memtotal' | grep -o '[[:digit:]]*')`
	out := z.Shell(&cmd)
	out = strings.TrimSpace(out)
	TOTAL_MEM, _ := strconv.ParseInt(out, 10, 64)
	log.Println(out)
	log.Printf("System total memory is: %d\n", TOTAL_MEM)

	if TOTAL_MEM > 8000000 {
		// Running from host instead of arch-chroot for now
		// To-do: Fix this
		cmd = `sed -n "s/#MAKEFLAGS=\"-j2\"/MAKEFLAGS=\"-j%s\"/g" /mnt/etc/makepkg.conf`
		z.Shell(&cmd)

		cmd = `sed -n "s/COMPRESSXZ=(xz -c -z -)/COMPRESSXZ=(xz -c -T %s -z -)/g" /mnt/etc/makepkg.conf`
		z.Shell(&cmd)
	}
}

func SetupLanguage(c *conf.Config) {
	log.Printf(`
	-------------------------------------------------------------------------
			Setup Language to %s and setting locale
	-------------------------------------------------------------------------
	`, c.User.Language.Locale)

	// en_US.UTF-8
	cmd := []string{
		`sed`,
		`-i`,
		fmt.Sprintf(
			`s/#%s %s/%s %s/g`,
			c.User.Language.Locale,
			c.User.Language.CharSet,
			c.User.Language.Locale,
			c.User.Language.CharSet,
		),
		`/etc/locale.gen`,
	}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`locale-gen`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`timedatectl`, `set-timezone`, c.Timezone}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`timedatectl`, `set-ntp`, `1`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`touch`, `/etc/locale.conf`}
	z.Arch_chroot(&cmd, false, c)

	locale_list := []string{
		fmt.Sprintf(`LANG=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LANGUAGE=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_ADDRESS=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_CTYPE=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_IDENTIFICATION=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_MEASUREMENT=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_MESSAGES=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_MONETARY=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_NAME=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_NUMERIC=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_PAPER=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_TELEPHONE=%s`, c.User.Language.Locale),
		fmt.Sprintf(`LC_TIME=%s`, c.User.Language.Locale),
	}

	for locale := range locale_list {
		cmd = []string{`awk`,
			fmt.Sprintf(`BEGIN{ printf "%s\n" >> "/etc/locale.conf" }`, locale_list[locale]),
		}
		z.Arch_chroot(&cmd, false, c)
	}

	cmd = []string{`ln`, `-s`, fmt.Sprintf(`/usr/share/zoneinfo/%s`, c.Timezone), `/etc/localtime`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`localectl`, `set-keymap`, c.User.Language.Keyboard}
	z.Arch_chroot(&cmd, false, c)
}

func SetupBaseSystem(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                            Installing Base System
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`}
	for i := range c.Pacman.Packages {
		cmd_list = append(cmd_list, c.Pacman.Packages[i])
	}
	z.Arch_chroot(&cmd_list, false, c)

	switch u.GetHostStatus().Chassis {
	case "vm":
		log.Println("Virtual machine detected! Installing open-vm-tools")
		cmd_list = []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`, `open-vm-tools`}
		z.Arch_chroot(&cmd_list, false, c)

		cmd_list = []string{`systemctl`, `enable`, `vmtoolsd.service`, `vmware-vmblock-fuse.service`}
		z.Arch_chroot(&cmd_list, false, c)
	}

	log.Println("Add sudo no password rights")
	cmd := []string{`sed`, `-i`, `s/^# %wheel ALL=(ALL) NOPASSWD: ALL/%wheel ALL=(ALL) NOPASSWD: ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`sed`, `-i`, `s/^# %wheel ALL=(ALL:ALL) NOPASSWD: ALL/%wheel ALL=(ALL:ALL) NOPASSWD: ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)
}

func SetupServices(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                            Enabling Services
	-------------------------------------------------------------------------
	`)

	cmd_list := []string{`systemctl`, `enable`}

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "cockpit":
			cmd_list = append(cmd_list,
				`cockpit.socket`,
			)
		case "docker":
			cmd_list = append(cmd_list,
				`docker.service`,
			)
		case "firewalld":
			cmd_list = append(cmd_list,
				`firewalld`,
			)
		case "virt-manager":
			cmd_list = append(cmd_list,
				`libvirtd.service`,
			)
		case "reflector":
			cmd_list = append(cmd_list,
				`reflector.timer`,
			)
		case "sddm":
			cmd_list = append(cmd_list,
				`sddm`,
			)
		case "openssh":
			cmd_list = append(cmd_list,
				`sshd`,
			)
		}
	}

	z.Arch_chroot(&cmd_list, false, c)
}

func SetupCustomRepos(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
			Setting up custom repositories 
	-------------------------------------------------------------------------
	`)

	cmd := []string{}

	switch c.Storage.Filesystem {
	case "zfs":
		log.Println("Adding custom repo to install ArchZFS")
		cmd = []string{`wget`, `https://archzfs.com/archzfs.gpg`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "[archzfs]\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`pacman-key`, `-a`, `archzfs.gpg`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`pacman-key`, `-r`, `DDF7DB817396A49B2A2723F7403BD972F75D9D76`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`pacman-key`, `--lsign-key`, `DDF7DB817396A49B2A2723F7403BD972F75D9D76`}
		z.Arch_chroot(&cmd, false, c)

		// # Check the fingerprint and verify it matches the one on the archzfs page
		cmd = []string{`pacman-key`, `-f`, `DDF7DB817396A49B2A2723F7403BD972F75D9D76`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "# Origin Server - France\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "Server = http://archzfs.com/$repo/x86_64\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "# Mirror - Germany\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "Server = http://mirror.sum7.eu/archlinux/archzfs/$repo/x86_64\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "# Mirror - Germany\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "Server = https://mirror.biocrafting.net/archlinux/archzfs/$repo/x86_64\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "# Mirror - India\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "Server = https://mirror.in.themindsmaze.com/archzfs/$repo/x86_64\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "# ArchZFS - US Mirror\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`awk`, `BEGIN{ printf "Server = https://zxcvfdsa.com/archzfs/$repo/$arch\n" >> "/etc/pacman.conf" }`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`, `zfs-dkms`, `zfs-utils`}
		z.Arch_chroot(&cmd, false, c)

		cmd_list := []string{`systemctl`, `enable`}
		service_names := []string{
			`zfs-import-cache`,
			`zfs-import-scan`,
			`zfs-mount`,
			`zfs-share`,
			`zfs-zed`,
			`zfs.target`,
		}

		for i := range service_names {
			cmd_list = append(cmd_list, service_names[i])
		}
		z.Arch_chroot(&cmd_list, false, c)

		err := os.Remove("/mnt/archzfs.gpg")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func SetupMicrocode(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                            Installing Microcode
	-------------------------------------------------------------------------
	`)

	cpu := u.ListCPU()

	cmd := []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`}

	for i := range cpu.Processor {
		switch cpu.Processor[i].Data {
		case "GenuineIntel":
			log.Printf("Detected CPU is: %q", cpu.Processor[i].Data)
			cmd = append(cmd, `intel-ucode`)
			z.Arch_chroot(&cmd, false, c)
		case "AuthenticAMD":
			log.Printf("Detected CPU is: %q", cpu.Processor[i].Data)
			cmd = append(cmd, `amd-ucode`)
			z.Arch_chroot(&cmd, false, c)
		}
	}
}

func SetupGraphics(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                        Installing Graphics Drivers
	-------------------------------------------------------------------------
	`)

	p := u.ListPCI()

	cmd := []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`}

	for i := range p.PCIDevices {
		if p.PCIDevices[i].Class == "VGA compatible controller" {

			log.Printf("Detected Graphics Device: %s\n", p.PCIDevices[i].Device)
			log.Printf("Graphics Vendor:%s\n", p.PCIDevices[i].Vendor)

			switch p.PCIDevices[i].Vendor {
			case "NVIDIA Corporation":
				switch c.Kernel {
				case "Standard", "linux":
					cmd = append(cmd, `nvidia`)
				case "Longterm", "linux-hardened":
					cmd = append(cmd, `nvidia-lts`)
				default:
					cmd = append(cmd, `nvidia-dkms`)
				}
				cmd = append(cmd,
					`nvidia-xconfig`,
				)
			case "Advanced Micro Devices, Inc. [AMD]", "Advanced Micro Devices, Inc. [AMD/ATI]":
				cmd = append(cmd,
					`lib32-libva-mesa-driver`,
					`lib32-mesa`,
					`lib32-vulkan-radeon`,
					`libva-mesa-driver`,
					`mesa`,
					`vulkan-radeon`,
					`xf86-video-amdgpu`,
				)
			case "Intel Corporation":
				cmd = append(cmd,
					`lib32-mesa`,
					`lib32-vulkan-intel`,
					`libva-intel-driver`,
					`libva-utils`,
					`mesa`,
					`vulkan-intel`,
					`xf86-video-intel`,
				)
			}
		}
	}

	z.Arch_chroot(&cmd, false, c)
}

func SetupUser(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                        Adding User
	-------------------------------------------------------------------------
	`)

	// https://wiki.archlinux.org/title/Users_and_groups#Group_management
	group_list := []string{
		`adm`,     // Read access to protected logs
		`audio`,   // Direct access to sound hardware, for all sessions
		`floppy`,  // Floppy drives
		`log`,     // Access to /var/log
		`lp`,      // Parallel port devices (Printers)
		`network`, // NetworkManager
		`optical`, // Optical devices
		`power`,   // [Unused]
		`rfkill`,  // Control wireless devices power state
		`scanner`, // Access to Scanner hardware
		`storage`, // Removable drives such as USB hard drives
		`sys`,     // Administer printers in CUPS
		`users`,   // The primary group for users when user private groups are not used
		`video`,   // Access to video capture devices, 2D/3D hardware acceleration, framebuffer
		`wheel`,   // Become superuser/root by using su
	}

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "docker":
			group_list = append(group_list,
				`docker`,
			)
		case "virt-manager":
			group_list = append(group_list,
				`kvm`,
			)
		}
	}

	cmd := []string{
		`useradd`, `-mU`,
		`-s`, `/bin/bash`, // /bin/zsh
		`-G`, strings.Join(group_list, ","),
		`-d`, fmt.Sprintf(`/home/%s/`, c.User.Username), c.User.Username,
	}
	z.Arch_chroot(&cmd, false, c)

	log.Printf("Setting password for %s\n", c.User.Username)
	cmd = []string{
		fmt.Sprintf(`echo "%s:%s" | chpasswd --root /mnt`, c.User.Username, c.User.Password),
	}
	z.Shell(&cmd[0])

	log.Println("Initializing user directories...")
	user_directories := []string{
		`.cache`, `.config`, `.local/share/applications`, `.local/share/icons`,
		`Desktop`, `Documents`, `Downloads`, `Music`, `Pictures`, `Public`,
		`Templates`, `Videos`,
	}

	for i := range user_directories {
		log.Println(user_directories[i])
		err := os.MkdirAll(fmt.Sprintf(`/mnt/home/%s/%s`, c.User.Username, user_directories[i]), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd = []string{
		`chown`,
		`-R`,
		fmt.Sprintf(`%s:%s`, c.User.Username, c.User.Username),
		fmt.Sprintf(`/home/%s/`, c.User.Username),
	}
	z.Arch_chroot(&cmd, false, c)

	log.Println("Configuring user as sudoer...")
	cmd = []string{`sudo`, `sed`, `-i`, `s/# %wheel ALL=(ALL:ALL) ALL/%wheel ALL=(ALL:ALL) ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`sudo`, `sed`, `-i`, `s/# %wheel ALL=(ALL:ALL) NOPASSWD: ALL/%wheel ALL=(ALL:ALL) NOPASSWD: ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`sudo`, `sed`, `-i`, `s/# %sudo ALL=(ALL:ALL) ALL/%sudo ALL=(ALL:ALL) ALL/g`, `/etc/sudoers`}
	z.Arch_chroot(&cmd, false, c)
}

func SetupAUR(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                        Adding AUR Helper
	-------------------------------------------------------------------------
	`)
	cmd := []string{
		`git`, `clone`,
		fmt.Sprintf(`https://aur.archlinux.org/%s.git`, c.Pacman.AUR.Helper),
		fmt.Sprintf(`/opt/%s`, c.Pacman.AUR.Helper),
	}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{
		`chown`,
		`-R`,
		fmt.Sprintf(`%s:users`, c.User.Username),
		fmt.Sprintf(`/opt/%s/`, c.Pacman.AUR.Helper),
	}
	z.Arch_chroot(&cmd, false, c)

	cmd_list := []string{
		fmt.Sprintf(`echo "cd /opt/%s/" >> /mnt/install_aur.sh`, c.Pacman.AUR.Helper),
		fmt.Sprintf(`echo "su -c 'makepkg -sc --noconfirm' %s" >> /mnt/install_aur.sh`, c.User.Username),
		fmt.Sprintf(`echo "%s ALL=(ALL) NOPASSWD: ALL" >> /mnt/etc/sudoers.d/local_users`, c.User.Username),
	}
	for i := range cmd_list {
		z.Shell(&cmd_list[i])
	}

	log.Println("Making AUR Package")
	cmd = []string{`chmod`, `+x`, `/install_aur.sh`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`./install_aur.sh`}
	z.Arch_chroot(&cmd, false, c)

	aur_path := fmt.Sprintf("/mnt/opt/%s", c.Pacman.AUR.Helper)
	f, err := os.Open(aur_path)
	if err != nil {
		log.Fatalln(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	aur_pkg := ""
	for _, file := range files {
		log.Println(file.Name())
		if strings.HasSuffix(file.Name(), ".pkg.tar.zst") {
			aur_pkg = file.Name()
		}
	}

	cmd = []string{`sudo`, `pacman`, `-U`, `--needed`, `--noconfirm`, fmt.Sprintf("/opt/%s/%s", c.Pacman.AUR.Helper, aur_pkg)}
	z.Arch_chroot(&cmd, true, c)

	log.Println("Cleaning up AUR cruft...")
	err = os.RemoveAll(fmt.Sprintf("/opt/%s/", c.Pacman.AUR.Helper))
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Remove("/mnt/install_aur.sh")
	if err != nil {
		log.Fatalln(err)
	}
}

func SetupFlatpaks(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                        Installing Flatpaks
	-------------------------------------------------------------------------
	`)

	for i := range c.Flatpak.Packages {
		cmd_list := []string{`flatpak`, `install`, `--assumeyes`, `flathub`, c.Flatpak.Packages[i]}
		z.Arch_chroot(&cmd_list, false, c)

		switch c.Flatpak.Packages[i] {
		case "com.brave.Browser", "com.google.Chrome", "com.microsoft.Edge":
			log.Println("Chromium browser Flatpak detected!")
			log.Println("Adding permissions for Progressive Web Apps")
			cmd_list = []string{`flatpak`, `override`, c.Flatpak.Packages[i], fmt.Sprintf(`--filesystem=/home/%s/.local/share/applications`, c.User.Username)}
			z.Arch_chroot(&cmd_list, false, c)

			cmd_list = []string{`flatpak`, `override`, c.Flatpak.Packages[i], fmt.Sprintf(`--filesystem=/home/%s/.local/share/icons`, c.User.Username)}
			z.Arch_chroot(&cmd_list, false, c)
		}
	}
}

func SetupEFI(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                        Adding EFI Configs
	-------------------------------------------------------------------------
	`)

	// Assuming Grub by default for now.
	// To-do: Future work will consider other EFI loaders
	cmd := []string{`sed`, `-i`, `s/quiet/quiet video=1920x1080/g`, `/etc/default/grub`}
	z.Arch_chroot(&cmd, false, c)

	switch c.Storage.Filesystem {
	case "zfs":
		cmd = []string{`sed`, `-i`, `s/GRUB_CMDLINE_LINUX=""/GRUB_CMDLINE_LINUX="root=ZFS=zroot\/ROOT\/default"/g`, `/etc/default/grub`}
		z.Arch_chroot(&cmd, false, c)
	}

	cmd = []string{`grub-install`, `--target=x86_64-efi`, `--efi-directory=/boot`, `--bootloader-id=ArchLinux`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`grub-mkconfig`, `-o`, `/boot/grub/grub.cfg`}
	z.Arch_chroot(&cmd, false, c)

	switch c.Storage.Filesystem {
	case "luks":
		// HOOKS=(base udev autodetect modconf block keyboard encrypt filesystems)
		cmd = []string{`sed`, `-i`, `s/HOOKS=(base udev autodetect modconf block filesystems keyboard fsck)/HOOKS=(base udev autodetect modconf block keyboard encrypt filesystems shutdown)/g`, `/etc/mkinitcpio.conf`}
		z.Arch_chroot(&cmd, false, c)
	case "zfs":
		// HOOKS=(base udev autodetect modconf block keyboard zfs filesystems)
		cmd = []string{`sed`, `-i`, `s/HOOKS=(base udev autodetect modconf block filesystems keyboard fsck)/HOOKS=(base udev autodetect modconf block keyboard zfs filesystems shutdown)/g`, `/etc/mkinitcpio.conf`}
		z.Arch_chroot(&cmd, false, c)
	}

	cmd = []string{`mkinitcpio`, `-p`, c.Kernel}
	z.Arch_chroot(&cmd, false, c)
}
