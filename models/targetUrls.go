package models

import (
	"fmt"
)

type TargetUrl struct {
	TargetUrl    string `json:"target_url"`
	OriginSource string `json:"origin_source"`
}

func NewTargetUrl(targetUrl string, originSource string) *TargetUrl {
	return &TargetUrl{
		targetUrl,
		originSource,
	}
}

func (t *TargetUrl) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (target_url, origin_source) VALUES (?, ?)", TABLE_NAME_TARGET_URLS)
	_, err := DbConnection.Exec(cmd, t.TargetUrl, t.OriginSource)
	if err != nil {
		return err
	}
	return err
}

func (t *TargetUrl) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET origin_source = ? WHERE target_url = ?", TABLE_NAME_TARGET_URLS)
	_, err := DbConnection.Exec(cmd, t.TargetUrl, t.OriginSource)
	if err != nil {
		return err
	}
	return err
}

func GetTargetUrl(targetUrl string) *TargetUrl {
	cmd := fmt.Sprintf("SELECT target_url, origin_source FROM %s WHERE target_url = ?", TABLE_NAME_TARGET_URLS)
	row := DbConnection.QueryRow(cmd, targetUrl)
	var t TargetUrl
	err := row.Scan(&t.TargetUrl, &t.OriginSource)
	if err != nil {
		return nil
	}
	return NewTargetUrl(targetUrl, t.OriginSource)
}

func IsTargetUrl(targetUrl string) bool {
	t := GetTargetUrl(targetUrl)
	return t != nil
}

func GetAllTargetUrl(limit int) (dfTargetUrl *DataFrameTargetUrl, err error) {
	cmd := fmt.Sprintf(`SELECT * FROM target_url, origin_source FROM %s LIMIT ?;`, TABLE_NAME_TARGET_URLS)
	rows, err := DbConnection.Query(cmd, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	dfTargetUrl = &DataFrameTargetUrl{}
	for rows.Next() {
		var targetUrl TargetUrl
		rows.Scan(&targetUrl.TargetUrl, &targetUrl.OriginSource)
		dfTargetUrl.TargetUrls[targetUrl.OriginSource] = targetUrl.TargetUrl
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return dfTargetUrl, nil
}