package models

import ()

type DataFrameTargetUrl struct {
	TargetUrls map[string]string `json:"target_urls"`
}

func (df *DataFrameTargetUrl) Urls() []string {
	s := make([]string, len(df.TargetUrls))
	for _, targetUrl := range df.TargetUrls {
		s = append(s, targetUrl)
	}
	return s
}
