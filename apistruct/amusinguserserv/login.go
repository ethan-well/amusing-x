package amusinguserserv

type LoginRequest struct {
	Type     int    `json:"type"`
	AreaCode string `json:"area_code"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
