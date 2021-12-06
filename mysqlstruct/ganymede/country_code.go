package ganymede

type CountryCode struct {
	ID         int64  `json:"id" db:"id"`
	Cname      string `json:"cname" db:"cname"`
	Code       string `json:"country_code" db:"country_code"`
	CreateTime string `json:"create_time,omitempty" db:"create_time"`
	UpdateTime string `json:"update_time,omitempty" db:"update_time"`
}
