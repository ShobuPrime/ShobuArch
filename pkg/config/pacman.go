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
package config

func BasePkgs(c *Config) {

	c.Pacman.Packages = append(c.Pacman.Packages,
		"cockpit",                // A systemd web based user interface for Linux servers
		"cockpit-machines",       // Cockpit UI for virtual machines
		"cockpit-packagekit",     // Packages UI for Cockpit
		"cockpit-storaged",       // Storage UI for Cockpit
		"code",                   // The Open Source build of Visual Studio Code (vscode) editor
		"dmidecode",              // Desktop Management Interface table related utilities
		"dnsmasq",                // Lightweight, easy to configure DNS forwarder and DHCP server
		"docker",                 // Pack, ship and run any application as a lightweight container
		"efibootmgr",             // Linux user-space application to modify the EFI Boot Manager
		"firefox",                // "Default" Web Browser
		"firewalld",              // Firewall daemon with D-Bus interface
		"flatpak",                // Sandboxed desktop applications on Linux.
		"fwupd",                  // Daemon to update some devices' firmware
		"go",                     // Core compiler tools for the Go programming language
		"keepassxc",              // Cross-platform community-driven port of Keepass password manager
		"libdbusmenu-glib",       // Library for passing menus over DBus
		"lib32-pipewire",         // Low-latency audio/video router and processor -- 32bit
		"lib32-pipewire-jack",    // Pipewire JACK support -- 32bit
		"lib32-pipewire-v4l2",    // Pipewire V4L2 interceptor -- 32bit
		"networkmanager",         // Network connection manager and user applications
		"network-manager-applet", // Applet for managing network connections
		"obs-studio",             // Free, open source software for live streaming and recording
		"openssh",                // Premier connectivity tool for remote login with the SSH protocol
		"os-prober",              // Utility to detect other OSes on a set of drives
		"packagekit",             // A system designed to make installation and updates of packages easier
		"pipewire",               // Low-latency audio/video router and processor
		"pipewire-docs",          // Pipewire Documentation
		"pipewire-jack",          // Pipewire JACK support
		"pipewire-pulse",         // Pipewire PulseAudio replacement
		"pipewire-v4l2",          // Pipewire V4L2 interceptor
		"qemu-full",              // A generic and open source machine emulator and virtualizer
		"qpwgraph",               // PipeWire Graph Qt GUI Interface
		"reflector",              // A Python 3 module and script to retrieve and filter the latest Pacman mirror list.
		"rsync",                  // A fast and versatile file copying tool for remote and local files
		"steam",                  // Valve's digital software delivery system
		"terminus-font",          // Monospace bitmap font (for X11 and console)
		"virt-manager",           // Desktop user interface for managing virtual machines
		"v4l2loopback-dkms",      // Virtual Camera for OBS-Studio
		"wget",                   // Network utility to retrieve files from the
		"wireplumber",            // Session & policy manager implementation for PipeWire
		"wireplumber-docs",       // Documentation
		"wpa_supplicant",         // A utility providing key negotiation for WPA wireless networks
		"xdg-desktop-portal-gtk", // Prevent blurry text from GTK Flatpaks
		"xdg-user-dirs",          // Manage user directories like ~/Desktop and ~/Music
		"xdg-utils",              // Command line tools that assist applications with a variety of desktop integration tasks
	)

	SaveConfig(c)
}

func Gnome(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"gnome",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"gnome",
			"gnome-extra",
		)
	}

	SaveConfig(c)
}

func KDE(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"plasma-desktop",
			"sddm",
		)
	case "FULL":
		// Consider `kde-applications-meta` if you want the default "Full" experience
		c.Pacman.Packages = append(c.Pacman.Packages,
			"ark",                         // Archiving tool
			"audiocd-kio",                 // Kioslave for accessing audio CDs
			"bluedevil",                   // Integrate the Bluetooth technology within KDE workspace and applications
			"breeze-gtk",                  // Breeze widget theme for GTK 2 and 3
			"discover",                    // KDE and Plasma resources management GUI
			"dolphin",                     // KDE File Manager
			"dragon",                      // A simple multimedia player
			"drkonqi",                     // The KDE crash handler
			"ffmpegthumbs",                // Video thumbnail generator for KDE file managers.
			"filelight",                   // View disk usage information
			"gwenview",                    // Fast and easy to use image viewer
			"juk",                         // A jukebox, tagger and music collection manager
			"k3b",                         // CD burning application
			"kamera",                      // KDE integration for gphoto2 cameras
			"kamoso",                      // A webcam recorder
			"kalarm",                      // Personal alarm scheduler
			"kalendar",                    // A calendar application using Akonadi to sync
			"kate",                        // Advanced Text Editor
			"kbackup",                     // Backup directories or files
			"kcalc",                       // Scientific Calculator
			"kcharselect",                 // Character selector
			"kcolorchooser",               // Color Chooser
			"kcodecs",                     // Manipulate strings using various encodings
			"kcron",                       // Configure and schedule tasks
			"kde-gtk-config",              // GTK2 and GTK3 Configurator for KDE
			"kde-network-meta",            // KDE network applications
			"kdebugsettings",              // Enable/disable qCDebug
			"kdegraphics-thumbnailers",    // Thumbnailers for various graphics file formats
			"kdenlive",                    // Non-linear video editor for Linux
			"kdeplasma-addons",            // All kind of addons to improve your Plasma experience
			"kdialog",                     // Display dialog boxes from shell scripts
			"kdf",                         // View Disk Usage
			"keditbookmarks",              // Bookmark Organizer and Editor
			"kfind",                       // Find files/folders
			"kgamma5",                     // Adjust monitor gamma settings
			"kgpg",                        // GnuPG frontend
			"khelpcenter",                 // KDE Applications documentation
			"khotkeys",                    // KHotKeys
			"kimagemapeditor",             // HTML Image Map Editor
			"kinfocenter",                 // Provides information about a computer system
			"kleopatra",                   // Certificate Manager and Unified Crypto GUI
			"kmag",                        // Screen Magnifier
			"kmousetool",                  // Clicks the mouse for you, reducing the effects of RSI
			"kmouth",                      // Speech Synthesizer Frontend
			"kolourpaint",                 // Paint Program
			"konsole",                     // KDE terminal emulator
			"kontrast",                    // Check contrast for colors that allows verifying they are accessible
			"kruler",                      // Screen ruler
			"kscreen",                     // KDE screen management software
			"ksshaskpass",                 // ssh-add helper that uses kwallet and kpassworddialog
			"ksystemlog",                  // System log viewer tool
			"kteatime",                    // Handy timer for steeping tea
			"ktimer",                      // Countdown launcher
			"kwallet-pam",                 // KWallet PAM integration
			"kwalletmanager",              // Wallet management tool
			"kwave",                       // Sound editor
			"kwayland-integration",        // Integration plugins for KDE frameworks and Wayland
			"kwrited",                     // KDE daemon listening for wall and write messages
			"markdownpart",                // KPart for rendering Markdown content
			"okular",                      // Document Viewer
			"oxygen",                      // KDE Oxygen style
			"packagekit-qt5",              // Qt5 bindings for PackageKit
			"partitionmanager",            // Manage disks, partitions, and file systems
			"plasma-browser-integration",  // Integrate browsers into the Plasma Desktop
			"plasma-desktop",              // KDE Plasma Desktop
			"plasma-disks",                // Monitors S.M.A.R.T. capable devices
			"plasma-firewall",             // Control Panel for your system firewall
			"plasma-nm",                   // Plasma applet written in QML for managing network connections
			"plasma-pa",                   // Plasma applet for audio volume management using PulseAudio
			"plasma-systemmonitor",        // Interface for monitoring system sensors, process information and other system resources
			"plasma-thunderbolt",          // Control Thunderbolt devices
			"plasma-vault",                // Plasma applet and services for creating encrypted vaults
			"plasma-wayland-session",      // Plasma Wayland session
			"plasma-workspace-wallpapers", // Additional wallpapers for the Plasma Workspace
			"powerdevil",                  // Manages the power consumption settings of a Plasma Shell
			"print-manager",               // Tool for managing print jobs and printers
			"sddm",                        // QML based X11 and Wayland display manager
			"sddm-kcm",                    // KDE Config Module for SDDM
			"skanpage",                    // Utility to scan images and multi-page documents
			"spectacle",                   // KDE screenshot capture utility
			"svgpart",                     // A KPart for viewing SVGs
			"sweeper",                     // System Cleaner
			"xdg-desktop-portal-kde",      // Backend implementation for x-d-p using Qt/KF5
			"yakuake",                     // Drop-down terminal emulator based on KDE konsole technology
		)
	}

	SaveConfig(c)
}

func Cinnamon(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"cinnamon",
			"metacity",
			"gnome-shell",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"cinnamon",
			"metacity",
			"gnome-shell",
		)
	}

	SaveConfig(c)
}

func XFCE(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"xfce4",
			"xfce-goodies",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"xfce4",
			"xfce-goodies",
		)
	}

	SaveConfig(c)
}

func Mate(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"mate",
			"mate-extra",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"mate",
			"mate-extra",
		)
	}

	SaveConfig(c)
}

func Budgie(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"budgie-desktop",
			"gnome",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"budgie-desktop",
			"gnome",
		)
	}

	SaveConfig(c)
}

func LXDE(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"lxde",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"lxde",
		)
	}

	SaveConfig(c)
}

func Deepin(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"deepin",
			"deepin-extra",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"deepin",
			"deepin-extra",
		)
	}

	SaveConfig(c)
}

func Openbox(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"lightdm",
			"lightdm-gtk-greeter",
			"lxsession",
			"openbox",
			"rxvt-unicode",
			"thunar",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"alsa-utils",
			"dunst",
			"feh",
			"geany",
			"gsimplecal",
			"gtk2-perl",
			"htop",
			"imagemagick",
			"jq",
			"lightdm",
			"lightdm-gtk-greeter",
			"lightdm-webkit2-greeter",
			"lightdm-webkit-theme-litarvan",
			"lxappearance",
			"lxsession",
			"nano",
			"neofetch",
			"obconf",
			"openbox",
			"parcellite",
			"pavucontrol",
			"picom",
			"playerctl",
			"pulseaudio",
			"pulseaudio-alsa",
			"qt5ct",
			"rofi",
			"rxvt-unicode",
			"scrot",
			"thunar",
			"thunar-archive-plugin",
			"thunar-media-tags-plugin",
			"thunar-volman",
			"tint2",
			"tumbler",
			"viewnior",
			"w3m",
			"wireless_tools",
			"xautolock",
			"xclip",
			"xfce4-power-manager",
			"xsettingsd",
			"zsh",
		)
	}

	SaveConfig(c)
}

func Server(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"binutils",
			"dosfstools",
			"linux-headers",
			"mesa",
			"noto-fonts-emoji",
			"usbutils",
			"xdg-user-dirs",
			"xorg",
			"xorg-apps",
			"xorg-drivers",
			"xorg-server",
			"xorg-xinit",
			"xorg-xkill",
			"xterm",
		)
	case "FULL":
		c.Pacman.Packages = append(c.Pacman.Packages,
			"alsa-plugins",
			"alsa-utils",
			"autoconf",
			"automake",
			"awesome-terminal-fonts",
			"bash-completion",
			"bind",
			"binutils",
			"bison",
			"bluez",
			"bluez-libs",
			"bluez-utils",
			"bridge-utils",
			"btrfs-progs",
			"celluloid",
			"cmatrix",
			"code",
			"cronie",
			"cups",
			"dialog",
			"dmidecode",
			"dnsmasq",
			"dosfstools",
			"dtc",
			"efibootmgr",
			"egl-wayland",
			"exfat-utils",
			"flex",
			"fuse2",
			"fuse3",
			"fuseiso",
			"gamemode",
			"gcc",
			"gimp",
			"gparted",
			"gptfdisk",
			"grub-customizer",
			"gst-libav",
			"gst-plugins-good",
			"gst-plugins-ugly",
			"haveged",
			"htop",
			"jdk-openjdk",
			"kitty",
			"libdvdcss",
			"libtool",
			"linux-headers",
			"lsof",
			"lutris",
			"lzop",
			"m4",
			"make",
			"mesa",
			"neofetch",
			"noto-fonts-emoji",
			"ntfs-3g",
			"ntp",
			"openbsd-netcat",
			"openssh",
			"os-prober",
			"p7zip",
			"papirus-icon-theme",
			"patch",
			"picom",
			"pkgconf",
			"powerline-fonts",
			"pulseaudio",
			"pulseaudio-alsa",
			"pulseaudio-bluetooth",
			"python-notify2",
			"python-pip",
			"python-psutil",
			"python-pyqt5",
			"qemu",
			"snap-pac",
			"snapper",
			"steam",
			"swtpm",
			"synergy",
			"terminus-font",
			"traceroute",
			"ttf-droid",
			"ttf-hack",
			"ttf-roboto",
			"ufw",
			"unrar",
			"unzip",
			"usbutils",
			"virt-manager",
			"virt-viewer",
			"which",
			"wine-gecko",
			"wine-mono",
			"winetricks",
			"xdg-user-dirs",
			"xorg",
			"xorg-apps",
			"xorg-drivers",
			"xorg-server",
			"xorg-xinit",
			"xorg-xkill",
			"xterm",
			"zip",
			"zsh",
			"zsh-autosuggestions",
			"zsh-syntax-highlighting",
		)
	}

	SaveConfig(c)
}

func AUR(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
		// Nothing yet
		// c.Pacman.AUR.Packages = append(c.Pacman.AUR.Packages,
		// )
	case "FULL":
		c.Pacman.AUR.Packages = append(c.Pacman.AUR.Packages,
			"bcompare",                    // Beyond Compare 4
			//"cockpit-navigator",         // File Browser for Cockpit
			"enpass-bin",                  // A multiplatform password manager
			"firefox-pwa-bin",             // Install, manage and use PWAs in Firefox (native component)
			"macchina-bin",                // A system information fetcher/frontend, with an (unhealthy) emphasis on performance
			//"morgen-bin",                // Modern, intuitive and smart calendar application (Successor to MineTime)
			"openrazer-meta",              // Razer device drivers
			"polychromatic",               // RGB management GUI for Razer Devices
			"razer-nari-pipewire-profile", // Razer Nari headsets pipewire profile
			"sddm-sugar-dark",             // "Sweetest" Dark theme for SDDM
			//"synology-drive",            // Powerful private cloud storage with no recurring fees
			"system-monitoring-center",    // Windows Task Manager Clone
		)

		switch c.Storage.Filesystem {
		case "zfs":
			c.Pacman.AUR.Packages = append(c.Pacman.AUR.Packages,
				"cockpit-zfs-manager", // ZFS on Linux admin package for Cockpit
			)
		}
	}

	SaveConfig(c)
}

func Flatpak(c *Config) {

	switch c.Desktop.InstallType {
	case "MINIMAL":
	case "FULL":
		c.Flatpak.Packages = append(c.Flatpak.Packages,
			"com.belmoussaoui.Obfuscate",        // Censor private information from images
			//"com.brave.Browser",               // Web browser from Brave
			"com.microsoft.Edge",                // Microsoft Edge Chromium Browser
			"com.getmailspring.Mailspring",      // Email client
			"com.github.alainm23.planner",       // Never worry about forgetting things again
			"com.github.liferooter.textpieces",  // Transform text without random websites
			"com.github.tchx84.Flatseal",        // Manage Flatpak permissions
			"com.gitlab.bitseater.meteo",        // Weather forecast app
			"com.github.wwmm.easyeffects",       // Audio Effects for PipeWire Applications
			//"com.obsproject.Studio",           // Live streaming and video recording software
			"com.steamgriddb.steam-rom-manager", // Manage ROMs in Steam
			"com.snes9x.Snes9x",                 // Super Nintendo Emulator
			"com.stepmania.StepMania",           // Rhythm and Dance Game (DDR for PC)
			"com.synology.SynologyDrive",        // Powerful private cloud storage with no recurring fees
			"com.usebottles.bottles",            // Sandox Windows Applications in Linux
			"im.riot.Riot",                      // Element -- Matrix Chat Client
			"io.github.antimicrox.antimicrox",   // Map gamepad buttons
			"io.github.peazip.PeaZip",           // File Archiver Utility (RAR, TAR, ZIP)
			"io.github.simple64.simple64",       // Nintendo 64 Emulator
			"io.github.seadve.Kooha",            // Record your screen
			"io.mgba.mGBA",                      // Nintendo Game Boy Advance Emulator
			"net.davidotek.pupgui2",             // ProtonUp-QT (Install Wine and Proton-based compatibility tools)
			"net.filebot.Filebot",               // Ultimate TV and Movie Renamer
			"net.kuribo64.melonDS",              // Nintendo DS and DSi emulator
			"net.pcsx2.PCSX2",                   // Playstation 2 emulator
			"org.citra_emu.citra",               // Nintendo 3DS emulator
			"org.DolphinEmu.dolphin-emu",        // Nintendo GameCube and Wii emulator
			"org.gnome.Firmware",                // Install firmware on devices
			//"org.gnome.Maps",                  // Find places around the world
			"org.libreoffice.LibreOffice",       // LibreOffice productivity suite
			"org.libretro.RetroArch",            // Frontend for emulators, game engines and media players
			"org.mixxx.Mixxx",                   // DJ Mixing software (Traktor alternative)
			"org.mozilla.Thunderbird",           // Email, RSS, and newsgroup client with integrated spam filter
			//"org.onlyoffice.desktopeditors",   // OnlyOffice productivity suite
			"org.ppsspp.PPSSPP",                 // PlayStation Portable emulator
			"org.yuzu_emu.yuzu",                 // Nintendo Switch emulator
			"tv.plex.PlexDesktop",               // Plex Client for desktop computers
		)
	}

	SaveConfig(c)
}

func PWA(c *Config) {

	youtube := Site{
		Name:        "YouTube",
		Description: "Share your videos with friends, family, and the world.",
		StartUrl:    "https://www.youtube.com",
		Manifest:    "manifest.webmanifest",
	}

	c.PWA.Sites = append(c.PWA.Sites, youtube)

	SaveConfig(c)
}
