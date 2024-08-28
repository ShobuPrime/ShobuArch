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
		"apparmor",               // Mandatory Access Control (MAC) using Linux Security Module (LSM)
		"arch-wiki-docs",         // Arch Wiki Documentation (Offline)
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
		"freerdp",                // Free implementation of the Remote Desktop Protocol (RDP)
		"fwupd",                  // Daemon to update some devices' firmware
		"go",                     // Core compiler tools for the Go programming language
		"gnu-netcat",             // GNU rewrite of netcat, the network piping application (QEMU)
		"iptables-nft",           // Linux kernel packet control tool (using nft interface)
		"jq",                     // Command-line JSON processor
		//"keepassxc",            // Cross-platform community-driven port of Keepass password manager
		"libdbusmenu-glib",       // Library for passing menus over DBus
		"lib32-pipewire",         // Low-latency audio/video router and processor -- 32bit
		"lib32-pipewire-jack",    // Pipewire JACK support -- 32bit
		"lib32-pipewire-v4l2",    // Pipewire V4L2 interceptor -- 32bit
		"networkmanager",         // Network connection manager and user applications
		"network-manager-applet", // Applet for managing network connections
		"noto-fonts-emoji",       // Google Noto emoji fonts
		"openssh",                // Premier connectivity tool for remote login with the SSH protocol
		"os-prober",              // Utility to detect other OSes on a set of drives
		"packagekit",             // A system designed to make installation and updates of packages easier
		"pipewire",               // Low-latency audio/video router and processor
		"pipewire-docs",          // Pipewire Documentation
		"pipewire-jack",          // Pipewire JACK support
		"pipewire-pulse",         // Pipewire PulseAudio replacement
		"pipewire-v4l2",          // Pipewire V4L2 interceptor
		"python",                 // Next generation of the python high-level scripting language
		"python-black",           // Uncompromising Python code formatter
		"python-pip",             // The PyPA recommended tool for installing Python packages
		"python-pipx",            // Install and Run Python Applications in Isolated Environments
		"qemu-full",              // A generic and open source machine emulator and virtualizer
		"reflector",              // A Python 3 module and script to retrieve and filter the latest Pacman mirror list.
		"rsync",                  // A fast and versatile file copying tool for remote and local files
		"steam",                  // Valve's digital software delivery system
		"system-config-printer",  // A CUPS printer configuration tool and status applet
		"terminus-font",          // Monospace bitmap font (for X11 and console)
		"virt-manager",           // Desktop user interface for managing virtual machines
		"v4l2loopback-dkms",      // Virtual Camera for OBS-Studio
		"wget",                   // Network utility to retrieve files from the
		"wireplumber",            // Session & policy manager implementation for PipeWire
		"wpa_supplicant",         // A utility providing key negotiation for WPA wireless networks
		//"x11-ssh-askpass",      // Lightweight passphrase dialog for SSH (QEMU)
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
			"kdeplasma-addons",            // All kind of addons to improve your Plasma experience
			"kdialog",                     // Display dialog boxes from shell scripts
			"kdf",                         // View Disk Usage
			"keditbookmarks",              // Bookmark Organizer and Editor
			"kfind",                       // Find files/folders
			"kgamma",                      // Adjust monitor gamma settings
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
			"maliit-keyboard",             // Virtual keyboard based on Maliit framework
			"markdownpart",                // KPart for rendering Markdown content
			"merkuro",                     // A calendar application using Akonadi to sync with external services
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
			"plasma-workspace",            // KDE Plasma Workspace
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
	case "FLATPAK":
		c.Flatpak.Packages = append(c.Flatpak.Packages,
			"org.kde.ark",
			"org.kde.dolphin",
			"org.kde.filelight",
			"org.kde.gwenview",
			"org.kde.kile",
			"org.kde.okular",
			"org.kde.kdenlive",
			"org.kde.koko",
			"org.kde.spectacle",
			"org.kde.okteta",
			"org.kde.kasts",
			"org.kde.kmplayer",
			"org.kde.haruna",
			"org.kde.neochat",
			"org.kde.tokodon",
			"org.kde.kxstitch",
			"org.kde.skrooge",
			"org.kde.itinerary",
			"org.kde.skanpage",
			"org.kde.kontact",
			"org.kde.konsole",
			"org.kde.akregator",
			"org.kde.arianna",
			"org.kde.kopete",
			"org.kde.elisa",
			"org.kde.kdevelop",
			"org.kde.lokalize",
			"org.kde.umbrello",
			"org.kde.ktechlab",
			"org.kde.peruse",
			"org.kde.ktorrent",
			"org.kde.kget",
			"org.kde.kgeotag",
			"org.kde.massif-visualizer",
			"org.kde.tellico",
			"org.kde.krdc",
			"org.kde.ksirk",
			"org.kde.konversation",
			"org.kde.gcompris",
			"org.kde.kpat",
			"org.kde.kbibtex",
			"org.kde.labplot2",
			"org.kde.cantor",
			"org.kde.digikam",
			"org.kde.kphotoalbum",
			"org.kde.krita",
			"org.kde.plasmatube",
			"org.kde.amarok",
			"org.kde.trojita",
			"org.kde.rocs",
			"org.kde.kleopatra",
			"org.kde.kdiff3",
			"org.kde.kgraphviewer",
			"org.kde.qmlkonsole",
			"org.kde.kjumpingcube",
			"org.kde.klickety",
			"org.kde.klettres",
			"org.kde.yakuake",
			"org.kde.kimagemapeditor",
			"org.kde.kweather",
			"org.kde.ktrip",
			"org.kde.kongress",
			"org.kde.keysmith",
			"org.kde.pix",
			"org.kde.picmi",
			"org.kde.parley",
			"org.kde.palapeli",
			"org.kde.minuet",
			"org.kde.marble",
			"org.kde.kwrite",
			"org.kde.kwordquiz",
			"org.kde.kwalletmanager5",
			"org.kde.kuiviewer",
			"org.kde.kubrick",
			"org.kde.kturtle",
			"org.kde.ktuberling",
			"org.kde.artikulate",
			"org.kde.atlantik",
			"org.kde.audiotube",
			"org.kde.blinken",
			"org.kde.bomber",
			"org.kde.bovo",
			"org.kde.alligator",
			"org.kde.angelfish",
			"org.kde.falkon",
			"org.kde.granatier",
			"org.kde.juk",
			"org.kde.kalgebra",
			"org.kde.kalzium",
			"org.kde.kamoso",
			"org.kde.kate",
			"org.kde.kblocks",
			"org.kde.kigo",
			"org.kde.killbots",
			"org.kde.kanagram",
			"org.kde.kapman",
			"org.kde.katomic",
			"org.kde.kblackbox",
			"org.kde.kreversi",
			"org.kde.kruler",
			"org.kde.ksnakeduel",
			"org.kde.kdf",
			"org.kde.kcolorchooser",
			"org.kde.kcalc",
			"org.kde.kcachegrind",
			"org.kde.kbreakout",
			"org.kde.kbruch",
			"org.kde.kfourinline",
			"org.kde.kfind",
			"org.kde.kdiamond",
			"org.kde.kgeography",
			"org.kde.khangman",
			"org.kde.kgoldrunner",
			"org.kde.kig",
			"org.kde.kiriki",
			"org.kde.kiten",
			"org.kde.kmahjongg",
			"org.kde.klines",
			"org.kde.kmines",
			"org.kde.vvave",
			"org.kde.krename",
			"org.kde.kmymoney",
			"org.kde.kontrast",
			"org.kde.kmplot",
			"org.kde.knetwalk",
			"org.kde.konquest",
			"org.kde.kompare",
			"org.kde.kolourpaint",
			"org.kde.kollision",
			"org.kde.kolf",
			"org.kde.knights",
			"org.kde.knavalbattle",
			"org.kde.ktouch",
			"org.kde.kteatime",
			"org.kde.ksudoku",
			"org.kde.ksquares",
			"org.kde.lskat",
			"org.kde.kclock",
			"org.kde.kalk",
			"org.kde.index",
			"org.kde.nota",
			"org.kde.kommit",
			"org.kde.subtitlecomposer",
			"org.kde.isoimagewriter",
			"org.kde.PlatformTheme.QGnomePlatform",
			"org.kde.kronometer",
			"org.kde.skanlite",
			"org.kde.krusader",
			"org.kde.francis",
			"org.kde.WaylandDecoration.QAdwaitaDecorations",
			"org.kde.kid3",
			"org.kde.ruqola",
			"org.kde.KStyle.Kvantum",
			"org.kde.telly-skout",
			"org.kde.kst",
			"org.kde.KStyle.Adwaita",
			"org.kde.KStyle.HighContrast",
			"org.kde.kstars",
			"org.kde.ktimetracker",
			"org.kde.SymbolEditor",
			"org.kde.kaffeine",
			"org.kde.Ikona",
			"org.kde.choqok",
			"org.kde.WaylandDecoration.QGnomePlatform-decoration",
			"org.kde.PlatformInputContexts.MaliitSailfishOS",
			"org.kde.PlatformTheme.QtSNI",
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
			"cockpit-navigator-git",       // File Browser for Cockpit
			"enpass-bin",                  // A multiplatform password manager
			"firefox-pwa-bin",             // Install, manage and use PWAs in Firefox (native component)
			"macchina-bin",                // A system information fetcher/frontend, with an (unhealthy) emphasis on performance
			//"mcmojave-cursors",          // X-cursor theme inspired by macOS and based on capitaine-cursors (Currently refuses to install via arch-chroot)
			//"morgen-bin",                // Modern, intuitive and smart calendar application (Successor to MineTime)
			//"nerd-dictation-git",        // Simple, hackable offline speech to text - using the VOSK-API.
			//"sddm-sugar-dark",           // "Sweetest" Dark theme for SDDM
			// "ttf-ms-win11-auto",        // Microsoft Windows 11 TrueType fonts (Currently refuses to install via arch-chroot)
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
			"com.belmoussaoui.Obfuscate",          // Censor private information from images
			"com.brave.Browser",                   // Web browser from Brave
			"com.getmailspring.Mailspring",        // Email client
			"com.github.alainm23.planner",         // Never worry about forgetting things again
			"com.github.jeromerobert.pdfarranger", // PDF Merging, Rearranging, Splitting, Rotating and Cropping
			"com.github.liferooter.textpieces",    // Transform text without random websites
			"com.github.Matoking.protontricks",    // Winetricks for Proton enabled games
			"io.github.prateekmedia.appimagepool", // Simple, modern AppImageHub client
			"com.gitlab.bitseater.meteo",          // Weather forecast app
			"com.gitlab.davem.ClamTk",             // GUI for Clam Antivirus (ClamAV)
			"com.github.tchx84.Flatseal",          // Manage Flatpak permissions
			"com.github.wwmm.easyeffects",         // Audio Effects for PipeWire Applications
			"com.leinardi.gst",                    // System utility designed to stress and monitor various hardware components
			"com.obsproject.Studio",               // Live streaming and video recording software
			"com.protonvpn.www",                   // High-speed Swiss VPN that safeguards your privacy
			"com.steamgriddb.steam-rom-manager",   // Manage ROMs in Steam
			"com.snes9x.Snes9x",                   // Super Nintendo Emulator
			"com.stepmania.StepMania",             // Rhythm and Dance Game (DDR for PC)
			"com.synology.SynologyDrive",          // Powerful private cloud storage with no recurring fees
			"com.usebottles.bottles",              // Sandox Windows Applications in Linux
			//"com.visualstudio.code-oss",         // Visual Studio Code. Code editing. Redefined.
			"im.riot.Riot",                        // Element -- Matrix Chat Client
			"in.srev.guiscrcpy",                   // Android Screen Mirroring Software
			"io.github.antimicrox.antimicrox",     // Map gamepad buttons
			"io.github.peazip.PeaZip",             // File Archiver Utility (RAR, TAR, ZIP)
			"io.github.simple64.simple64",         // Nintendo 64 Emulator
			"io.github.seadve.Kooha",              // Record your screen
			"io.mgba.mGBA",                        // Nintendo Game Boy Advance Emulator
			"io.missioncenter.MissionCenter",      // Monitor system resource usage
			"it.cuteworks.pacmanlogviewer",        // Inspect pacman logs
			"md.obsidian.Obsidian",                // Markdown-based knowledge base
			"media.emby.EmbyTheater",              // Emby (Media Server) client for desktop computers
			"net.davidotek.pupgui2",               // ProtonUp-QT (Install Wine and Proton-based compatibility tools)
			"net.filebot.FileBot",                 // Ultimate TV and Movie Renamer
			"net.kuribo64.melonDS",                // Nintendo DS and DSi emulator
			"net.pcsx2.PCSX2",                     // Playstation 2 emulator
			"org.citra_emu.citra",                 // Nintendo 3DS emulator
			"org.DolphinEmu.dolphin-emu",          // Nintendo GameCube and Wii emulator
			"org.gimp.GIMP",                       // Create images and edit photographs
			"org.gnome.Firmware",                  // Install firmware on devices
			//"org.gnome.Maps",                    // Find places around the world
			"org.kde.kdenlive",                    // Non-linear video editor for Linux
			//"org.keepassxc.KeePassXC",           // Cross-platform community-driven port of Keepass password manager
			"org.libreoffice.LibreOffice",         // LibreOffice productivity suite
			"org.libretro.RetroArch",              // Frontend for emulators, game engines and media players
			"org.mixxx.Mixxx",                     // DJ Mixing software (Traktor alternative)
			"org.mozilla.Thunderbird",             // Email, RSS, and newsgroup client with integrated spam filter
			//"org.onlyoffice.desktopeditors",     // OnlyOffice productivity suite
			"org.ppsspp.PPSSPP",                   // PlayStation Portable emulator
			"org.rncbc.qpwgraph",                  // PipeWire Graph Qt GUI Interface
			"org.upscayl.Upscayl",                 // Free and Open Source AI Image Upscaler
			"org.yuzu_emu.yuzu",                   // Nintendo Switch emulator
			"tv.plex.PlexDesktop",                 // Plex client for desktop computers
			"tv.plex.PlexHTPC",                    // Plex HTPC client for the big screen
		)
	}

	SaveConfig(c)
}

func PWA(c *Config) {

	messages := Site{
		Name:        "Messages",
		Description: "Simple, helpful messaging by Google",
		StartUrl:    "https://messages.google.com/web",
		Manifest:    "manifest.json",
	}

	c.PWA.Sites = append(c.PWA.Sites, messages)

	youtube := Site{
		Name:        "YouTube",
		Description: "Share your videos with friends, family, and the world.",
		StartUrl:    "https://www.youtube.com",
		Manifest:    "manifest.webmanifest",
	}

	c.PWA.Sites = append(c.PWA.Sites, youtube)

	SaveConfig(c)
}

// Make a note to update /usr/bin/steam-runtime to include `env LD_PRELOAD=/usr/lib32/libextest.so` between `exec` and `/usr/lib/steam/steam "$@"`
