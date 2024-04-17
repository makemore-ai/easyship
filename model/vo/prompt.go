package vo

type PromptVO struct {
	Id        int64   `json:"id"`
	PromptZh  *string `json:"prompt_zh"`  //中文prompt
	PromptEn  *string `json:"prompt_en"`  //英文prompt
	LabelName string  `json:"label_name"` //标签名称
}
