package models

type DataFrameTargetUrl struct {
	TargetUrls []string `json:"target_urls"`
}

// func (df *DataFrameTargetUrl) Urls() []string {
// 	s := make([]string, len(df.TargetUrls))
// 	for _, targetUrl := range df.TargetUrls {
// 		s = append(s, targetUrl)
// 	}
// 	return s
// }
