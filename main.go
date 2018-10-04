package main

import (
	"flag"
	"os"
)

const (
	WINDOWS_CHROME_COOKIE_LOC = "%s\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Cookies"
	LINUX_CHROME_COOKIE_LOC   = "%s/.config/google-chrome/Default/Cookies"
	MAC_CHROME_COOKIE_LOC     = "%s/Library/Application Support/Google/Chrome/Default/Cookies"

	FIREFOX_COOKIES_FILENAME    = "cookies.sqlite"
	FIREFOX_PROFILE_DIR_INCLUDE = ".default"
	WINDOWS_FIREFOX_COOKIE_LOC  = "%s\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles"
	LINUX_FIREFOX_COOKIE_LOC    = "%s/.mozilla/firefox"
	MAC_FIREFOX_COOKIE_LOC      = "%s/Library/Application Support/Firefox/Profiles"

	CHROME_COOKIES_TABLE_NAME  = "cookies"
	FIREFOX_COOKIES_TABLE_NAME = "moz_cookies"

	CHROME_COOKIES_HOST_COLUMN  = "host_key"
	FIREFOX_COOKIES_HOST_COLUMN = "host"

	CHROME_ID            = 1
	FIREFOX_ID           = 2
	CHROME_DISPLAY_NAME  = "Google Chrome"
	FIREFOX_DISPLAY_NAME = "Mozilla Firefox"

	EXIT_BAD_ARGS        = 1
	EXIT_SQLITE_ERROR    = 2
	EXIT_NO_USER         = 3
	EXIT_BAD_PROFILE_DIR = 4

	EXTRA_USAGE = "\n./cookiezap [args] <search_terms...>"
)

var (
	clearFirefox bool
	clearChrome  bool
	searchTerms  []string
)

func init() {
	flag.BoolVar(&clearFirefox, "f", false, "Clean Firefox cookies")
	flag.BoolVar(&clearChrome, "c", false, "Clean Chrome cookies")
	if len(os.Args) == 1 {
		exitBadArgs()
	}
	flag.Parse()
	if !clearFirefox && !clearChrome {
		exitBadArgs()
	}
	if flag.NArg() == 0 {
		exitBadArgs()
	} else {
		searchTerms = flag.Args()
	}
}

func main() {
	if clearChrome {
		deleteCookies(CHROME_ID, getChromeCookiesLocation())
	}
	if clearFirefox {
		for _, loc := range getFirefoxCookiesLocations() {
			deleteCookies(FIREFOX_ID, loc)
		}
	}
}
