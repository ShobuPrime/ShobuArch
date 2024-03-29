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
package util

import (
	"log"
	"os/exec"
	"strings"
)

type SBCTL struct {
	Installed  string `json:"Installed,omitempty"`
	OwnerGUID  string `json:"OwnerGUID,omitempty"`
	SetupMode  string `json:"SetupMode,omitempty"`
	SecureBoot string `json:"SecureBoot,omitempty"`
	VendorKeys string `json:"VendorKeys,omitempty"`
}

func SecureBootStatus() *SBCTL {

	log.Println("Grabbing Secure Boot Status...")

	sbctl, _ := exec.Command(
		"sbctl",
		"status",
	).CombinedOutput()

	sbctl_status := string(sbctl)

	return SBJSON(&sbctl_status)
}

func SecureBootCreateKeys() *[]string {

	log.Println("Creating Secure Boot keys...")

	sbctl := []string{
		"sbctl",
		"create-keys",
	}

	return &sbctl
}

func SecureBootEnrollKeys() *[]string {

	log.Println("Enrolling Secure Boot keys...")

	sbctl := []string{
		"sbctl",
		"enroll-keys",
		"--microsoft",
	}

	return &sbctl
}

func SecureBootSign(file *string) *[]string {
	log.Printf("Signing %s with Secure Boot Keys", *file)

	sbctl := []string{
		`sbctl`,
		`sign`,
		`-s`,
		*file,
	}

	return &sbctl
}

func SBJSON(sbctl *string) *SBCTL {

	sbctl_struct := &SBCTL{}

	sbctl_status := strings.Split(*sbctl, "\n")

	// ✓ - sbctl has native JSON output
	// ✘ - Doesn't work as of Release 0.9
	for i := range sbctl_status {
		if strings.HasPrefix(sbctl_status[i], "Installed:\t") {
			sbctl_struct.Installed = (strings.TrimPrefix(sbctl_status[i], "Installed:\t"))[4:]
		} else if strings.HasPrefix(sbctl_status[i], "Owner GUID:\t") {
			sbctl_struct.OwnerGUID = strings.TrimPrefix(sbctl_status[i], "Owner GUID:\t")
		} else if strings.HasPrefix(sbctl_status[i], "Setup Mode:\t") {
			sbctl_struct.SetupMode = (strings.TrimPrefix(sbctl_status[i], "Setup Mode:\t"))[4:]
		} else if strings.HasPrefix(sbctl_status[i], "Secure Boot:\t") {
			sbctl_struct.SecureBoot = (strings.TrimPrefix(sbctl_status[i], "Secure Boot:\t"))[4:]
		} else if strings.HasPrefix(sbctl_status[i], "Vendor Keys:\t") {
			sbctl_struct.VendorKeys = strings.TrimPrefix(sbctl_status[i], "Vendor Keys:\t")
		} else if sbctl_status[i] == "" {
			// Do nothing
			do := "nothing"
			_ = do
		}
	}

	return sbctl_struct
}
