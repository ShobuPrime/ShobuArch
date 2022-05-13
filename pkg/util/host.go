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
	"encoding/json"
	"log"
	"os/exec"
)

type HostStatus struct {
	Hostname                  string      `json:"Hostname"`
	StaticHostname            string      `json:"StaticHostname"`
	PrettyHostname            interface{} `json:"PrettyHostname"`
	DefaultHostname           string      `json:"DefaultHostname"`
	HostnameSource            string      `json:"HostnameSource"`
	IconName                  string      `json:"IconName"`
	Chassis                   string      `json:"Chassis"`
	Deployment                interface{} `json:"Deployment"`
	Location                  interface{} `json:"Location"`
	KernelName                string      `json:"KernelName"`
	KernelRelease             string      `json:"KernelRelease"`
	KernelVersion             string      `json:"KernelVersion"`
	OperatingSystemPrettyName string      `json:"OperatingSystemPrettyName"`
	OperatingSystemCPEName    interface{} `json:"OperatingSystemCPEName"`
	OperatingSystemHomeURL    string      `json:"OperatingSystemHomeURL"`
	HardwareVendor            string      `json:"HardwareVendor"`
	HardwareModel             string      `json:"HardwareModel"`
	ProductUUID               interface{} `json:"ProductUUID"`
}

func GetHostStatus() *HostStatus {
	hostnamectl, _ := exec.Command(
		"hostnamectl",
		"status",
		"--json=pretty",
	).Output()

	hostnamectl_struct := HostStatus{}
	err := json.Unmarshal(hostnamectl, &hostnamectl_struct)
	if err != nil {
		log.Fatalln("Invalid LSBLK Struct")
	}

	return &hostnamectl_struct
}
