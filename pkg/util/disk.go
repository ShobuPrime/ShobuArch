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
	"strings"
)

type LSBLK struct {
	Blockdevices []BlockDevice
}

type BlockDevice struct {
	// lsblk from util-linux 2.38
	Alignment    int           `json:"alignment"`
	DiscAln      int           `json:"disc-aln"`
	Dax          bool          `json:"dax"`
	DiscGran     string        `json:"disc-gran"`
	DiscMax      string        `json:"disc-max"`
	DiscZero     bool          `json:"disc-zero"`
	Fsavail      interface{}   `json:"fsavail"`
	Fsroots      []interface{} `json:"fsroots"`
	Fssize       interface{}   `json:"fssize"`
	Fstype       string        `json:"fstype"`
	Fsused       interface{}   `json:"fsused"`
	Fsuse        interface{}   `json:"fsuse%"`
	Fsver        string        `json:"fsver"`
	Group        string        `json:"group"`
	Hctl         string        `json:"hctl"`
	Hotplug      bool          `json:"hotplug"`
	Kname        string        `json:"kname"`
	Label        string        `json:"label"`
	LogSec       int           `json:"log-sec"`
	MajMin       string        `json:"maj:min"`
	MinIo        int           `json:"min-io"`
	Mode         string        `json:"mode"`
	Model        string        `json:"model"`
	Name         string        `json:"name"`
	OptIo        int           `json:"opt-io"`
	Owner        string        `json:"owner"`
	Partflags    interface{}   `json:"partflags"`
	Partlabel    interface{}   `json:"partlabel"`
	Parttype     interface{}   `json:"parttype"`
	Parttypename interface{}   `json:"parttypename"`
	Partuuid     interface{}   `json:"partuuid"`
	Path         string        `json:"path"`
	PhySec       int           `json:"phy-sec"`
	Pkname       interface{}   `json:"pkname"`
	Pttype       string        `json:"pttype"`
	Ptuuid       string        `json:"ptuuid"`
	Ra           int           `json:"ra"`
	Rand         bool          `json:"rand"`
	Rev          string        `json:"rev"`
	Rm           bool          `json:"rm"`
	Ro           bool          `json:"ro"`
	Rota         bool          `json:"rota"`
	RqSize       int           `json:"rq-size"`
	Sched        string        `json:"sched"`
	Serial       string        `json:"serial"`
	Size         string        `json:"size"`
	Start        interface{}   `json:"start,omitempty"`
	State        string        `json:"state"`
	Subsystems   string        `json:"subsystems"`
	Mountpoint   interface{}   `json:"mountpoint"`
	Mountpoints  []interface{} `json:"mountpoints"`
	Tran         string        `json:"tran"`
	Type         string        `json:"type"`
	UUID         string        `json:"uuid"`
	Vendor       string        `json:"vendor"`
	Wsame        string        `json:"wsame"`
	Wwn          interface{}   `json:"wwn"`
	Zoned        string        `json:"zoned"`
	ZoneSz       int           `json:"zone-sz,omitempty"`
	ZoneWgran    int           `json:"zone-wgran,omitempty"`
	ZoneApp      int           `json:"zone-app,omitempty"`
	ZoneNr       int           `json:"zone-nr,omitempty"`
	ZoneOmax     int           `json:"zone-omax,omitempty"`
	ZoneAmax     int           `json:"zone-amax,omitempty"`
	Children     []BlockDevice `json:"children,omitempty"`
}

func ListDisk() *LSBLK {
	// util-linux 2.38 breaks accepting all columns
	// lsblk, _ := exec.Command(
	// 	"lsblk",
	// 	"-J",
	// 	"-O",
	// ).Output()

	// Temporary fix for util-linux 2.38 until columns are fixed
	lsblk_columns := []string{
		"alignment",
		"disc-aln",
		"dax",
		"disc-gran",
		"disc-max",
		"disc-zero",
		"fsavail",
		"fsroots",
		"fssize",
		"fstype",
		"fsused",
		"fsuse%",
		"fsver",
		"group",
		"hctl",
		"hotplug",
		"kname",
		"label",
		"log-sec",
		"maj:min",
		"min-io",
		"mode",
		"model",
		"name",
		"opt-io",
		"owner",
		"partflags",
		"partlabel",
		"parttype",
		"parttypename",
		"partuuid",
		"path",
		"phy-sec",
		"pkname",
		"pttype",
		"ptuuid",
		"ra",
		"rand",
		"rev",
		"rm",
		"ro",
		"rota",
		"rq-size",
		"sched",
		"serial",
		"size",
		"state",
		"subsystems",
		"mountpoint",
		"mountpoints",
		"tran",
		"type",
		"uuid",
		"vendor",
		"wsame",
		"wwn",
		"zoned",
	}

	lsblk, _ := exec.Command(
		"lsblk",
		"-J",
		"-o",
		strings.Join(lsblk_columns, ","),
	).Output()

	lsblk_struct := LSBLK{}
	err := json.Unmarshal(lsblk, &lsblk_struct)
	if err != nil {
		log.Fatalln("Invalid LSBLK Struct")
	}

	return &lsblk_struct
}
