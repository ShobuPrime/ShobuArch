#!/bin/bash

# Install Arch User Repository helper (yay)
# git clone https://aur.archlinux.org/yay.git /opt/yay
# chown -R username:users /opt/yay

# Make yay package, and return to PWD
current_dir=$PWD
cd /opt/yay
makepkg -si -S --noconfirm
cd $current_dir

# Install AUR Packages
# yay -Syyu --noconfirm \
#     vulkan-amdgpu-pro \
#     amf-amdgpu-pro \
#     amdgpu-pro-libgl \
#     openrazer-meta \
#     polychromatic \
#     microsoft-edge-stable-bin \
#     cockpit-navigator \
#     cockpit-zfs-manager

yay -Syyu --noconfirm openrazer-driver-dkms polychromatic microsoft-edge-stable-bin cockpit-navigator

# Configure Razer Drivers
gpasswd -a shobuprime plugdev
systemctl enable openrazer-daemon.service

# Install zsh
pacman -Syyu --needed --noconfirm \
    zsh \
    grml-zsh-config

chsh --shell /bin/zsh