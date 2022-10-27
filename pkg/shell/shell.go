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
package shell

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	conf "github.com/ShobuPrime/ShobuArch/pkg/config"
)

func Shell(cmd *string) string {

	program := exec.Command("/bin/bash", "-c", *cmd)
	stdout, err := program.StdoutPipe()
	// // https://stackoverflow.com/a/35995372
	// program.Stderr = program.Stdout
	if err != nil {
		log.Println(err)
	}

	stderr, err := program.StderrPipe()
	if err != nil {
		log.Println(err)
	}

	combined_output := io.MultiReader(stdout, stderr)

	err = program.Start()
	log.Printf("Executing: %q\n", program)
	if err != nil {
		log.Println(err)
	}

	output_log := ""

	// print the output of the subprocess
	// scanner := bufio.NewScanner(stdout)
	scanner := bufio.NewScanner(combined_output)
	for scanner.Scan() {
		m := scanner.Text()
		output_log += m + "\n"
		log.Printf("%s\n", m)
	}

	err = program.Wait()
	if err != nil {
		log.Println(err)
	}

	fmt.Println()

	return string(output_log)
}

// For whatever reason, arch_chroot is giving me lots of problems at the moment with commands like sed and echo
// Unlike shell function above, we will attempt each command as a list of args, rather than raw string
// Nuances include: Shell needs single quote for sed, but exec.Command does not use shell.
// Consequently, no single quotes around argument
func Arch_chroot(cmd_args *[]string, runuser bool, c *conf.Config) string {

	log.Println("Preparing arch-chroot command---------------------------")

	arch_chroot := []string{`arch-chroot`, `/mnt`}

	switch runuser {
	case true:
		arch_chroot = []string{"arch-chroot", "/mnt", "/usr/bin/runuser", "-u", c.User.Username, "--"}
		// arch_chroot = []string{"arch-chroot", "-u", c.Username, "/mnt",}
	}

	cmd := []string{}
	cmd = append(cmd, arch_chroot...)
	cmd = append(cmd, *cmd_args...)

	program := exec.Command(cmd[0], cmd[1:]...)
	stdout, err := program.StdoutPipe()
	if err != nil {
		log.Println(err)
	}

	stderr, err := program.StderrPipe()
	if err != nil {
		log.Println(err)
	}

	combined_output := io.MultiReader(stdout, stderr)

	err = program.Start()
	log.Printf("Executing: %q\n", program)
	if err != nil {
		log.Println(err)
	}

	output_log := ""

	// print the output of the subprocess
	// scanner := bufio.NewScanner(stdout)
	scanner := bufio.NewScanner(combined_output)
	for scanner.Scan() {
		m := scanner.Text()
		output_log += m + "\n"
		log.Printf("%s\n", m)
	}

	err = program.Wait()
	if err != nil {
		log.Println(err)
	}

	fmt.Println()

	return string(output_log)
}

func Systemd_nspawn(cmd_args *[]string, boot bool, c *conf.Config) string {
	// https://wiki.archlinux.org/title/Systemd-nspawn
	log.Println("Preparing systemd-nspawn command---------------------------")

	systemd_nspawn := []string{`systemd-nspawn`, `-D`, `/mnt`}

	switch boot {
	case true:
		systemd_nspawn = []string{"systemd-nspawn", `-bnq`, `-D`, `/mnt`}
	}

	cmd := []string{}
	cmd = append(cmd, systemd_nspawn...)
	cmd = append(cmd, *cmd_args...)

	program := exec.Command(cmd[0], cmd[1:]...)
	stdout, err := program.StdoutPipe()
	if err != nil {
		log.Println(err)
	}

	stderr, err := program.StderrPipe()
	if err != nil {
		log.Println(err)
	}

	combined_output := io.MultiReader(stdout, stderr)

	err = program.Start()
	log.Printf("Executing: %q\n", program)
	if err != nil {
		log.Println(err)
	}

	output_log := ""

	// print the output of the subprocess
	// scanner := bufio.NewScanner(stdout)
	scanner := bufio.NewScanner(combined_output)
	for scanner.Scan() {
		m := scanner.Text()
		output_log += strings.TrimSpace(m) + "\n"
		log.Printf("%s\n", strings.TrimSpace(m))
	}

	err = program.Wait()
	if err != nil {
		log.Println(err)
	}

	fmt.Println()

	return string(output_log)
}
