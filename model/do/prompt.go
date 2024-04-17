package do

type PromptEsDo struct {
	Id        int64
	PromptZh  *string //中文prompt
	PromptEn  *string //英文prompt
	LabelName string  //标签名称
}
