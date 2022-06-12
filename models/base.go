package models

import (
	"database/sql"
	"fmt"
	"github.com/notpop/url_getter/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	TABLE_NAME_TARGET_URLS = "target_urls"
)

var DbConnection *sql.DB

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			target_url STRING PRIMARY KEY NOT NULL,
			origin_source STRING)`, TABLE_NAME_TARGET_URLS)
	DbConnection.Exec(cmd)
}
