package models

import (
	"fmt"
)

type TargetUrlSource struct {
	ImageSourceUrl       string `json:"image_source_url"`
	TargetUrl            string `json:"target_url"`
	StorageDirectoryPath string `json:"storage_directory_path"`
	StorageFilePath      string `json:"storage_file_path"`
	StorageFullPath      string `json:"storage_full_path"`
}

func NewTargetUrlSources(ImageSourceUrl, TargetUrl, StorageDirectoryPath, StorageFilePath, StorageFullPath string) *TargetUrlSource {
	return &TargetUrlSource{
		ImageSourceUrl,
		TargetUrl,
		StorageDirectoryPath,
		StorageFilePath,
		StorageFullPath,
	}
}

func (t *TargetUrlSource) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (image_source_url, target_url, storage_directory_path, storage_file_path, storage_full_path) VALUES (?, ?, ?, ?, ?)", TABLE_NAME_TARGET_URL_SOURCES)
	_, err := DbConnection.Exec(cmd, t.ImageSourceUrl, t.TargetUrl, t.StorageDirectoryPath, t.StorageFilePath, t.StorageFullPath)
	if err != nil {
		return err
	}
	return err
}

func (t *TargetUrlSource) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET target_url = ?, storage_directory_path = ?, storage_file_path = ?, storage_full_path = ? WHERE image_source_url = ?", TABLE_NAME_TARGET_URL_SOURCES)
	_, err := DbConnection.Exec(cmd, t.TargetUrl, t.StorageDirectoryPath, t.StorageFilePath, t.StorageFullPath, t.ImageSourceUrl)
	if err != nil {
		return err
	}
	return err
}

func GetImageSourceUrl(ImageSourceUrl string) *TargetUrlSource {
	cmd := fmt.Sprintf("SELECT image_source_url, target_url, storage_directory_path, storage_file_path, storage_full_path FROM %s WHERE image_source_url = ?", TABLE_NAME_TARGET_URL_SOURCES)
	row := DbConnection.QueryRow(cmd, ImageSourceUrl)
	var t TargetUrlSource
	err := row.Scan(&t.ImageSourceUrl, &t.TargetUrl, &t.StorageDirectoryPath, &t.StorageFilePath, &t.StorageFullPath)
	if err != nil {
		return nil
	}
	return NewTargetUrlSources(ImageSourceUrl, t.TargetUrl, t.StorageDirectoryPath, t.StorageFilePath, t.StorageFullPath)
}

func IsTargetUrlSource(ImageSourceUrl string) bool {
	t := GetImageSourceUrl(ImageSourceUrl)
	return t != nil
}

func GetImageSourceUrls(TargetUrl string) (dfTargetUrlSource *DataFrameTargetUrlSource, err error) {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE target_url = ?;`, TABLE_NAME_TARGET_URL_SOURCES)
	rows, err := DbConnection.Query(cmd, TargetUrl)
	if err != nil {
		return
	}
	defer rows.Close()

	dfTargetUrlSource = &DataFrameTargetUrlSource{}
	for rows.Next() {
		var TargetUrlSource TargetUrlSource
		rows.Scan(&TargetUrlSource.ImageSourceUrl, &TargetUrlSource.TargetUrl, &TargetUrlSource.StorageDirectoryPath, &TargetUrlSource.StorageFilePath, &TargetUrlSource.StorageFullPath)
		dfTargetUrlSource.ImageSourceUrls = append(dfTargetUrlSource.ImageSourceUrls, TargetUrlSource.ImageSourceUrl)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return dfTargetUrlSource, nil
}
