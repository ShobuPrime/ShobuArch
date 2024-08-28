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
	"path/filepath"
	"regexp"
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

	cmd = []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`}
	z.Arch_chroot(&cmd, false, c)
}

func SetupMirrors(c *conf.Config) {
	iso := "US"

	log.Printf(`
		-------------------------------------------------------------------------
					Setting up %s mirrors for faster downloads
		-------------------------------------------------------------------------
	`, iso)

	// cmd := `(curl -4 ifconfig.co/country-iso)`
	// iso, _ := z.Shell(&cmd)

	// Hard-coded ISO
	// To-do: Fix receiving ISO from ifconfig

	cmd := []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`, `pacman-contrib`, `curl`}
	z.Arch_chroot(&cmd, false, c)

	cmd = []string{`pacman`, `-Syy`, `--noconfirm`, `--needed`, `reflector`, `rsync`, `arch-install-scripts`, `git`}
	z.Arch_chroot(&cmd, false, c)

	// log.Println("Enabling parallel downloading...")
	// cmd = []string{`sed`, `-i`, `s/^#ParallelDownloads/ParallelDownloads/g`, `/etc/pacman.conf`}
	// z.Arch_chroot(&cmd, false, c)

	// This specific command is giving errors on Arch-Chroot. Using host shell with hard-coded mount point
	log.Println("Enabling multilib libraries...")
	bash := `sed -i "/\[multilib\]/,/Include/"'s/^#//' /mnt/etc/pacman.conf`
	z.Shell(&bash)

	log.Println(`Backing up mirror list...`)
	cmd = []string{`cp`, `/etc/pacman.d/mirrorlist`, `/etc/pacman.d/mirrorlist.bak`}
	z.Arch_chroot(&cmd, false, c)

	log.Println(`Detecting best mirrors...`)
	cmd = []string{`reflector`, `--protocol`, `https`, `--country`, iso, `--latest`, `10`, `--sort`, `rate`, `--ipv6`, `--fastest`, `5`, `--save`, `/etc/pacman.d/mirrorlist`}
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

	log.Println(`Setting TimeZone...`)
	cmd = []string{`timedatectl`, `set-timezone`, c.Timezone}
	z.Arch_chroot(&cmd, false, c)

	log.Println(`Enabling Network Time Protocol...`)
	cmd = []string{`timedatectl`, `set-ntp`, `true`}
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
	cmd_list = append(cmd_list, c.Pacman.Packages...)
	z.Arch_chroot(&cmd_list, false, c)

	switch u.GetHostStatus().Chassis {
	case "vm":
		log.Println("Virtual machine detected! Installing open-vm-tools")
		cmd_list = []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`, `open-vm-tools`, `spice-vdagent`}
		z.Arch_chroot(&cmd_list, false, c)

		cmd_list = []string{`systemctl`, `enable`,
			`vmtoolsd.service`, `vmware-vmblock-fuse.service`,
			`spice-vdagentd.service`, `spice-vdagentd.socket`, `spice-webdavd.service`,
		}
		z.Arch_chroot(&cmd_list, false, c)
	}

	log.Println("Adding sudo no password rights...")
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
		case "apparmor":
			cmd_list = append(cmd_list,
				`apparmor.service`,
			)
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
		case "networkmanager":
			cmd_list = append(cmd_list,
				`bluetooth.service`,
			)
		case "openssh":
			cmd_list = append(cmd_list,
				`sshd`,
			)
		case "reflector":
			cmd_list = append(cmd_list,
				`reflector.timer`,
			)
		case "sddm":
			cmd_list = append(cmd_list,
				`sddm`,
			)
		case "virt-manager":
			cmd_list = append(cmd_list,
				`libvirtd.service`,
			)
		}
	}

	// Don't forget to enable TRIM for SSDs
	if c.Storage.SystemDiskRota || strings.Contains(c.Storage.SystemDiskID, "SSD") {
		cmd_list = append(cmd_list,
			`fstrim.timer`,
		)
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
		log.Println("Adding custom repo to install ArchZFS...")
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

		cmd = []string{`awk`, `BEGIN{ printf "# Origin Server - Finland\n" >> "/etc/pacman.conf" }`}
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

		cmd = []string{`awk`, `BEGIN{ printf "# ArchZFS - United States\n" >> "/etc/pacman.conf" }`}
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

		cmd_list = append(cmd_list, service_names...)
		z.Arch_chroot(&cmd_list, false, c)

		err := os.Remove("/mnt/archzfs.gpg")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func SetupProcessor(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                            Setup Processor
	-------------------------------------------------------------------------
	`)

	cpu := u.ListCPU()

	cmd := []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`}

	for i := range cpu.Processor {
		switch cpu.Processor[i].Data {
		case "GenuineIntel":
			log.Printf("Detected CPU is: %q\n", cpu.Processor[i].Data)
			cmd = append(cmd, `intel-ucode`)
			z.Arch_chroot(&cmd, false, c)
		case "AuthenticAMD":
			log.Printf("Detected CPU is: %q\n", cpu.Processor[i].Data)
			cmd = append(cmd, `amd-ucode`)
			z.Arch_chroot(&cmd, false, c)

			// Note: amd_pstate is now built into the kernel as of Kernel 6.1
			// module_dir := filepath.Join("/", "mnt", "etc", "modprobe.d")
			// module_config := "amd.conf"

			// module_contents := []string{
			// 	`options amd_pstate replace=1`,
			// }

			// log.Printf("Creating %q...\n", module_config)
			// u.WriteFile(&module_dir, &module_config, &module_contents, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
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
					`lib32-vulkan-icd-loader`,
					`lib32-vulkan-radeon`,
					`libva-mesa-driver`,
					`mesa`,
					`rocm-opencl-runtime`,
					`vulkan-icd-loader`,
					`vulkan-radeon`,
					`xf86-video-amdgpu`,
				)
				c.Parameters = append(c.Parameters, `amdgpu.ppfeaturemask=0xffffffff`) // Adjust clocks and voltages (https://wiki.archlinux.org/title/AMDGPU#Boot_parameter)
				c.Parameters = append(c.Parameters, `amdgpu.abmlevel=0`)               // Disable panel_power_savings (https://lore.kernel.org/lkml/a1d2749b-8db5-46d1-bf60-7820902cfc8f@amd.com/T/)
			case "Intel Corporation":
				cmd = append(cmd,
					`lib32-mesa`,
					`lib32-vulkan-icd-loader`,
					`lib32-vulkan-intel`,
					`libva-intel-driver`,
					`libva-utils`,
					`mesa`,
					`vulkan-icd-loader`,
					`vulkan-intel`,
					`xf86-video-intel`,
				)
			}
		}
	}

	z.Arch_chroot(&cmd, false, c)
}

func SetupBiometrics(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                    Detecting Biometric Hardware
	-------------------------------------------------------------------------
	`)

	usb := u.ListUSB()
	bio := u.BiometricIDs()

	cmd := []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`}

	// PAM Module Config(s): kde[✓], system-local-login[✓], login[✓], su[✘], sudo[✓], sddm[✘]
	pam_modules := []string{`login`, `system-local-login`, `sudo`}
	pam_module := []string{}

	switch c.Desktop.Environment {
	case "kde":
		pam_modules = append(pam_modules, `kde`)
	}

	// for i := range c.Pacman.Packages {
	// 	switch c.Pacman.Packages[i] {
	// 	case "sddm":
	// 		pam_modules = append(pam_modules, `sddm`)
	// 	}
	// }

	udev_rule_dir := filepath.Join("/", "mnt", "etc", "udev", "rules.d")
	u2f_udev_file := "70-u2f.rules"
	udev_rules := []string{`ACTION!="add|change", GOTO="fido_end"`}

	for i := range usb.USBDevices {

		if strings.Contains(usb.USBDevices[i].Description, `Camera`) {
			for j := range bio.Face {
				if usb.USBDevices[i].ID == bio.Face[j] {
					log.Println(`howdy compatible device found: `, usb.USBDevices[i].ID)
				}
			}

			// To-do:
			// - Add howdy PAM Authentication
		}

		if strings.Contains(usb.USBDevices[i].Description, `Fingerprint`) {
			// https://wiki.archlinux.org/title/Fprint
			for j := range bio.Fingerprint {
				if usb.USBDevices[i].ID == bio.Fingerprint[j] {
					log.Println(`fprint compatible device found: `, usb.USBDevices[i].ID)
					cmd = append(cmd, `fprintd`, `libfprint`)
				}
			}

			log.Println("Configuring PAM Authentication modules for fprint")
			pam_module_dir := filepath.Join("/", "mnt", "etc", "pam.d")

			for i := range pam_modules {
				pam_file := pam_modules[i]

				switch pam_modules[i] {
				case "kde":
					pam_module = []string{
						`#%PAM-1.0`,
						``,
						`auth            sufficient      pam_unix.so try_first_pass likeauth nullok`,
						`auth            sufficient      pam_fprintd.so`,
						``,
						`auth            include         system-login`,
						``,
						`account         include         system-login`,
						``,
						`password        include         system-login`,
						``,
						`session         include         system-login`,
					}
				case "login":
					pam_module = []string{
						`#%PAM-1.0`,
						``,
						`auth            sufficient      pam_unix.so try_first_pass likeauth nullok`,
						`auth            sufficient      pam_fprintd.so`,
						``,
						`auth       required     pam_securetty.so`,
						`auth       requisite    pam_nologin.so`,
						`auth       include      system-local-login`,
						`account    include      system-local-login`,
						`session    include      system-local-login`,
						`password   include      system-local-login`,
					}
				case "system-local-login":
					pam_module = []string{
						`#%PAM-1.0`,
						``,
						`auth            sufficient      pam_unix.so try_first_pass likeauth nullok`,
						`auth            sufficient      pam_fprintd.so`,
						``,
						`auth      include   system-login`,
						`account   include   system-login`,
						`password  include   system-login`,
						`session   include   system-login`,
					}
				case "sudo":
					pam_module = []string{
						`#%PAM-1.0`,
						``,
						`auth            sufficient      pam_unix.so try_first_pass likeauth nullok`,
						`auth            sufficient      pam_fprintd.so`,
						``,
						`auth            include         system-auth`,
						`account         include         system-auth`,
						`session         include         system-auth`,
					}
				}
				log.Printf("Adding fprint PAM Authentication for '%q'\n", pam_modules[i])
				u.WriteFile(&pam_module_dir, &pam_file, &pam_module, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
				u.ReadFile(&pam_module_dir, &pam_file)
			}
		}

		for j := range bio.SecurityKey {
			if usb.USBDevices[i].ID == bio.SecurityKey[j] {
				log.Println(`fido compatible device found: `, usb.USBDevices[i].ID)
				cmd = append(cmd, `libfido2`, `pam-u2f`)
			}

			// To-do:
			// - Add FIDO PAM Authentication modules
		}
	}
	z.Arch_chroot(&cmd, false, c)

	log.Println(`Adding udev rules for FIDO U2F...`)
	// Examples:
	// https://github.com/Yubico/libfido2/blob/main/udev/70-u2f.rules
	// https://www.trustkeysolutions.com/support/
	for i := range bio.SecurityKey {
		udev_rules = append(udev_rules, ``,
			fmt.Sprintf(`KERNEL=="hidraw*", SUBSYSTEM=="hidraw", ATTRS{idVendor}=="%s", ATTRS{idProduct}=="%s", TAG+="uaccess", GROUP="plugdev", MODE="0660"`, strings.Split(bio.SecurityKey[i], ":")[0], strings.Split(bio.SecurityKey[i], ":")[1]),
		)
	}
	udev_rules = append(udev_rules, ``, `LABEL="fido_end"`)
	u.WriteFile(&udev_rule_dir, &u2f_udev_file, &udev_rules, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	u.ReadFile(&udev_rule_dir, &u2f_udev_file)

	// To-do:
	// - Add more scalable way to handle a "sed-like" implementation in Go.
	// --- howdy and fprint can conflict with current overwrite method
}

func SetupMiscHardware(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                        Detecting Misc Hardware
	-------------------------------------------------------------------------
	`)

	// Define regular expression for RGB hardware
	var (
		rgbRegExp = regexp.MustCompile(`^RGB$|^Aura$|^Fusion$|^Mystic Light$|^Polychrome$`)
	)

	p := u.ListPCI()
	usb := u.ListUSB()

	for i := range p.PCIDevices {
		// Do nothing for now
		_ = i
	}

	for i := range usb.USBDevices {
		switch {
		case rgbRegExp.MatchString(usb.USBDevices[i].Description):
			log.Printf("'%s' detected!\n", usb.USBDevices[i].Description)
			log.Println("Installing RGB compatible packages(s) + udev rules...")
			c.Flatpak.Packages = append(c.Flatpak.Packages, `org.openrgb.OpenRGB`)

			log.Println("Manually load 'i2c-dev' module to reduce RGB errors")
			module_dir := filepath.Join("/", "mnt", "etc", "modules-load.d")
			module_config := "i2c_dev.conf"

			module_contents := []string{
				`i2c-dev`,
			}

			log.Printf("Creating %q...\n", module_config)
			u.WriteFile(&module_dir, &module_config, &module_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite

			// OpenRGB Flatpak instructions: https://github.com/flathub/org.openrgb.OpenRGB
			cmd := []string{`wget`, `'https://gitlab.com/CalcProgrammer1/OpenRGB/-/jobs/artifacts/master/raw/60-openrgb.rules?job=Linux+64+AppImage&inline=false'`, `/usr/lib/udev/rules.d/`, `-O`, `60-openrgb.rules`}
			z.Arch_chroot(&cmd, false, c)
		case strings.Contains(usb.USBDevices[i].Description, `NZXT Kraken`):
			log.Printf("'%s' detected!\n", usb.USBDevices[i].Description)

			log.Println("Note: [As of 12/23/22] NZXT Kraken X53/X63/X73 is not native to the mainline kernel")
			// https://github.com/liquidctl/liquidtux#installing-with-dkms
			c.Pacman.AUR.Packages = append(c.Pacman.AUR.Packages, `liquidtux-dkms-git`)

			log.Println("Manually configuring to load at boot...")
			module_dir := filepath.Join("/", "mnt", "etc", "modules-load.d")
			module_config := "nzxt.conf"

			module_contents := []string{
				`nzxt-kraken3`,
			}

			log.Printf("Creating %q...\n", module_config)
			u.WriteFile(&module_dir, &module_config, &module_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite

			log.Printf("Appending '%s' compatible package(s)...\n", usb.USBDevices[i].Description)
		case strings.HasPrefix(usb.USBDevices[i].ID, "1532:"):
			log.Printf("'%s' detected!\n", usb.USBDevices[i].Description)
			log.Println("Appending Razer compatible package(s)...")

			c.Pacman.Packages = append(c.Pacman.Packages,
				"openrazer-daemon",      // Userspace daemon that abstracts access to the kernel driver. Provides a DBus service for applications to use
				"openrazer-driver-dkms", // OpenRazer kernel modules sources
				"python-openrazer",      // Library for interacting with the OpenRazer daemon
			)

			c.Pacman.AUR.Packages = append(c.Pacman.AUR.Packages,
				"polychromatic", // RGB management GUI for Razer Devices
			)
		case strings.Contains(usb.USBDevices[i].Description, `Razer USA, Ltd Nari`):
			log.Println("Appending Razer Nari compatible package(s)...")

			c.Pacman.AUR.Packages = append(c.Pacman.AUR.Packages,
				"razer-nari-pipewire-profile", // Razer Nari headsets pipewire profile
			)
		}
	}

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
				`libvirt`,
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
		`.cache`, `.config/autostart`, `.local/share/applications`, `.local/share/icons`,
		`Applications`, `Desktop`, `Developer`, `Documents`, `Downloads`, `Music`,
		`Pictures`, `Public`, `Templates`, `Videos`,
	}

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "code":
			user_directories = append(user_directories,
				`.config/Code - OSS/User`,
				`.vscode-oss`,
			)
		}
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

func SetupSecurityModules(c *conf.Config) {

	log.Println(`
	-------------------------------------------------------------------------
                    Setup Linux Security Modules
	-------------------------------------------------------------------------
	`)

	// Currently only supporting AppArmor
	// ToDo: Will provde considerations for SElinux

	cmd := []string{`pacman`, `-Syy`, `--needed`, `--noconfirm`}

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "apparmor":
			log.Println("AppArmor Linux Security Module detected...")

			log.Println(`Installing Audit framework...`)
			cmd = append(cmd,
				`audit`,
				`python-notify2`,
				`python-psutil`,
			)
			z.Arch_chroot(&cmd, false, c)

			log.Println(`Enabling Audit service...`)
			cmd = []string{`systemctl`, `enable`, `--now`, `auditd.service`}
			z.Arch_chroot(&cmd, false, c)

			log.Println(`Creating autostart for AppArmor Notify service`)
			autostart_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "autostart")
			autostart_file := "apparmor-notify.desktop"
			autostart_contents := []string{
				`[Desktop Entry]`,
				`Comment[en_US]=Receive on screen notifications of AppArmor denials`,
				`Comment=Receive on screen notifications of AppArmor denials`,
				`Exec=aa-notify -p -s 1 -w 60 -f /var/log/audit/audit.log`,
				`GenericName[en_US]=`,
				`GenericName=`,
				`Icon=preferences-security-apparmor`,
				`MimeType=`,
				`Name[en_US]=AppArmor Notify`,
				`Name=AppArmor Notify`,
				`NoDisplay=true`,
				`Path=`,
				`StartupNotify=false`,
				`Terminal=false`,
				`TerminalOptions=`,
				`TryExec=aa-notify`,
				`Type=Application`,
				`X-DBUS-ServiceName=`,
				`X-DBUS-StartupType=`,
				`X-KDE-SubstituteUID=false`,
				`X-KDE-Username=`,
			}
			u.WriteFile(&autostart_dir, &autostart_file, &autostart_contents, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		}

		log.Println(`Creating config for Audit daemon...`)
		audit_dir := filepath.Join("/", "mnt", "etc", "audit")
		audit_file := "auditd.conf"
		audit_contents := []string{
			`#`,
			`# This file controls the configuration of the audit daemon`,
			`#`,
			`local_events = yes`,
			`write_logs = yes`,
			`log_file = /var/log/audit/audit.log`,
			`log_group = adm`,
			`log_format = ENRICHED`,
			`flush = INCREMENTAL_ASYNC`,
			`freq = 50`,
			`max_log_file = 8`,
			`num_logs = 5`,
			`priority_boost = 4`,
			`name_format = NONE`,
			`##name = mydomain`,
			`max_log_file_action = ROTATE`,
			`space_left = 75`,
			`space_left_action = SYSLOG`,
			`verify_email = yes`,
			`action_mail_acct = root`,
			`admin_space_left = 50`,
			`admin_space_left_action = SUSPEND`,
			`disk_full_action = SUSPEND`,
			`disk_error_action = SUSPEND`,
			`use_libwrap = yes`,
			`##tcp_listen_port = 60`,
			`tcp_listen_queue = 5`,
			`tcp_max_per_addr = 1`,
			`##tcp_client_ports = 1024-65535`,
			`tcp_client_max_idle = 0`,
			`transport = TCP`,
			`krb5_principal = auditd`,
			`##krb5_key_file = /etc/audit/audit.key`,
			`distribute_network = no`,
			`q_depth = 1200`,
			`overflow_action = SYSLOG`,
			`max_restarts = 10`,
			`plugin_dir = /etc/audit/plugins.d`,
			`end_of_event_timeout = 2`,
		}
		u.WriteFile(&audit_dir, &audit_file, &audit_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
	}
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

	log.Println("Making AUR Package...")
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

	log.Println("Compiling Flatpak commands...")
	fp_install_cmd := `sudo flatpak install --verbose --assumeyes flathub`
	fp_override_cmd := `sudo flatpak override`
	cmd_list := []string{`sleep 3`}
	for i := range c.Flatpak.Packages {
		cmd_list = append(cmd_list, fmt.Sprintf(`%s %s`, fp_install_cmd, c.Flatpak.Packages[i]))

		switch c.Flatpak.Packages[i] {
		case "com.brave.Browser", "com.google.Chrome", "com.microsoft.Edge":
			log.Println("Chromium browser Flatpak detected!")
			log.Println("Adding permissions for Progressive Web Apps")
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=/home/%s/.local/share/applications`, fp_override_cmd, c.Flatpak.Packages[i], c.User.Username))
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=/home/%s/.local/share/icons`, fp_override_cmd, c.Flatpak.Packages[i], c.User.Username))
		case "com.github.wwmm.easyeffects":
			log.Println("EasyEffects for PipeWire detected!")
			log.Println("Configuring AutoStart for EasyEffects")

			autostart_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "autostart")
			autostart_file := "com.github.wwmm.easyeffects.desktop"
			autostart_contents := []string{
				`[Desktop Entry]`,
				`Comment[en_US]=`,
				`Comment=`,
				`Exec=flatpak run --command=easyeffects com.github.wwmm.easyeffects --gapplication-service`,
				`GenericName[en_US]=`,
				`GenericName=`,
				`Icon=com.github.wwmm.easyeffects`,
				`MimeType=`,
				`Name[en_US]=com.github.wwmm.easyeffects`,
				`Name=com.github.wwmm.easyeffects`,
				`Path=`,
				`StartupNotify=true`,
				`Terminal=false`,
				`TerminalOptions=`,
				`Type=Application`,
				`X-DBUS-ServiceName=`,
				`X-DBUS-StartupType=`,
				`X-Flatpak=com.github.wwmm.easyeffects`,
				`X-KDE-SubstituteUID=false`,
				`X-KDE-Username=`,
			}
			u.WriteFile(&autostart_dir, &autostart_file, &autostart_contents, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

			// Insert code for Dolby Atmos here
		case "com.getmailspring.Mailspring":
			log.Println("Mailspring Flatpak detected!")
			log.Println("Adding permissions for Freedesktop.org Secret Service Integration")
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --socket=session-bus`, fp_override_cmd, c.Flatpak.Packages[i]))
		case "com.synology.SynologyDrive":
			log.Println("Synology Drive Flatpak detected!")
			log.Println("Configuring AutoStart for Synology Drive")

			autostart_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "autostart")
			autostart_file := "com.synology.SynologyDrive.desktop"
			autostart_contents := []string{
				`[Desktop Entry]`,
				`Comment[en_US]=`,
				`Comment=`,
				`Exec=flatpak run com.synology.SynologyDrive`,
				`GenericName[en_US]=`,
				`GenericName=`,
				`Icon=com.synology.SynologyDrive`,
				`MimeType=`,
				`Name[en_US]=com.synology.SynologyDrive`,
				`Name=com.synology.SynologyDrive`,
				`Path=`,
				`StartupNotify=true`,
				`Terminal=false`,
				`TerminalOptions=`,
				`Type=Application`,
				`X-DBUS-ServiceName=`,
				`X-DBUS-StartupType=`,
				`X-Flatpak=com.synology.SynologyDrive`,
				`X-KDE-SubstituteUID=false`,
				`X-KDE-Username=`,
			}
			u.WriteFile(&autostart_dir, &autostart_file, &autostart_contents, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		case "com.usebottles.bottles":
			log.Println("Bottles Flatpak detected!")
			log.Println("Adding permissions for Desktop/Steam Entries")

			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=xdg-data/applications`, fp_override_cmd, c.Flatpak.Packages[i]))
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=/home/%s/.local/share/Steam`, fp_override_cmd, c.Flatpak.Packages[i], c.User.Username))
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=/home/%s/.var/app/com.valvesoftware.Steam/data/Steam`, fp_override_cmd, c.Flatpak.Packages[i], c.User.Username))
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=/home/%s/Applications`, fp_override_cmd, c.Flatpak.Packages[i], c.User.Username))
		case "org.keepassxc.KeePassXC":
			log.Println("KeePassXC Flatpak detected!")
			log.Println("Configuring AutoStart for KeePassXC")

			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --filesystem=xdg-data/applications`, fp_override_cmd, c.Flatpak.Packages[i]))
			autostart_dir := filepath.Join("/", "mnt", "home", c.User.Username, ".config", "autostart")
			autostart_file := "org.keepassxc.KeePassXC.desktop"
			autostart_contents := []string{
				`[Desktop Entry]`,
				`Comment[en_US]=`,
				`Comment=`,
				`Exec=flatpak run org.keepassxc.KeePassXC`,
				`GenericName[en_US]=`,
				`GenericName=`,
				`Icon=org.keepassxc.KeePassXC`,
				`MimeType=`,
				`Name[en_US]=org.keepassxc.KeePassXC`,
				`Name=org.keepassxc.KeePassXC`,
				`Path=`,
				`StartupNotify=true`,
				`Terminal=false`,
				`TerminalOptions=`,
				`Type=Application`,
				`X-DBUS-ServiceName=`,
				`X-DBUS-StartupType=`,
				`X-Flatpak=org.keepassxc.KeePassXC`,
				`X-KDE-SubstituteUID=false`,
				`X-KDE-Username=`,
			}
			u.WriteFile(&autostart_dir, &autostart_file, &autostart_contents, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

			log.Println("Adding permissions for Freedesktop.org Secret Service Integration")
			cmd_list = append(cmd_list, fmt.Sprintf(`%s %s --socket=session-bus`, fp_override_cmd, c.Flatpak.Packages[i]))
		}
	}
	log.Println("Appending systemd-nspawn 'Get out of Jail for free' command...")
	cmd_list = append(cmd_list, `sudo poweroff`)

	log.Println("Ensuring Flatpak will automatically execute after mounting systemd-nspawn container...")

	systemd_autorun_dir := filepath.Join("/", "mnt", "etc", "profile.d")
	flatpak_script := "install_flatpaks.sh"

	log.Println(`Creating Flatpak script for systemd-nspawn container...`)
	u.WriteFile(&systemd_autorun_dir, &flatpak_script, &cmd_list, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	log.Println("Making executable...")
	cmd := fmt.Sprintf(`chmod +x %v`, filepath.Join(systemd_autorun_dir, flatpak_script))
	z.Shell(&cmd)

	log.Println("Installing Flatpaks via systemd-nspawn...")
	z.Systemd_nspawn(&[]string{}, true, c)

	log.Println("Cleaning up cruft...")
	log.Println("Deleting: ", autologin_dir)
	err = os.RemoveAll(autologin_dir)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Deleting: ", filepath.Join(systemd_autorun_dir, flatpak_script))
	err = os.Remove(filepath.Join(systemd_autorun_dir, flatpak_script))
	if err != nil {
		log.Fatalln(err)
	}
}

func SetupEFI(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                        Adding EFI Configs
	-------------------------------------------------------------------------
	`)

	var uuid_command string
	if strings.HasPrefix(c.Storage.SystemDisk, "/dev/nvme") {
		uuid_command = fmt.Sprintf(`lsblk -dno UUID %vp2`, c.Storage.SystemDisk)
	} else {
		uuid_command = fmt.Sprintf(`lsblk -dno UUID %v2`, c.Storage.SystemDisk)
	}
	root_uuid := strings.TrimRight(strings.TrimSpace(z.Shell(&uuid_command)), "\n")

	switch c.Storage.Filesystem {
	case "btrfs":
		c.Modules = append(c.Modules, "btrfs")
		c.Parameters = append(c.Parameters, fmt.Sprintf(`root=UUID=%v rootflags=subvol=@`, root_uuid))
	case "luks":
		c.Modules = append(c.Modules, "btrfs") // Since `luks` is currently treated as `btrfs`
		c.Parameters = append(c.Parameters, fmt.Sprintf(`rd.luks.name=%v=luks_ROOT root=/dev/mapper/luks_ROOT rootflags=subvol=@ rd.luks.options=%v=timeout=15s,discard,quiet,rw`, root_uuid, root_uuid))
	case "zfs":
		c.Parameters = append(c.Parameters, `root=ZFS=zroot/ROOT/default`)
	}

	cpu := u.ListCPU()

	for i := range cpu.Processor {
		switch cpu.Processor[i].Data {
		case "GenuineIntel":
			c.Modules = append(c.Modules, `intel_pstate`)
			c.Parameters = append(c.Parameters, `intel_pstate=active`)
		case "AuthenticAMD":
			c.Modules = append(c.Modules, `amd_pstate`)
			// Note: `amd_pstate={active,guided}` implemented for Kernel v6.3+
			c.Parameters = append(c.Parameters, `amd_pstate=guided`)
		}
	}

	for i := range c.Pacman.Packages {
		switch c.Pacman.Packages[i] {
		case "apparmor":
			c.Parameters = append(c.Parameters, `lsm=landlock,lockdown,yama,integrity,apparmor,bpf`)
		}
	}

	p := u.ListPCI()

	for i := range p.PCIDevices {
		if p.PCIDevices[i].Class == "VGA compatible controller" {
			switch p.PCIDevices[i].Vendor {
			case "NVIDIA Corporation":
				c.Modules = append(c.Modules, `nvidia`, `nvidia_modeset`, `nvidia_uvm`, `nvidia_drm`)
			case "Advanced Micro Devices, Inc. [AMD]", "Advanced Micro Devices, Inc. [AMD/ATI]":
				c.Modules = append(c.Modules, `amdgpu`)
			case "Intel Corporation":
				c.Modules = append(c.Modules, `i915`)
			}
		}
	}

	switch c.Bootloader {
	case "grub": // GRand Unified Bootloader
		cmd := []string{`sudo`, `pacman`, `-Syy`, `--needed`, `--noconfirm`, c.Bootloader}
		z.Arch_chroot(&cmd, false, c)

		grub_path := filepath.Join("/", "mnt", "etc", "default")
		grub_file := `grub`
		grub_contents := u.ReadFile(&grub_path, &grub_file)

		for line := range *grub_contents {
			switch {
			case strings.HasPrefix((*grub_contents)[line], `GRUB_CMDLINE_LINUX=""`):
				(*grub_contents)[line] =
					fmt.Sprintf("GRUB_CMDLINE_LINUX=\"loglevel 3 video=1920x1080 %v\"", strings.Join(c.Parameters, " "))
			case strings.HasPrefix((*grub_contents)[line], `#GRUB_ENABLE_CRYPTODISK=`):
				switch c.Storage.Filesystem {
				case "luks":
					(*grub_contents)[line] = strings.TrimPrefix((*grub_contents)[line], "#")
				}
			}
		}
		u.WriteFile(&grub_path, &grub_file, grub_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
		u.ReadFile(&grub_path, &grub_file)

		cmd = []string{`grub-install`, `--target=x86_64-efi`, `--efi-directory=/boot`, `--bootloader-id=ArchLinux`}
		z.Arch_chroot(&cmd, false, c)

		cmd = []string{`grub-mkconfig`, `-o`, `/boot/grub/grub.cfg`}
		z.Arch_chroot(&cmd, false, c)

		grub_path = filepath.Join("/", "mnt", "boot", "grub")
		grub_file = `grub.cfg`
		u.ReadFile(&grub_path, &grub_file)
	case "systemd-boot":
		cmd := []string{`bootctl`, `--path=/boot`, `install`}
		z.Arch_chroot(&cmd, false, c)

		// Build boot entry for Arch
		boot_entry_dir := filepath.Join("/", "mnt", "boot", "loader", "entries")
		boot_entry := "arch.conf"
		boot_config := []string{
			`title ArchLinux`,
			fmt.Sprintf(`linux /vmlinuz-%v`, c.Kernel),
		}

		for i := range cpu.Processor {
			switch cpu.Processor[i].Data {
			case "GenuineIntel":
				boot_config = append(boot_config, `initrd /intel-ucode.img`)
			case "AuthenticAMD":
				boot_config = append(boot_config, `initrd /amd-ucode.img`)
			}
		}

		boot_config = append(boot_config, fmt.Sprintf(`initrd /initramfs-%v.img`, c.Kernel))
		boot_config = append(boot_config, fmt.Sprintf(`options %v`, strings.Join(c.Parameters, " ")))
		u.WriteFile(&boot_entry_dir, &boot_entry, &boot_config, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
		u.ReadFile(&boot_entry_dir, &boot_entry)

		// Configure system-d bootloader
		boot_loader_dir := filepath.Join("/", "mnt", "boot", "loader")
		boot_loader := "loader.conf"
		boot_loader_config := []string{
			fmt.Sprintf(`default %v`, boot_entry),
			`timeout 3`,
			`console-mode max`,
			`editor 0`,
		}
		u.WriteFile(&boot_loader_dir, &boot_loader, &boot_loader_config, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
		u.ReadFile(&boot_loader_dir, &boot_loader)
	}

	// Configure mkinitcpio
	mkinitcpio_dir := filepath.Join("/", "mnt", "etc")
	mkinitcpio_conf := "mkinitcpio.conf"

	mkinitcpio_contents := u.ReadFile(&mkinitcpio_dir, &mkinitcpio_conf)

	// Determine HOOKS
	switch c.Storage.Filesystem {
	case "btrfs", "luks":
		c.Hooks = []string{
			`base`,
			`udev`,
			`systemd`,
			`autodetect`,
			`keyboard`,
			`modconf`,
			`block`,
			`sd-encrypt`,
			`filesystems`,
			`shutdown`,
		}
	case "zfs":
		c.Hooks = []string{
			`base`,
			`udev`,
			`systemd`,
			`autodetect`,
			`modconf`,
			`block`,
			`keyboard`,
			`zfs`,
			`filesystems`,
			`shutdown`,
		}
	}

	// Save to file
	for line := range *mkinitcpio_contents {
		switch {
		case strings.HasPrefix((*mkinitcpio_contents)[line], `MODULES=()`):
			(*mkinitcpio_contents)[line] = fmt.Sprintf(`MODULES=(%v)`, strings.Join(c.Modules, " "))
		case strings.HasPrefix((*mkinitcpio_contents)[line], `HOOKS=(`):
			(*mkinitcpio_contents)[line] = fmt.Sprintf(`HOOKS=(%v)`, strings.Join(c.Hooks, " "))
		}
	}
	u.WriteFile(&mkinitcpio_dir, &mkinitcpio_conf, mkinitcpio_contents, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755) // Overwrite
	u.ReadFile(&mkinitcpio_dir, &mkinitcpio_conf)

	cmd := []string{`mkinitcpio`, `-p`, c.Kernel}
	z.Arch_chroot(&cmd, false, c)
}

func SetupSecureBoot(c *conf.Config) {
	log.Println(`
	-------------------------------------------------------------------------
                        Setup Secure Boot
	-------------------------------------------------------------------------
	`)

	sb_status := u.SecureBootStatus()
	log.Println(u.PrettyJson(sb_status))

	efi_files := []string{}

	// Paths are not prefixed with `/mnt` since we are using chroot
	switch c.Bootloader {
	case "grub":
		efi_files = append(efi_files, `/boot/EFI/ArchLinux/grubx64.efi`)
	case "systemd-boot":
		efi_files = append(efi_files, `/boot/EFI/BOOT/BOOTX64.EFI`)
		efi_files = append(efi_files, fmt.Sprintf(`/boot/EFI/Linux/linux-%v.efi`, c.Kernel))
		efi_files = append(efi_files, `/boot/EFI/systemd/systemd-bootx64.efi`)
	}

	efi_files = append(efi_files, fmt.Sprintf(`/boot/vmlinuz-%v`, c.Kernel))

	switch sb_status.SetupMode {
	case "Enabled":
		log.Println(`"Setup Mode" is enabled for Secure Boot!`)
		z.Arch_chroot(u.SecureBootCreateKeys(), false, c)
		z.Arch_chroot(u.SecureBootEnrollKeys(), false, c)

		for i := range efi_files {
			z.Arch_chroot(u.SecureBootSign(&efi_files[i]), false, c)
		}

		switch c.Storage.Filesystem {
		case "luks":
			var disk1_part2 string
			switch strings.HasPrefix(c.Storage.SystemDisk, "/dev/nvme") {
			case true:
				disk1_part2 = fmt.Sprintf(`%vp2`, c.Storage.SystemDisk)
			case false:
				disk1_part2 = fmt.Sprintf(`%v2`, c.Storage.SystemDisk)
			}

			cmd := fmt.Sprintf(`export PASSWORD='%v' && systemd-cryptenroll --wipe-slot=tpm2 --tpm2-device=auto --tpm2-pcrs=0,7 %v`, c.Storage.EncryptionKey, disk1_part2)
			z.Shell(&cmd)
		}
	case "Disabled":
		log.Println(`WARNING: "Setup Mode" is disabled for Secure Boot!`)
		log.Println(`WARNING: Ensure to delete all keys or enable Setup Mode in your System BIOS after installation`)
		log.Println(`WARNING: After preparations, to use Secure Boot, please run the following commands:`)
		for i := range efi_files {
			log.Printf(`sudo sbctl sign -s %q\n`, strings.TrimPrefix(efi_files[i], "/mnt"))
		}

		switch c.Storage.Filesystem {
		case "luks":
			log.Println(`EXTRA: Run "systemd-cryptenroll --tpm2-device=auto --tpm2-pcrs=0,7" to save LUKS key to TPM`)
		}
	}

	// To-do:
	// - Add post-install hook for kernel upgrades and auto-signing
	// - Add DKMS Kernel Module Signing
}
