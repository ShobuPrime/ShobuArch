#!/bin/bash

# Comment out all zroot entries, because zfs handles this by itself
sed -i 's/zroot/# zroot/g' /etc/fstab
sed -i 's/# # zroot/# zroot/g' /etc/fstab

# Add parallel downloading
# sed -i 's/^#ParallelDownloads/ParallelDownloads/' /etc/pacman.conf

# Enable multilib repos
sed -i "/\[multilib\]/,/Include/"'s/^#//' /etc/pacman.conf

# plasma-desktop and plasma-nm for more minimul install
pacman -Syu --needed --noconfirm \
    efibootmgr \
    grub \
    cockpit \
    cockpit-machines \
    packagekit \
    docker \
    go \
    virt-manager \
    qemu \
    vde2 \
    dnsmasq \
    dmidecode \
    code \
    libdbusmenu-glib \
    networkmanager \
    network-manager-applet \
    openssh \
    firewalld \
    os-prober \
    reflector \
    rsync \
    terminus-font \
    wpa_supplicant \
    xdg-user-dirs \
    xdg-utils \
    wget \
    firefox \
    plasma \
    kde-applications \
    plasma-wayland-session \
    sddm \
    lib32-mesa \
    vulkan-radeon \
    vulkan-tools \
    lib32-vulkan-radeon \
    vulkan-icd-loader \
    lib32-vulkan-icd-loader \
    steam

# Configure kvm
sed -i 's/#unix_sock_group/unix_sock_group/g' /etc/libvirt/libvirtd.conf
sed -i 's/#unix_sock_rw_perms/unix_sock_rw_perms/g' /etc/libvirt/libvirtd.conf

# Add custom repo to install ArchZFS
wget https://archzfs.com/archzfs.gpg
echo "" >> /etc/pacman.conf
echo '[archzfs]'  >> /etc/pacman.conf
pacman-key -a archzfs.gpg
pacman-key -r DDF7DB817396A49B2A2723F7403BD972F75D9D76
pacman-key --lsign-key DDF7DB817396A49B2A2723F7403BD972F75D9D76

# Check the fingerprint and verify it matches the one on the archzfs page
pacman-key -f DDF7DB817396A49B2A2723F7403BD972F75D9D76

# echo "SigLevel = Optional TrustAll"  >> /etc/pacman.conf
echo "# Origin Server - France" >> /etc/pacman.conf
echo 'Server = http://archzfs.com/$repo/x86_64' >> /etc/pacman.conf

echo "# Mirror - Germany" >> /etc/pacman.conf
echo 'Server = http://mirror.sum7.eu/archlinux/archzfs/$repo/x86_64' >> /etc/pacman.conf

echo "# Mirror - Germany" >> /etc/pacman.conf
echo 'Server = https://mirror.biocrafting.net/archlinux/archzfs/$repo/x86_64' >> /etc/pacman.conf

echo "# Mirror - India" >> /etc/pacman.conf
echo 'Server = https://mirror.in.themindsmaze.com/archzfs/$repo/x86_64' >> /etc/pacman.conf

echo "# ArchZFS - US Mirror" >> /etc/pacman.conf
echo 'Server = https://zxcvfdsa.com/archzfs/$repo/$arch' >> /etc/pacman.conf

pacman -Syu --needed --noconfirm \
    zfs-dkms \
    zfs-utils

systemctl enable cockpit.socket
systemctl enable docker.service
systemctl enable firewalld
systemctl enable libvirtd.service
systemctl enable NetworkManager
systemctl enable reflector.timer
systemctl enable sddm
systemctl enable sshd
systemctl enable zfs-import-cache
systemctl enable zfs-import-scan
systemctl enable zfs-mount
systemctl enable zfs-share
systemctl enable zfs-zed
systemctl enable zfs.target

# systemctl start firewalld
# firewall-cmd --permanent --add-service=cockpit

# Configure TimeZone
ln -sf /usr/share/zoneinfo/America/New_York /etc/localtime

# Sync to hardware clock
hwclock --systohc

# Configure locale
sed -i 's/#en_US.UTF-8/en_US.UTF-8/g' /etc/locale.gen && locale-gen
echo "LANG=en_US.UTF-8" >> /etc/locale.conf

# Configure Hostname
echo "<Hostname01>" >> /etc/hostname
echo "127.0.0.1 localhost" >> /etc/hosts
echo "::1       localhost" >> /etc/hosts
echo "127.0.1.1 <Hostname01>.local Hostname02" >> /etc/hosts

# Install zsh
pacman -Syu --needed --noconfirm \
    zsh \
    grml-zsh-config

# Create main user
useradd -mU \
    -s /bin/zsh \
    -G sys,log,network,floppy,scanner,power,rfkill,users,video,storage,optical,lp,audio,wheel,adm,docker,kvm \
    -d /home/shobuprime/ \
    username

# Set password
passwd shobuprime

# Give root password
passwd

# visudo -- Uncomment 'wheel' group
sed -i 's/# %wheel ALL=(ALL:ALL) ALL/%wheel ALL=(ALL:ALL) ALL/g' /etc/sudoers

# Add user to sudoers
# touch /etc/sudoers.d/local_users
# echo 'shobuprime ALL=(ALL) ALL' >> /etc/sudoers.d/local_users

# To prioritize the local and fastest repositories, edit reflector.conf to have the following
# --save /etc/pacman.d/mirrorlist
# --protocol https
# --country US
# --latest 5
# --sort rate
cp /etc/xdg/reflector/reflector.conf /etc/xdg/reflector/reflector.conf.bak
sed -i 's/#--country France,Germany/--country US/g' /etc/xdg/reflector/reflector.conf
sed -i 's/--sort age/--sort rate/g' /etc/xdg/reflector/reflector.conf

# # Install Arch User Repository helper (yay)
git clone https://aur.archlinux.org/yay.git /opt/yay
chown -R shobuprime:users /opt/yay

# Make yay package, and return to PWD
current_dir=$PWD
cd /opt/yay
su -c 'makepkg -si -S --noconfirm' shobuprime
pacman -U yay_11.1.1_x86_64.tar.gz
cd $current_dir

# # Install AUR Packages
# # yay -Syu --noconfirm \
# #     vulkan-amdgpu-pro \
# #     amf-amdgpu-pro \
# #     amdgpu-pro-libgl \
# #     openrazer-meta \
# #     polychromatic \
# #     microsoft-edge-stable-bin \
# #     cockpit-navigator \
# #     cockpit-zfs-manager

su -c 'yay -Syu --noconfirm openrazer-meta polychromatic microsoft-edge-stable-bin cockpit-navigator' shobuprime

# # Configure Razer Drivers
gpassws -a shobuprime plugdev
systemctl enable openrazer-daemon.service

# GRUB_CMDLINE_LINUX_DEFAULT: add "video=1920x1080"
# GRUB_CMDLINE_LINUX="root=ZFS=zroot/ROOT/default"
sed -i 's/quiet/quiet video=1920x1080/g' /etc/default/grub 
sed -i 's/GRUB_CMDLINE_LINUX=""/GRUB_CMDLINE_LINUX="root=ZFS=zroot\/ROOT\/default"/g' /etc/default/grub

grub-install --target=x86_64-efi --efi-directory=/boot --bootloader-id=ArchLinux

grub-mkconfig -o /boot/grub/grub.cfg

# Make HOOKS=(base udev autodetect modconf block keyboard zfs filesystems)
sed -i 's/HOOKS=(base udev autodetect modconf block filesystems keyboard fsck)/HOOKS=(base udev autodetect modconf block keyboard zfs filesystems shutdown)/g' /etc/mkinitcpio.conf

mkinitcpio -p linux
