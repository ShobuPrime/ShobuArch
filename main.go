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
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
	i "github.com/ShobuPrime/ShobuArch/pkg/install"

	"github.com/manifoldco/promptui"
)

func clear() {

	switch host_os := runtime.GOOS; host_os {
	case "windows":
		clear := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		clear.Stdout = os.Stdout
		err := clear.Run()
		if err != nil {
			log.Println(err)
		}
	default:
		clear := exec.Command("clear") //Linux example, its tested
		clear.Stdout = os.Stdout
		err := clear.Run()
		if err != nil {
			log.Println(err)
		}
	}
}

func MainLogo() {

	fmt.Println(`
	-------------------------------------------------------------------------
	███████╗██╗  ██╗ ██████╗ ██████╗ ██╗   ██╗ █████╗ ██████╗  ██████╗██╗  ██╗
	██╔════╝██║  ██║██╔═══██╗██╔══██╗██║   ██║██╔══██╗██╔══██╗██╔════╝██║  ██║
	███████╗███████║██║   ██║██████╔╝██║   ██║███████║██████╔╝██║     ███████║
	╚════██║██╔══██║██║   ██║██╔══██╗██║   ██║██╔══██║██╔══██╗██║     ██╔══██║
	███████║██║  ██║╚██████╔╝██████╔╝╚██████╔╝██║  ██║██║  ██║╚██████╗██║  ██║
	╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
	-------------------------------------------------------------------------
			Automated Arch Linux Tools
	-------------------------------------------------------------------------
	`)
}

func getFormat(c *conf.Config) {
	prompt := promptui.Select{
		Label: "Preferred configuration format?",
		Items: []string{"JSON", "YAML"},
	}
	_, format_choice, err := prompt.Run()
	if err != nil {
		log.Printf("Format selection failed %v\n", err)
		return
	}

	c.Format = format_choice
	log.Printf("Config format is %q", c.Format)
}

func config(c *conf.Config) {

	conf.Logo()
	conf.SetKernel(c)
	conf.BasePkgs(c)
	clear()

	conf.Logo()
	conf.SetUserInfo(c)
	clear()

	conf.Logo()
	conf.SetHostInfo(c)
	clear()

	conf.Logo()
	c.Storage.SystemDisk, c.Storage.SystemDiskID, c.Storage.SystemDiskRota = conf.SelectDisk()
	conf.SetFilesystem(c)
	clear()

	conf.Logo()
	conf.SetInstallType(c)
	conf.SelectDesktopEnv(c)
	conf.SelectAUR(c)
	conf.Flatpak(c)
	conf.PWA(c)
	clear()

	conf.Logo()
	conf.SetTimezone(c)
	conf.SetKeyboardLayout(c)
	clear()

	conf.Logo()
	conf.SetLocale(c)
	clear()
}

func preinstall(c *conf.Config) {
	i.PreInstallLogo()
	i.Prerequisites(c)
	i.FormatDisks(c)
	i.Mirrors(c)
	i.ArchInstall(c)
	clear()

	log.Println(`
	-------------------------------------------------------------------------
				SYSTEM READY FOR SETUP
	-------------------------------------------------------------------------
	`)
}

func setup(c *conf.Config) {
	i.SetupLogo()
	i.SetupHostname(c)
	i.SetupNetwork(c)
	i.SetupMirrors(c)

	// Remember, exec.Command doesn't use quotes. trying without them first

	i.SetupResources(c)
	i.SetupLanguage(c)
	i.SetupBaseSystem(c)
	i.SetupServices(c)
	i.SetupCustomRepos(c)
	i.SetupMicrocode(c)
	i.SetupGraphics(c)
	i.SetupUser(c)
	i.SetupAUR(c)
	i.SetupFlatpaks(c)
	i.SetupEFI(c)

	clear()

	log.Println(`
	-------------------------------------------------------------------------
				SYSTEM READY FOR USER
	-------------------------------------------------------------------------
	`)
}

func user(c *conf.Config) {
	i.UserLogo()
	i.UserLocale(c)
	i.UserPackages(c)
	i.UserPWAs(c)
	i.UserShell(c)
	clear()

	log.Println(`
	-------------------------------------------------------------------------
				SYSTEM READY FOR POST-CHECKS
	-------------------------------------------------------------------------
	`)
}

func postinstall(c *conf.Config) {
	i.PostInstallLogo()
	i.PostInstallCleanup(c)
	i.PostInstallUnmount(c)

	log.Println(`
	-------------------------------------------------------------------------
				INSTALLATION COMPLETE!!!
	-------------------------------------------------------------------------
	`)
}

func main() {

	// Initialize Logs
	LOG_FILE := "./shobuarch.log"
	log_file, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log_out := io.MultiWriter(os.Stdout, log_file)
	defer log_file.Close()

	log.SetOutput(log_out)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Initializing ShobuArch logs...")

	clear()
	MainLogo()
	c := conf.Config{}

	// https://gobyexample.com/command-line-flags
	method_flag := flag.String("method", "m", `'a': Automated, 'm': Manual`)
	config_flag := flag.String("config", "n", `'y': Load config, 'n': Fresh config`)
	format_flag := flag.String("format", "json", `Accepted: 'JSON' || 'YAML'`)
	flag.Parse()

	c.Format = *format_flag

	if *method_flag == "a" {
		if *config_flag != "y" {
			log.Println("Automated installation detected, but 'config' flag is not set properly.")
			log.Fatalln("Type '--help' as a command-line argument to see available flag options")
		} else {
			conf.LoadConfig(&c)
			preinstall(&c)
			setup(&c)
			user(&c)
			postinstall(&c)
		}
	} else {
		log.Println("Manual installation detected.")
		prompt := promptui.Select{
			Label: "Start fresh, or load existing config?",
			Items: []string{"Start fresh", "Load config"},
		}
		_, prompt_choice, err := prompt.Run()
		if err != nil {
			log.Printf("Config selection failed %v\n", err)
			return
		}

		getFormat(&c)

		if *config_flag == "y" || prompt_choice == "Load config" {
			// Future release: load directory and use recursive ui to find config in file system
			// Or, we can accept filename as flag, config as singleline json, etc.
			conf.LoadConfig(&c)
			preinstall(&c)
			setup(&c)
			user(&c)
			postinstall(&c)
		} else {
			clear()
			config(&c)
			preinstall(&c)
			setup(&c)
			user(&c)
			postinstall(&c)
		}
	}
}
