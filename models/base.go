package models

import (
	"database/sql"
	"fmt"
	"github.com/notpop/url_getter/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	TABLE_NAME_TARGET_URLS        = "target_urls"
	TABLE_NAME_TARGET_URL_SOURCES = "target_url_sources"
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
			origin_source STRING NOT NULL,
			is_completed BOOLEAN DEFAULT FALSE NOT NULL)`, TABLE_NAME_TARGET_URLS)
	DbConnection.Exec(cmd)

	cmd = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			image_source_url STRING PRIMARY KEY NOT NULL,
			target_url STRING NOT NULL,
			storage_directory_path STRING NOT NULL,
			storage_file_path STRING NOT NULL,
			storage_full_path STRING NOT NULL)`, TABLE_NAME_TARGET_URL_SOURCES)
	DbConnection.Exec(cmd)
}
