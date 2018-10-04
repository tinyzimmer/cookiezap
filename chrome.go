package main

import (
	"fmt"
	"os/user"
	"runtime"
)

func getChromeCookiesLocation() string {
	usr, err := user.Current()
	checkErr(err, EXIT_NO_USER)
	if runtime.GOOS == "windows" {
		return fmt.Sprintf(WINDOWS_CHROME_COOKIE_LOC, usr.HomeDir)
	} else if runtime.GOOS == "darwin" {
		return fmt.Sprintf(MAC_CHROME_COOKIE_LOC, usr.HomeDir)
	} else {
		return fmt.Sprintf(LINUX_CHROME_COOKIE_LOC, usr.HomeDir)
	}
}
