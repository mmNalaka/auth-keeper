package models

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Signup struct {
	Email    string                 `json:"email" validate:"required,email"`
	Password string                 `json:"password" validate:"required,min=8"`
	Data     map[string]interface{} `json:"data"`
	Role     string                 `json:"-"`
}
