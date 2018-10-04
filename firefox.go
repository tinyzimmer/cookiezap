package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func getFirefoxCookiesLocations() (cookieFiles []string) {
	usr, err := user.Current()
	checkErr(err, EXIT_NO_USER)
	var profileDir string
	if runtime.GOOS == "windows" {
		profileDir = fmt.Sprintf(WINDOWS_FIREFOX_COOKIE_LOC, usr.HomeDir)
	} else if runtime.GOOS == "darwin" {
		profileDir = fmt.Sprintf(MAC_FIREFOX_COOKIE_LOC, usr.HomeDir)
	} else {
		profileDir = fmt.Sprintf(LINUX_FIREFOX_COOKIE_LOC, usr.HomeDir)
	}
	cookieFiles = getFirefoxProfiles(profileDir)
	return
}

func getFirefoxProfiles(profileDir string) (profiles []string) {
	files, err := ioutil.ReadDir(profileDir)
	checkErr(err, EXIT_BAD_PROFILE_DIR)
	for _, file := range files {
		if file.IsDir() {
			if strings.Contains(file.Name(), FIREFOX_PROFILE_DIR_INCLUDE) {
				profiles = append(
					profiles,
					filepath.Join(profileDir, file.Name(), FIREFOX_COOKIES_FILENAME),
				)
			}
		}
	}
	return
}
