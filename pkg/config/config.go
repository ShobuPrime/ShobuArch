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

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"

	u "github.com/ShobuPrime/ShobuArch/pkg/util"

	"github.com/manifoldco/promptui"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	Format   string   `json:"-" yaml:"-"`
	Kernel   string   `json:"kernel" yaml:"kernel"`
	Hostname string   `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Timezone string   `json:"timezone,omitempty" yaml:"timezone,omitempty"`
	User     User     `json:"user" yaml:"user"`
	Storage  Storage  `json:"storage" yaml:"storage"`
	Desktop  Desktop  `json:"desktop,omitempty" yaml:"desktop,omitempty"`
	Pacman   Pacman   `json:"pacman,omitempty" yaml:"pacman,omitempty"`
	Flatpak  Flatpaks `json:"flatpak,omitempty" yaml:"flatpak,omitempty"`
	PWA      PWAs     `json:"pwa,omitempty" yaml:"pwa,omitempty"`
}

type User struct {
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
	Language Language `json:"language" yaml:"language"`
	Shell    string   `json:"user_shell,omitempty" yaml:"shell,omitempty"`
	Git      Git      `json:"git,omitempty" yaml:"git,omitempty"`
}

type Language struct {
	Locale   string `json:"locale,omitempty" yaml:"locale,omitempty"`
	CharSet  string `json:"charset,omitempty" yaml:"charset,omitempty"`
	Keyboard string `json:"keyboard_layout,omitempty" yaml:"keyboard,omitempty"`
}

type Git struct {
	Username string `json:"git_username,omitempty" yaml:"git_username,omitempty"`
	Email    string `json:"git_email,omitempty" yaml:"git_password,omitempty"`
}

type Storage struct {
	SystemDisk     string `json:"system_disk" yaml:"system_disk"`
	SystemDiskID   string `json:"system_disk_id,omitempty" yaml:"system_disk_id,omitempty"`
	SystemDiskRota bool   `json:"system_disk_rota,omitempty" yaml:"sysem_disk_rota,omitempty"`
	MirrorInstall  bool   `json:"mirror_install" yaml:"mirror_install"`
	MirrorDisk     string `json:"mirror_disk,omitempty" yaml:"mirror_disk,omitempty"`
	MirrorDiskID   string `json:"mirror_disk_id,omitempty" yaml:"mirror_disk_id,omitempty"`
	MirrorDiskRota bool   `json:"mirror_disk_rota,omitempty" yaml:"mirror_disk_rota,omitempty"`
	Filesystem     string `json:"filesystem,omitempty" yaml:"filesystem,omitempty"`
	EncryptionKey  string `json:"encryption_key,omitempty" yaml:"encryption_key,omitempty"`
	EncryptionUUID string `json:"encryption_uuid,omitempty" yaml:"encryption_uuid,omitempty"`
}

type Desktop struct {
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`
	InstallType string `json:"install_type,omitempty" yaml:"install_type,omitempty"`
}

type Pacman struct {
	AUR      AURs     `json:"aur,omitempty" yaml:"aur,omitempty"`
	Packages []string `json:"packages,omitempty" yaml:"packages,omitempty"`
}

type AURs struct {
	Helper   string   `json:"aur_helper,omitempty" yaml:"aur_helper,omitempty"`
	Packages []string `json:"aur_packages,omitempty" yaml:"aur_packages,omitempty"`
}

type Flatpaks struct {
	Packages []string `json:"packages" yaml:"packages"`
}

type PWAs struct {
	Sites []Site `json:"sites" yaml:"sites"`
}

type Site struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	StartUrl    string `json:"start_url" yaml:"start_url"`
	Manifest    string `json:"manifest" yaml:"manifest"`
}

func Logo() {

	log.Println(`
	-------------------------------------------------------------------------
	███████╗██╗  ██╗ ██████╗ ██████╗ ██╗   ██╗ █████╗ ██████╗  ██████╗██╗  ██╗
	██╔════╝██║  ██║██╔═══██╗██╔══██╗██║   ██║██╔══██╗██╔══██╗██╔════╝██║  ██║
	███████╗███████║██║   ██║██████╔╝██║   ██║███████║██████╔╝██║     ███████║
	╚════██║██╔══██║██║   ██║██╔══██╗██║   ██║██╔══██║██╔══██╗██║     ██╔══██║
	███████║██║  ██║╚██████╔╝██████╔╝╚██████╔╝██║  ██║██║  ██║╚██████╗██║  ██║
	╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
	-------------------------------------------------------------------------
		Please select presetup settings for your system              
	-------------------------------------------------------------------------
	`)
}

func InitConfig() *Config {
	c := Config{}
	return &c
}

func SetKernel(c *Config) {
	prompt := promptui.Select{
		Label: "Please select your Linux kernel",
		Items: []string{"Stable", "Hardened", "Longterm", "Zen"},
		Size:  4,
	}

	_, kernel_choice, err := prompt.Run()

	if err != nil {
		log.Printf("Kernel selection failed %v\n", err)
		return
	}

	// https://wiki.archlinux.org/title/Kernel#Officially_supported_kernels
	switch kernel_choice {
	case "Stable":
		c.Kernel = "linux"
	case "Hardened":
		c.Kernel = "linux-hardened"
	case "Longterm":
		c.Kernel = "linux-lts"
	case "Zen":
		c.Kernel = "linux-zen"
	}

	log.Printf("Selected %q kernel -- %q\n", kernel_choice, c.Kernel)

	SaveConfig(c)
}

func SetUserInfo(c *Config) {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your username: ")
	username, _ := input.ReadString('\n')
	username = strings.TrimSuffix(username, "\n")
	log.Printf("Your username is: %q\n", username)
	c.User.Username = strings.ToLower(username)
	c.User.Password = SetPassword("user")
	SaveConfig(c)
}

func SetPassword(target string) string {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter " + target + " password: ")
	password, _ := input.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	log.Printf("Hash to store: %v\n", string(hash))
	log.Println("Store this 'hash' somewhere safe")

	// After a while, the user wants to log in and you need to check the password he entered
	input = bufio.NewReader(os.Stdin)
	fmt.Print("Please confirm " + target + " password: ")
	confirmed_pass, _ := input.ReadString('\n')
	confirmed_pass = strings.TrimSuffix(confirmed_pass, "\n")
	hashFromDatabase := hash

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(confirmed_pass)); err != nil {
		// TODO: Properly handle error
		log.Println("Encryption Keys do not match")
		log.Fatal(err)
	} else {
		log.Println("Encryption Keys match!")
		return password
	}

	// Placeholder statement
	return ""
}

func SetHostInfo(c *Config) {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your hostname: ")
	hostname, _ := input.ReadString('\n')
	hostname = strings.TrimSuffix(hostname, "\n")
	log.Printf("Your username is: %q\n", hostname)
	c.Hostname = hostname

	SaveConfig(c)
}

// To-do: Ask for git info
func SetGitInfo() {

}

func SelectAUR(c *Config) {
	prompt := promptui.Select{
		Label: "Please select desired AUR helper",
		Items: []string{"yay", "aura", "pacaur", "pakku", "paru", "pikaur", "trizen", "none"},
		Size:  7,
	}

	_, aur_choice, err := prompt.Run()

	if err != nil {
		log.Printf("AUR selection failed %v\n", err)
		return
	}

	c.Pacman.AUR.Helper = aur_choice

	AUR(c)
	SaveConfig(c)
}

func SelectDesktopEnv(c *Config) {
	prompt := promptui.Select{
		Label: "Please select your desired Desktop Environment",
		Items: []string{
			"kde", "gnome", "cinnamon",
			"xfce", "mate", "budgie",
			"lxde", "deepin", "openbox",
			"server",
		},
	}

	_, desktop_choice, err := prompt.Run()

	if err != nil {
		log.Printf("Desktop selection failed %v\n", err)
		return
	}

	c.Desktop.Environment = desktop_choice
	log.Printf("Your desktop environment is: %q\n", c.Desktop.Environment)

	switch c.Desktop.Environment {
	case "gnome":
		Gnome(c)
	case "kde":
		KDE(c)
	case "cinnamon":
		Cinnamon(c)
	case "xfce":
		XFCE(c)
	case "mate":
		Mate(c)
	case "budgie":
		Budgie(c)
	case "lxde":
		LXDE(c)
	case "deepin":
		Deepin(c)
	case "openbox":
		Openbox(c)
	case "server":
		Server(c)
	}

	SaveConfig(c)
}

func SetInstallType(c *Config) {

	prompt := promptui.Select{
		Label: "Please select Desktop Install Type",
		Items: []string{"FULL", "MINIMAL"},
	}

	_, choice, err := prompt.Run()

	if err != nil {
		log.Printf("Desktop Install Type selection failed %v\n", err)
		return
	}

	log.Printf("Your install method is: %q\n", choice)

	switch choice {
	case "FULL":
		c.Desktop.InstallType = choice
	case "MINIMAL":
		c.Desktop.InstallType = choice
	}
}

func SetFilesystem(c *Config) {
	prompt := promptui.Select{
		Label: "Select Filesystem",
		Items: []string{"btrfs", "ext4", "luks", "zfs"},
	}

	_, fs_choice, err := prompt.Run()

	if err != nil {
		fmt.Printf("Filesystem selection failed %v\n", err)
		return
	}

	c.Storage.Filesystem = fs_choice

	switch c.Storage.Filesystem {
	case "btrfs":
		SetMirror(c)
	case "luks":
		c.Storage.EncryptionKey = SetPassword("luks")
	case "zfs":
		SetMirror(c)
	}

	log.Printf("Filesystem: %q\n", c.Storage.Filesystem)
	SaveConfig(c)
}

func SetMirror(c *Config) {

	input := bufio.NewReader(os.Stdin)
	fmt.Printf("Would you like to configure a mirrored %s pool (Y/N): ", c.Storage.Filesystem)
	mirror_choice, _ := input.ReadString('\n')
	mirror_choice = strings.TrimSuffix(mirror_choice, "\n")
	switch strings.ToLower(mirror_choice) {
	case "y", "yes":
		log.Printf("%s mirror selected! Please select mirror disk.\n", c.Storage.Filesystem)
		c.Storage.MirrorInstall = true
		// Go version of "do while" -- Keep prompting user if mirrored disk matches root disk
		for same_disk := true; same_disk; same_disk = (c.Storage.MirrorDisk == c.Storage.SystemDisk) {
			c.Storage.MirrorDisk, c.Storage.MirrorDiskID, c.Storage.SystemDiskRota = SelectDisk()
			if c.Storage.MirrorDisk == c.Storage.SystemDisk {
				log.Println("WARNING: Same disk selected for mirror! Please select a different disk")
			}
			SaveConfig(c)
		}
	case "n", "no":
		log.Printf("%s will be installed on a single disk.\n", c.Storage.Filesystem)
		c.Storage.MirrorInstall = false
	}

	SaveConfig(c)
}

func SelectDisk() (string, string, bool) {

	lsblk := u.ListDisk()

	device_paths := []string{}
	device_ids := []string{}
	device_list := []string{}
	for i := range lsblk.Blockdevices {
		device_paths = append(device_paths, lsblk.Blockdevices[i].Path)
		lsblk.Blockdevices[i].Model = strings.ReplaceAll(lsblk.Blockdevices[i].Model, " ", "_")
		device_ids = append(device_ids, lsblk.Blockdevices[i].Model+"-"+lsblk.Blockdevices[i].Serial)
		lsblk.Blockdevices[i].Serial = strings.ReplaceAll(lsblk.Blockdevices[i].Serial, " ", "_")
		device_list = append(device_list, lsblk.Blockdevices[i].Path+" --> "+lsblk.Blockdevices[i].Model+"-"+lsblk.Blockdevices[i].Serial)
	}
	prompt := promptui.Select{
		Label: "Select Disk list method",
		Items: []string{"By Path", "By ID", "Combined"},
	}

	_, disk_list_method, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	log.Printf("Listing disks '%q'\n", disk_list_method)

	switch disk_list_method {
	case "By Path":
		prompt = promptui.Select{
			Label: "Select Disk",
			Items: device_paths,
		}
		i, choice, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		log.Printf("Selected disk: %v\n", choice)
		return device_paths[i], device_ids[i], lsblk.Blockdevices[i].Rota
	case "By ID":
		prompt = promptui.Select{
			Label: "Select Disk",
			Items: device_ids,
		}
		i, choice, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		log.Printf("Selected disk: %v\n", choice)
		return device_paths[i], device_ids[i], lsblk.Blockdevices[i].Rota
	case "Combined":
		prompt = promptui.Select{
			Label: "Select Disk",
			Items: device_list,
		}
		i, choice, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		log.Printf("Selected disk: %v\n", choice)
		return device_paths[i], device_ids[i], lsblk.Blockdevices[i].Rota
	}
	// These values are never expected to be returned.
	return "SystemDisk", "SystemDiskID", true
}

func SetTimezone(c *Config) {

	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter timezone (IANA Format -- ex: America/New_York): ")
	tz, _ := input.ReadString('\n')
	tz = strings.TrimSuffix(tz, "\n")
	switch tz {
	case "":
		tz = `America/New_York`
	}
	log.Printf("Your timezone is: %v\n", tz)
	c.Timezone = tz
	SaveConfig(c)
}

func SetKeyboardLayout(c *Config) {

	// These are default key maps as presented in official arch repo archinstall
	prompt := promptui.Select{
		Label: "What is your keyboard layout?",
		Items: []string{
			"us", "by", "ca", "cf", "cz",
			"de", "es", "et", "fa", "fi",
			"fr", "gr", "hu", "il", "it",
			"lt", "lv", "mk", "nl", "no",
			"pl", "ro", "ru", "sg", "ua",
			"uk",
		},
		Size: 13,
	}

	_, kb_layout, err := prompt.Run()

	if err != nil {
		log.Printf("Keyboard selection failed %v\n", err)
		return
	}

	c.User.Language.Keyboard = kb_layout
	SaveConfig(c)
}

func SetLocale(c *Config) {

	// Using UTF-8 as default for now since it's recommended from Arch's own wiki
	c.User.Language.CharSet = "UTF-8"

	prompt := promptui.Select{
		Label: "What is your locale?",
		Items: []string{
			fmt.Sprintf("aa_DJ.%s", c.User.Language.CharSet),
			"aa_ER", "aa_ET",
			fmt.Sprintf("af_ZA.%s", c.User.Language.CharSet),
			"agr_PE", "ak_GH", "am_ET",
			fmt.Sprintf("an_ES.%s", c.User.Language.CharSet),
			"anp_IN",
			fmt.Sprintf("ar_AE.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_BH.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_DZ.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_EG.%s", c.User.Language.CharSet),
			"ar_IN",
			fmt.Sprintf("ar_IQ.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_JO.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_KW.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_LB.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_LY.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_MA.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_OM.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_QA.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_SA.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_SD.%s", c.User.Language.CharSet),
			"ar_SS",
			fmt.Sprintf("ar_SY.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_TN.%s", c.User.Language.CharSet),
			fmt.Sprintf("ar_YE.%s", c.User.Language.CharSet),
			"ayc_PE", "az_AZ", "az_IR", "as_IN",
			fmt.Sprintf("ast_ES.%s", c.User.Language.CharSet),
			fmt.Sprintf("be_BY.%s", c.User.Language.CharSet),
			"bem_ZM", "ber_DZ", "ber_MA",
			fmt.Sprintf("bg_BG.%s", c.User.Language.CharSet),
			fmt.Sprintf("bhb_IN.%s", c.User.Language.CharSet),
			"bho_NP", "bi_VU", "bn_BD", "bn_IN", "bo_CN", "bo_IN",
			fmt.Sprintf("br_FR.%s", c.User.Language.CharSet),
			"brx_IN",
			fmt.Sprintf("bs_BA.%s", c.User.Language.CharSet),
			"byn_ER",
			fmt.Sprintf("C.%s", c.User.Language.CharSet),
			fmt.Sprintf("ca_AD.%s", c.User.Language.CharSet),
			fmt.Sprintf("ca_ES.%s", c.User.Language.CharSet),
			fmt.Sprintf("ca_FR.%s", c.User.Language.CharSet),
			fmt.Sprintf("ca_IT.%s", c.User.Language.CharSet),
			"ce_RU", "chr_US", "ckb_IQ", "cmn_TW", "crh_UA",
			fmt.Sprintf("cs_CZ.%s", c.User.Language.CharSet),
			"csb_PL", "cv_RU",
			fmt.Sprintf("cy_GB.%s", c.User.Language.CharSet),
			fmt.Sprintf("da_DK.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_AT.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_BE.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_CH.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_DE.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_IT.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_LI.%s", c.User.Language.CharSet),
			fmt.Sprintf("de_LU.%s", c.User.Language.CharSet),
			"doi_IN", "dsb_DE", "dv_MV", "dz_BT",
			fmt.Sprintf("el_GR.%s", c.User.Language.CharSet),
			fmt.Sprintf("el_CY.%s", c.User.Language.CharSet),
			"en_AG",
			fmt.Sprintf("en_AU.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_BW.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_CA.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_DK.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_GB.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_HK.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_IE.%s", c.User.Language.CharSet),
			"en_IL", "en_IN", "en_NG",
			fmt.Sprintf("en_NZ.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_PH.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_SC.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_SG.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_US.%s", c.User.Language.CharSet),
			fmt.Sprintf("en_ZA.%s", c.User.Language.CharSet),
			"en_ZM",
			fmt.Sprintf("en_ZW.%s", c.User.Language.CharSet),
			"eo",
			fmt.Sprintf("es_AR.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_BO.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_CL.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_CO.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_CR.%s", c.User.Language.CharSet),
			"es_CU",
			fmt.Sprintf("es_DO.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_EC.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_ES.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_GT.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_HN.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_MX.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_NI.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_PA.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_PE.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_PR.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_PY.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_SV.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_US.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_UY.%s", c.User.Language.CharSet),
			fmt.Sprintf("es_VE.%s", c.User.Language.CharSet),
			fmt.Sprintf("et_EE.%s", c.User.Language.CharSet),
			fmt.Sprintf("eu_US.%s", c.User.Language.CharSet),
			"fa_IR", "ff_SN",
			fmt.Sprintf("fi_FI.%s", c.User.Language.CharSet),
			"fil_PH",
			fmt.Sprintf("fo_FO.%s", c.User.Language.CharSet),
			fmt.Sprintf("fr_BE.%s", c.User.Language.CharSet),
			fmt.Sprintf("fr_CA.%s", c.User.Language.CharSet),
			fmt.Sprintf("fr_CH.%s", c.User.Language.CharSet),
			fmt.Sprintf("fr_FR.%s", c.User.Language.CharSet),
			fmt.Sprintf("fr_LU.%s", c.User.Language.CharSet),
			"fur_IT", "fy_NL", "fy_DE",
			fmt.Sprintf("ga_IE.%s", c.User.Language.CharSet),
			fmt.Sprintf("gd_GB.%s", c.User.Language.CharSet),
			"gez_ER", "gez_ET",
			fmt.Sprintf("gl_ES.%s", c.User.Language.CharSet),
			"gu_IN",
			fmt.Sprintf("gv_GB.%s", c.User.Language.CharSet),
			"ha_NG", "hak_TW",
			fmt.Sprintf("he_IL.%s", c.User.Language.CharSet),
			"hi_IN", "hif_FJ", "hne_IN",
			fmt.Sprintf("hr_HR.%s", c.User.Language.CharSet),
			fmt.Sprintf("hsb_DE.%s", c.User.Language.CharSet),
			"ht_HT",
			fmt.Sprintf("hu_HU.%s", c.User.Language.CharSet),
			"hy_AM", "ia_FR",
			fmt.Sprintf("id_ID.%s", c.User.Language.CharSet),
			"ig_NG", "ik_CA",
			fmt.Sprintf("is_IS.%s", c.User.Language.CharSet),
			fmt.Sprintf("it_CH.%s", c.User.Language.CharSet),
			fmt.Sprintf("it_IT.%s", c.User.Language.CharSet),
			"iu_CA",
			fmt.Sprintf("ja_JP.%s", c.User.Language.CharSet),
			fmt.Sprintf("ka_GE.%s", c.User.Language.CharSet),
			"kab_DZ",
			fmt.Sprintf("kk_KZ.%s", c.User.Language.CharSet),
			fmt.Sprintf("kl_GL.%s", c.User.Language.CharSet),
			"km_KH", "kn_IN",
			fmt.Sprintf("ko_KR.%s", c.User.Language.CharSet),
			"kok_IN", "ks_IN",
			fmt.Sprintf("ku_TR.%s", c.User.Language.CharSet),
			fmt.Sprintf("kw_GB.%s", c.User.Language.CharSet),
			"ky_KG", "lb_LU",
			fmt.Sprintf("lg_UG.%s", c.User.Language.CharSet),
			"li_BE", "li_NL", "lij_IT", "ln_CD", "lo_LA",
			fmt.Sprintf("lt_LT.%s", c.User.Language.CharSet),
			fmt.Sprintf("lv_LV.%s", c.User.Language.CharSet),
			"lzh_TW", "mag_IN", "mai_IN", "mai_NP", "mfe_MU",
			fmt.Sprintf("mg_MG.%s", c.User.Language.CharSet),
			"mhr_RU",
			fmt.Sprintf("mi_NZ.%s", c.User.Language.CharSet),
			"miq_NI", "mjw_IN",
			fmt.Sprintf("mg_MK.%s", c.User.Language.CharSet),
			"ml_IN", "mn_MN", "mni_MN", "mnw_MM", "mr_IN",
			fmt.Sprintf("ms_MY.%s", c.User.Language.CharSet),
			fmt.Sprintf("mt_MT.%s", c.User.Language.CharSet),
			"my_MM", "nan_TW",
			fmt.Sprintf("nb_NO.%s", c.User.Language.CharSet),
			"nds_DE", "nds_NL", "ne_NP", "nhn_MX", "niu_NU", "niu_NZ", "nl_AW",
			fmt.Sprintf("nl_BE.%s", c.User.Language.CharSet),
			fmt.Sprintf("nl_NL.%s", c.User.Language.CharSet),
			fmt.Sprintf("nn_NO.%s", c.User.Language.CharSet),
			"nr_ZA",
			"nso_ZA",
			fmt.Sprintf("oc_FR.%s", c.User.Language.CharSet),
			"om_ET",
			fmt.Sprintf("om_KE.%s", c.User.Language.CharSet),
			"or_IN", "os_RU", "pa_IN", "pa_PK", "pap_AW", "pap_CW",
			fmt.Sprintf("pl_PL.%s", c.User.Language.CharSet),
			"ps_AF",
			fmt.Sprintf("pt_BR.%s", c.User.Language.CharSet),
			fmt.Sprintf("pt_PT.%s", c.User.Language.CharSet),
			"quz_PE", "raj_IN",
			fmt.Sprintf("ro_RO.%s", c.User.Language.CharSet),
			fmt.Sprintf("ru_RU.%s", c.User.Language.CharSet),
			fmt.Sprintf("ru_UA.%s", c.User.Language.CharSet),
			"rw_RW", "sa_IN", "sah_RU", "sat_IN", "sc_IT", "sd_IN", "se_NO", "sgs_LT", "shn_MM", "shs_CA", "si_LK", "sid_ET",
			fmt.Sprintf("sk_SK.%s", c.User.Language.CharSet),
			fmt.Sprintf("sl_SI.%s", c.User.Language.CharSet),
			"sm_WS",
			fmt.Sprintf("so_DJ.%s", c.User.Language.CharSet),
			"so_ET",
			fmt.Sprintf("so_KE.%s", c.User.Language.CharSet),
			fmt.Sprintf("so_SO.%s", c.User.Language.CharSet),
			fmt.Sprintf("sq_AL.%s", c.User.Language.CharSet),
			"sq_MK", "sr_ME", "sr_RS", "ss_ZA",
			fmt.Sprintf("st_ZA.%s", c.User.Language.CharSet),
			fmt.Sprintf("sv_FI.%s", c.User.Language.CharSet),
			fmt.Sprintf("sv_SE.%s", c.User.Language.CharSet),
			"sw_KE", "sw_TZ", "szl_PL", "ta_IN", "ta_LK",
			fmt.Sprintf("tcy_IN.%s", c.User.Language.CharSet),
			"te_IN",
			fmt.Sprintf("tg_TJ.%s", c.User.Language.CharSet),
			fmt.Sprintf("th_TH.%s", c.User.Language.CharSet),
			"the_NP", "ti_ER", "ti_ET", "tig_ER", "tk_TM",
			fmt.Sprintf("tl_PH.%s", c.User.Language.CharSet),
			"tn_ZA", "to_TO", "tpi_PG",
			fmt.Sprintf("tr_CY.%s", c.User.Language.CharSet),
			fmt.Sprintf("tr_TR.%s", c.User.Language.CharSet),
			"ts_ZA", "tt_RU", "ug_CN",
			fmt.Sprintf("uk_UA.%s", c.User.Language.CharSet),
			"unm_US", "ur_IN", "ur_PK",
			fmt.Sprintf("uz_UZ.%s", c.User.Language.CharSet),
			"ve_ZA", "vi_VN",
			fmt.Sprintf("wa_BE.%s", c.User.Language.CharSet),
			"wae_CH", "wal_ET", "wo_SN",
			fmt.Sprintf("xh_ZA.%s", c.User.Language.CharSet),
			fmt.Sprintf("yi_US.%s", c.User.Language.CharSet),
			"yo_NG", "yue_HK", "yuw_PG",
			fmt.Sprintf("zh_CN.%s", c.User.Language.CharSet),
			fmt.Sprintf("zh_HK.%s", c.User.Language.CharSet),
			fmt.Sprintf("zh_SG.%s", c.User.Language.CharSet),
			fmt.Sprintf("zh_TW.%s", c.User.Language.CharSet),
			fmt.Sprintf("zh_ZA.%s", c.User.Language.CharSet),
		},
		Size: 15,
	}

	_, locale_choice, err := prompt.Run()

	if err != nil {
		log.Printf("Locale selection failed %v\n", err)
		return
	}

	c.User.Language.Locale = locale_choice
	SaveConfig(c)
}

func LoadConfig(c *Config) {
	switch c.Format {
	case "yaml", "YAML":
		config_file, err := ioutil.ReadFile("shobuarch_config.yaml")
		log.Printf("\n%s\n", config_file)
		if err != nil {
			log.Panic("Config file not found!")
		}
		_ = yaml.Unmarshal([]byte(config_file), &c)
	default:
		config_file, err := ioutil.ReadFile("shobuarch_config.json")
		log.Printf("\n%s\n", config_file)
		if err != nil {
			log.Panic("Config file not found!")
		}
		_ = json.Unmarshal([]byte(config_file), &c)
	}

	log.Println("Config successfully loaded!")
}

func SaveConfig(c *Config) {
	switch c.Format {
	case "yaml", "YAML":
		file, _ := yaml.Marshal(c)
		_ = ioutil.WriteFile("./shobuarch_config.yaml", file, 0644)
	default:
		file, _ := json.MarshalIndent(c, "", " ")
		_ = ioutil.WriteFile("./shobuarch_config.json", file, 0644)
	}
}
