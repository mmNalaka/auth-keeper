package models

type LoginHistory struct {
	Id        int64  `json:"id"`
	ClintId   int64  `json:"clint_id"`
	UserId    int64  `json:"user_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	CreatedAt int64  `json:"created_at"`
}
