# ShobuArch -- Automated Arch Linux Tools (Written in Go)
[![Lint Status](https://github.com/ShobuPrime/ShobuArch/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/marketplace/actions/run-golangci-lint)
[![Build Status](https://github.com/ShobuPrime/ShobuArch/actions/workflows/go-build.yml/badge.svg)](https://github.com/ShobuPrime/ShobuArch/actions/workflows/go-build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShobuPrime/ShobuArch)](https://goreportcard.com/report/github.com/ShobuPrime/ShobuArch)
[![License](https://img.shields.io/badge/License-GPLv3-success.svg)](./LICENSE.md)

<img src="https://blog.cloudflare.com/content/images/2018/02/gopher-tux-1.png" />

Have you ever wanted to use an IaC (Infrastructure as Code) approach towards automating an Arch Linux environment? If so, this README should help with getting you get what you need.

---
### Why does this exist?
There are other Automated Linux installers, and there's nothing wrong with them if your needs are met. However, if you're a ZFS fan most alternatives don't account for ZFS on Root, let alone a mirrored configuration!

On a personal level, I'm finally making efforts to run Linux as my daily driver since I believe we are approaching the time of the Linux Desktop. Most tools I enjoy to use these days are written in Go, and I wanted to become a Gopher myself.

In addition, the release of the Steam Deck has motivated me to ensure I get Arch Linux configured exactly the way I want.

Implementing a project like this was practical for me to easily and consistently repeat multiple OS installs while ensuring I improve my skills in another language.

---
## Create Arch ISO or Use Image
Download ArchISO from <https://archlinux.org/download/> and put on a USB drive with [Etcher](https://www.balena.io/etcher/), [Ventoy](https://www.ventoy.net/en/index.html), or [Rufus](https://rufus.ie/en/)

If you're looking for ZFS on Root, ensure to use an ISO with ZFS packages pre-installed.

Some archiso wrappers include: <https://github.com/ShobuPrime/arch-iso-zfs> and <https://github.com/stevleibelt/arch-linux-live-cd-iso-with-zfs>

Another helpful guide can be found at <https://michaelabrahamsen.com/posts/arch-linux-iso-zfs/>

## Boot Arch ISO

If running ShobuArch from official repository, run the following commands on a fresh boot of your ISO:

```
pacman -Syy git
git clone https://github.com/ShobuPrime/ShobuArch.git
cd ShobuArch/.build
chmod +x ShobuArch
./ShobuArch
```

ShobuArch runs both with and without arguments. Currently, supported flags are:

```
  -config string
        'y': Load config, 'n': Fresh config (default "n")
  -format string
        Accepted: 'JSON' || 'YAML' (default "json")
  -method string
        'a': Automated, 'm': Manual (default "m")
```

### System Description
Depending on arguments used, ShobuArch prompts the user for information deemed important for a coherent OS install:
- Username
- Password
- Filesystem
- Kernel
- Desktop Environment
- Window Manager
- AUR helper
- etc.

## Troubleshooting

__[Official Arch Installation Guide](https://wiki.archlinux.org/title/Installation_guide)__

__[Install Ach Linux on ZFS](https://wiki.archlinux.org/title/Install_Arch_Linux_on_ZFS)__

__[Arch Linux Installed on ZFS](https://michaelabrahamsen.com/posts/arch-linux-installed-on-zfs/)__

ShobuArch generates a logfile including every command ran, and its STDOUT and STDERR outputs. If something ended up not working the way you want, you should be able to dig through the logs to find what commands started to Go wonky.

# To-Do
- Add some GIFS or screenshots of the tool working in action

# Planned Features
- Offline ISO builder
- Verbosity flag, since right now the log is naturally very verbose
- Implement some sort of `resume` mechanism if a specific function failed very poorly, so we can continue from where we left off

# Known Issues
- When generating a fresh config through the UI, [util-linux v2.38](https://github.com/util-linux/util-linux/releases/tag/v2.38) added new key-value pairs which are being returned incorrectly. [A fix was committed](https://github.com/util-linux/util-linux/commit/97ce458a0dc2465fca2b36a6a4aa3378b607476e).
- You can workaround this issue my manually creating your config and loading it. A temp fix was also implemented via [this commit](https://github.com/ShobuPrime/ShobuArch/commit/5811003470471982ddcca68f726ea1d5584e4dc7) to only return "working" fields.

## Credits
 
- Inspiration for the project came from the philosophy of [ArchTitus](https://github.com/ChrisTitusTech/ArchTitus)
- Thank you Reddit User /u/BrenekH for [helping solve my Go Report Card issues](https://www.reddit.com/r/golang/comments/u0kg78/shobuarch_automated_arch_linux_tools_written_in_go/i46pgnq?utm_medium=android_app&utm_source=share&context=3)
- Thank you Reddit User /u/AladW for [identifying an issue with hostnames](https://www.reddit.com/r/archlinux/comments/u0kh1a/shobuarch_automated_arch_linux_tools_written_in_go/i46kvsk?utm_medium=android_app&utm_source=share&context=3)
- Thank you Reddit user /u/Max_yask for [suggesting significant feature requests such as YAML and sourcing dotfiles](https://www.reddit.com/r/linuxadmin/comments/u0ki4j/comment/i49qxag/?utm_source=share&utm_medium=web2x&context=3)
