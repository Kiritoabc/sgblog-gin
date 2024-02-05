package dto

type TagListDto struct {
	Name   string `json:"name" description:"标签名"`
	Remark string `json:"remark" description:"备注"`
}
