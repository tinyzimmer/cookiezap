package main

import (
	"flag"
	"fmt"
	"os"
)

func exitBadArgs() {
	flag.Usage()
	fmt.Println(EXTRA_USAGE)
	os.Exit(EXIT_BAD_ARGS)
}

func resolveBrowserId(id int) (browser string) {
	if id == CHROME_ID {
		browser = CHROME_DISPLAY_NAME
	} else if id == FIREFOX_ID {
		browser = FIREFOX_DISPLAY_NAME
	}
	return
}

func checkErr(err error, errCode int) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(errCode)
	}
}
