package dto

type EveryDayData struct {
	DataContent   map[string]DataContentDetail `json:"data_content" redis:"data_content"`
	ProvidePerson string                       `json:"provide_person" redis:"provide_person"`
}

type DataContentDetail struct {
	ThumbUp         int      `json:"thumb_up" redis:"thumb_up"`
	ThumbDown       int      `json:"thumb_down" redis:"thumb_down"`
	ThumbUpPerson   []string `json:"thumb_up_person" redis:"thumb_up_person"`
	ThumbDownPerson []string `json:"thumb_down_person" redis:"thumb_down_person"`
}
