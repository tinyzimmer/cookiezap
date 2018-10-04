package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func getCookiesTable(browser int) (table string) {
	if browser == CHROME_ID {
		table = CHROME_COOKIES_TABLE_NAME
	} else if browser == FIREFOX_ID {
		table = FIREFOX_COOKIES_TABLE_NAME
	}
	return
}

func getCookiesHostColumn(browser int) (column string) {
	if browser == CHROME_ID {
		column = CHROME_COOKIES_HOST_COLUMN
	} else if browser == FIREFOX_ID {
		column = FIREFOX_COOKIES_HOST_COLUMN
	}
	return
}

func deleteCookies(browser int, dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	checkErr(err, EXIT_SQLITE_ERROR)
	stmt, err := db.Prepare(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s LIKE ?",
			getCookiesTable(browser),
			getCookiesHostColumn(browser),
		),
	)
	checkErr(err, EXIT_SQLITE_ERROR)
	for _, x := range searchTerms {
		res, err := stmt.Exec(fmt.Sprintf("%%%s%%", x))
		checkErr(err, EXIT_SQLITE_ERROR)
		affected, err := res.RowsAffected()
		checkErr(err, EXIT_SQLITE_ERROR)
		fmt.Printf("%v cookie(s) related to %s deleted from %s\n", affected, x, resolveBrowserId(browser))
	}
}
