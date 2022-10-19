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
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Define regular expression for Linux and Windows linebreaks
var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func ReadFile(filePath *string, fileName *string) *[]string {

	// Grab current directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Changing directory to %q", *filePath)
	if err := os.Chdir(*filePath); err != nil {
		log.Fatalln(err)
	}

	file, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatalln(err)
	}

	fileContents := lineBreakRegExp.Split(string(file), -1)

	log.Printf("Reading contents of %q", *fileName)
	for i, line := range fileContents {
		log.Println(i, "\t", line)
	}

	log.Println("Returning to original directory")
	// Return to original directory
	if err := os.Chdir(pwd); err != nil {
		log.Fatalln(err)
	}

	return &fileContents
}

func WriteFile(filePath *string, fileName *string, fileContents *[]string, fileFlag int, filePerm int) {
	// Grab current directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	
	log.Printf("Changing directory to %q", *filePath)
	if err := os.Chdir(*filePath); err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile(filepath.Join(*filePath, *fileName), fileFlag, fs.FileMode(filePerm))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(`Saving settings`)
	if _, err := f.Write([]byte(strings.Join(*fileContents, "\n"))); err != nil {
		log.Fatalln(err)
	}
	log.Println(`Closing file`)
	if err := f.Close(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Done!")

	log.Println("Returning to original directory")
	// Return to original directory
	if err := os.Chdir(pwd); err != nil {
		log.Fatalln(err)
	}
}
