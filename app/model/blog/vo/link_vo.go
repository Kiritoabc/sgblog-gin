package vo

type LinkVo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Address     string `json:"address" description:"网站地址"`
}
