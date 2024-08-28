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

type LSPCI struct {
	PCIDevices []PCIDevice
}

type PCIDevice struct {
	Slot       string `json:"Slot,omitempty"`
	Class      string `json:"Class,omitempty"`
	Vendor     string `json:"Vendor,omitempty"`
	Device     string `json:"Device,omitempty"`
	SVendor    string `json:"SVendor,omitempty"`
	SDevice    string `json:"SDevice,omitempty"`
	Rev        string `json:"Rev,omitempty"` //int
	ProgIf     string `json:"ProgIf,omitempty"`
	PhySlot    string `json:"PhySlot,omitempty"`    //int
	NUMANode   string `json:"NUMANode,omitempty"`   //int
	IOMMUGroup string `json:"IOMMUGroup,omitempty"` //int
}

func ListPCI() *LSPCI {

	lspci, _ := exec.Command(
		`lspci`,
		`-vmmq`, // [-v] Verbose --[-mm] machine-readable output -- [-q] query database for unknown IDs
	).CombinedOutput()

	pci_devices := string(lspci)

	return PCIJSON(&pci_devices)
}

func PCIJSON(lspci *string) *LSPCI {

	pci_struct := &LSPCI{}

	device := PCIDevice{}

	pci_list := strings.Split(*lspci, "\n")

	for i := range pci_list {
		if strings.HasPrefix(pci_list[i], "Slot:\t") {
			device.Slot = strings.TrimPrefix(pci_list[i], "Slot:\t")
		} else if strings.HasPrefix(pci_list[i], "Class:\t") {
			device.Class = strings.TrimPrefix(pci_list[i], "Class:\t")
		} else if strings.HasPrefix(pci_list[i], "Vendor:\t") {
			device.Vendor = strings.TrimPrefix(pci_list[i], "Vendor:\t")
		} else if strings.HasPrefix(pci_list[i], "Device:\t") {
			device.Device = strings.TrimPrefix(pci_list[i], "Device:\t")
		} else if strings.HasPrefix(pci_list[i], "SVendor:\t") {
			device.SVendor = strings.TrimPrefix(pci_list[i], "SVendor:\t")
		} else if strings.HasPrefix(pci_list[i], "SDevice:\t") {
			device.SDevice = strings.TrimPrefix(pci_list[i], "SDevice:\t")
		} else if strings.HasPrefix(pci_list[i], "Rev:\t") {
			device.Rev = strings.TrimPrefix(pci_list[i], "Rev:\t")
		} else if strings.HasPrefix(pci_list[i], "ProgIf:\t") {
			device.ProgIf = strings.TrimPrefix(pci_list[i], "ProgIf:\t")
		} else if strings.HasPrefix(pci_list[i], "PhySlot:\t") {
			device.PhySlot = strings.TrimPrefix(pci_list[i], "PhySlot:\t")
		} else if strings.HasPrefix(pci_list[i], "NUMANode:\t") {
			device.NUMANode = strings.TrimPrefix(pci_list[i], "NUMANode:\t")
		} else if strings.HasPrefix(pci_list[i], "IOMMUGroup:\t") {
			device.IOMMUGroup = strings.TrimPrefix(pci_list[i], "IOMMUGroup:\t")
		} else if pci_list[i] == "" {
			// If device is not an empty struct, append
			if (device != PCIDevice{}) {
				pci_struct.PCIDevices = append(pci_struct.PCIDevices, device)
			}
			device = PCIDevice{}
		}
	}

	return pci_struct
}
