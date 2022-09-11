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
	"os/exec"
	"strings"
)

type LSUSB struct {
	USBDevices []USBDevice
}

type USBDevice struct {
	Bus         string `json:"bus,omitempty"`
	Device      string `json:"device,omitempty"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

func ListUSB() *LSUSB {

	lsusb, _ := exec.Command(
		`lsusb`,
	).CombinedOutput()

	usb_devices := string(lsusb)

	return PCIUSB(&usb_devices)
}

func PCIUSB(lsusb *string) *LSUSB {

	usb_struct := &LSUSB{}

	device := USBDevice{}

	usb_list := strings.Split(*lsusb, "\n")

	for i := range usb_list {
		device.Bus = usb_list[i][:7]
		device.Device = usb_list[i][8:19]
		device.ID = usb_list[i][20:32]
		device.Description = usb_list[i][33:]
		usb_struct.USBDevices = append(usb_struct.USBDevices, device)
		device = USBDevice{}
	}

	return usb_struct
}